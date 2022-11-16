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

type Hook struct {
	ID                  *int64     `json:"id,omitempty"`
	URL                 *string    `json:"url,omitempty"`
	CreatedAt           *Timestamp `json:"created_at,omitempty"`
	Password            *string    `json:"password,omitempty"`
	ProjectID           *int64     `json:"project_id,omitempty"`
	Result              *string    `json:"result,omitempty"`
	ResultCode          *int       `json:"result_code,omitempty"`
	PushEvents          *bool      `json:"push_events,omitempty"`
	TagPushEvents       *bool      `json:"tag_push_events,omitempty"`
	IssuesEvents        *bool      `json:"issues_events,omitempty"`
	NoteEvents          *bool      `json:"note_events,omitempty"`
	MergeRequestsEvents *bool      `json:"merge_requests_events,omitempty"`
}

func (h Hook) String() string {
	return Stringify(h)
}

// ListHooks lists all Hooks for the specified repository.
//
func (s *RepositoriesService) ListHooks(ctx context.Context, owner, repo string, opts *ListOptions) ([]*Hook, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/hooks", owner, repo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var hooks []*Hook
	resp, err := s.client.Do(ctx, req, &hooks)
	if err != nil {
		return nil, resp, err
	}

	return hooks, resp, nil
}
