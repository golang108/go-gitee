//Copyright magesfc bright.ma
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Package gitee golang 版本的 gitee API 实现 模仿 go-github 实现的
package gitee

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/google/go-querystring/query"
)

const (
	Version          = "v1.0.0"
	defaultBaseURL   = "https://gitee.com/api/v5/"
	defaultUserAgent = "go-gitee" + "/" + Version
)

var errNonNilContext = errors.New("context must be non-nil")

// A Client manages communication with the gitee API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the GitHub API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the gitee API.
	Users         *UsersService
	Repositories  *RepositoriesService
	Search        *SearchService
	PullRequests  *PullRequestsService
	Organizations *OrganizationsService
	Miscellaneous *MiscellaneousService
}

type service struct {
	client *Client
}

// Client returns the http.Client used by this GitHub client.
func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.client
	return &clientCopy
}

// NewClient returns a new gitee API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient // 等价于 &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: defaultUserAgent}
	c.common.client = c

	c.Users = (*UsersService)(&c.common)
	c.Repositories = (*RepositoriesService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.PullRequests = (*PullRequestsService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.Miscellaneous = (*MiscellaneousService)(&c.common)

	return c

}

// Response is a gitee API response. This wraps the standard http.Response
// returned from gitee and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response

	// These fields provide the page values for paginating through a set of
	// results. Any or all of these may be set to the zero value for
	// responses that are not part of a paginated set, or for which there
	// are no additional pages.
	//
	// These fields support what is called "offset pagination" and should
	// be used with the ListOptions struct.
	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int

	TotalCount int
	TotalPage  int
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.populatePageValues()
	response.parseOther()
	return response
}

func (r *Response) parseOther() {
	if tc := r.Header.Get("Total_count"); tc != "" {
		r.TotalCount, _ = strconv.Atoi(tc)
	}
	if tc := r.Header.Get("Total_page"); tc != "" {
		r.TotalPage, _ = strconv.Atoi(tc)
	}
}

// populatePageValues parses the HTTP Link response headers and populates the
// various pagination link values in the Response.解析 响应头 里面 header 里面的 分页相关的值
func (r *Response) populatePageValues() {
	if links, ok := r.Response.Header["Link"]; ok && len(links) > 0 {
		for _, link := range strings.Split(links[0], ",") {
			segments := strings.Split(strings.TrimSpace(link), ";")
			// link must at least have href and rel
			if len(segments) < 2 {
				continue
			}
			// ensure href is properly formatted
			if !strings.HasPrefix(segments[0], "<") || !strings.HasSuffix(segments[0], ">") {
				continue
			}

			// try to pull out page parameter
			url, err := url.Parse(segments[0][1 : len(segments[0])-1])
			if err != nil {
				continue
			}

			q := url.Query()
			page := q.Get("page")
			perPage := q.Get("per_page")

			if page == "" && perPage == "" {
				continue
			}

			for _, segment := range segments[1:] {
				switch strings.TrimSpace(segment) {
				case `rel="next"`, `rel='next'`:
					r.NextPage, _ = strconv.Atoi(page)
				case `rel="prev"`, `rel='prev'`: //前一页
					r.PrevPage, _ = strconv.Atoi(page)
				case `rel="first"`, `rel='first'`: //第一页，这个好像总是 1 吧
					r.FirstPage, _ = strconv.Atoi(page)
				case `rel="last"`, `rel='last'`: //在 Header Link 属性中 有 next 表示下一页的page的值的，有last 表示最后一页的值的。
					r.LastPage, _ = strconv.Atoi(page)
				}
			}
		} // end for
	} // end if Response.Header["Link"]
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

type AcceptedError struct {
	// Raw contains the response body.
	Raw []byte
}

func (*AcceptedError) Error() string {
	return "job scheduled on gitee side; try again later"
}

// Is returns whether the provided error equals this error.
func (ae *AcceptedError) Is(target error) bool {
	v, ok := target.(*AcceptedError)
	if !ok {
		return false
	}
	return bytes.Compare(ae.Raw, v.Raw) == 0
}

type ErrorResponse struct {
	Response *http.Response         // HTTP response that caused this error
	ErrorMap map[string]interface{} `json:"error"`   // more detail on individual errors
	Message  string                 `json:"message"` // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d: %v, %v, %+v",
		r.Response.Request.Method,
		sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode,
		r.Response.Status,
		r.ErrorMap, r.Message,
	)
}

func CheckResponse(r *http.Response) error {
	if r.StatusCode == http.StatusAccepted { // = 202 equal to 202 Accepted.
		return &AcceptedError{}
	}
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	switch {
	case r.StatusCode == http.StatusUnauthorized: // 401 error
		return errorResponse
	case r.StatusCode == http.StatusForbidden: // 403 error
		return errorResponse
	default: // 以上几个报错, github 这里的报错会有 是由于api 接口访问频率限制等等
		return errorResponse
	}
}

// parseBoolResponse determines the boolean result from a API response.
// Several API methods return boolean responses indicated by the HTTP
// status code in the response (true indicated by a 204, false indicated by a
// 404). This helper function will determine that result and hide the 404
// error if present. Any other error will be returned through as-is.
func parseBoolResponse(err error) (bool, error) {
	if err == nil {
		return true, nil
	}

	if err, ok := err.(*ErrorResponse); ok && err.Response.StatusCode == http.StatusNotFound {
		// Simply false. In this one case, we do not pass the error through.
		return false, nil
	}

	// some other real error occurred
	return false, err
}

// BareDo sends an API request and lets you handle the api response. If an error
// or API Error occurs, the error will contain more information. Otherwise you
// are supposed to read and close the response's Body. If rate limit is exceeded
// and reset time is in the future, BareDo returns *RateLimitError immediately
// without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is
// canceled or times out, ctx.Err() will be returned.
func (c *Client) BareDo(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}

		return nil, err
	}

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil { // 这里 特殊处理 AcceptedError 这种错误，提交 返回了
		defer resp.Body.Close() // 这里就提前关闭了, 其他返回的 会在调用的 地方 func Do 里面关闭

		aerr, ok := err.(*AcceptedError)
		if ok {
			b, readErr := ioutil.ReadAll(resp.Body)
			if readErr != nil {
				return response, readErr
			}

			aerr.Raw = b
			err = aerr
		}
	}
	return response, err
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it. If v is nil, and no error hapens, the response is returned as is.
// If rate limit is exceeded and reset time is in the future, Do returns
// *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it
// is canceled or times out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

// ListOptions specifies the optional parameters to various List methods that
// support offset pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
