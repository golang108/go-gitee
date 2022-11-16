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

// ActivityService handles communication with the activity related
// methods of the gitee API.
type ActivityService service

// ListStargazers lists people who have starred the specified repo.
// 列出 star 了仓库的用户 GET https://gitee.com/api/v5/repos/{owner}/{repo}/stargazers
func (s *ActivityService) ListStargazers(ctx context.Context, owner, repo string, opts *ListOptions) ([]*User, *Response, error) {
	u := fmt.Sprintf("repos/%s/%s/stargazers", owner, repo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var stargazers []*User
	resp, err := s.client.Do(ctx, req, &stargazers)
	if err != nil {
		return nil, resp, err
	}

	return stargazers, resp, nil
}

// ListWatchers lists watchers of a particular repo.
//
// 列出 watch 了仓库的用户 GET https://gitee.com/api/v5/repos/{owner}/{repo}/subscribers
func (s *ActivityService) ListWatchers(ctx context.Context, owner, repo string, opts *ListOptions) ([]*User, *Response, error) {
	u := fmt.Sprintf("repos/%s/%s/subscribers", owner, repo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var watchers []*User
	resp, err := s.client.Do(ctx, req, &watchers)
	if err != nil {
		return nil, resp, err
	}

	return watchers, resp, nil
}
