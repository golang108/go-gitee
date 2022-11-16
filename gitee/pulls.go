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

// PullRequestsService handles communication with the pull request related
// methods of the gitee API.
type PullRequestsService service

// PullRequest represents a GitHub pull request on a repository.
type PullRequest struct {
	ID     *int64  `json:"id,omitempty"`
	Number *int    `json:"number,omitempty"`
	State  *string `json:"state,omitempty"`
	Locked *bool   `json:"locked,omitempty"`
	Title  *string `json:"title,omitempty"`
}

func (p PullRequest) String() string {
	return Stringify(p)
}

// PullRequestListOptions specifies the optional parameters to the
// PullRequestsService.List method.
type PullRequestListOptions struct {
	// State filters pull requests based on their state. Possible values are:
	// open, closed, all. Default is "open". 可选。Pull Request 状态
	State string `url:"state,omitempty"`

	// Head filters pull requests by head user and branch name in the format of:
	// "user:ref-name". 可选。Pull Request 提交的源分支。格式：branch 或者：username:branch
	Head string `url:"head,omitempty"`

	// Base filters pull requests by base branch name. 可选。Pull Request 提交目标分支的名称。
	Base string `url:"base,omitempty"`

	// Sort specifies how to sort pull requests. Possible values are: created,
	// updated, popularity, long-running. Default is "created". 可选。排序字段，默认按创建时间
	Sort string `url:"sort,omitempty"`

	Since string `url:"since,omitempty"` //可选。起始的更新时间，要求时间格式为 ISO 8601

	// Direction in which to sort pull requests. Possible values are: asc, desc.
	// If Sort is "created" or not specified, Default is "desc", otherwise Default
	// is "asc" 可选。升序/降序
	Direction string `url:"direction,omitempty"`

	MilestoneNumber int64 `url:"milestone_number,omitempty"` //可选。里程碑序号(id)

	Labels string `url:"labels,omitempty"` //用逗号分开的标签。如: bug,performance

	ListOptions
}

// List the pull requests for the specified repository.
// 获取Pull Request列表 GET https://gitee.com/api/v5/repos/{owner}/{repo}/pulls
func (s *PullRequestsService) List(ctx context.Context, owner string, repo string, opts *PullRequestListOptions) ([]*PullRequest, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/pulls", owner, repo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var pulls []*PullRequest
	resp, err := s.client.Do(ctx, req, &pulls)
	if err != nil {
		return nil, resp, err
	}

	return pulls, resp, nil
}
