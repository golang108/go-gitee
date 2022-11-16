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

package gitee

import (
	"context"
	"fmt"
)

// LicensesService handles communication with the license related
// methods of the gitee API.
type LicensesService service

// RepositoryLicense represents the license for a repository.
type RepositoryLicense struct {
	License *string `json:"license,omitempty"`
}

func (l RepositoryLicense) String() string {
	return Stringify(l)
}

// License gets the contents of a repository's license if one is detected.
// 这个接口github放到repos 里面了
// 获取一个仓库使用的开源许可协议 GET https://gitee.com/api/v5/repos/{owner}/{repo}/license
func (s *LicensesService) License(ctx context.Context, owner, repo string) (*RepositoryLicense, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/license", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(RepositoryLicense)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// List popular open source licenses.
//
// 列出可使用的开源许可协议 GET https://gitee.com/api/v5/licenses
func (s *LicensesService) List(ctx context.Context) ([]string, *Response, error) {
	req, err := s.client.NewRequest("GET", "licenses", nil)
	if err != nil {
		return nil, nil, err
	}

	var licenses []string
	resp, err := s.client.Do(ctx, req, &licenses)
	if err != nil {
		return nil, resp, err
	}

	return licenses, resp, nil
}
