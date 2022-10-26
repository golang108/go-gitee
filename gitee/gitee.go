package gitee

import (
	"net/http"
	"net/url"
	"sync"
)

const (
	Version          = "v1.0.0"
	defaultBaseURL   = "https://gitee.com/api/v5/"
	defaultUserAgent = "go-gitee" + "/" + Version
)

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
	Users *UsersService
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

	return c

}
