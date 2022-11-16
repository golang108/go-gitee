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

// GitignoresService provides access to the gitignore related functions in the
// gitee API.
type GitignoresService service

// List all available Gitignore templates.
//
// 列出可使用的 .gitignore 模板 GET https://gitee.com/api/v5/gitignore/templates
func (s *GitignoresService) List(ctx context.Context) ([]string, *Response, error) {
	req, err := s.client.NewRequest("GET", "gitignore/templates", nil)
	if err != nil {
		return nil, nil, err
	}

	var availableTemplates []string
	resp, err := s.client.Do(ctx, req, &availableTemplates)
	if err != nil {
		return nil, resp, err
	}

	return availableTemplates, resp, nil
}

type Gitignore struct {
	License *string `json:"license,omitempty"`
	Source  *string `json:"source,omitempty"`
}

func (l Gitignore) String() string {
	return Stringify(l)
}

// Get extended metadata for one gitignore.
//
// 获取一个 .gitignore 模板 GET https://gitee.com/api/v5/gitignore/templates/{name}
func (s *GitignoresService) Get(ctx context.Context, licenseName string) (*Gitignore, *Response, error) {
	u := fmt.Sprintf("gitignore/templates/%s", licenseName)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ignore := new(Gitignore)
	resp, err := s.client.Do(ctx, req, ignore)
	if err != nil {
		return nil, resp, err
	}

	return ignore, resp, nil
}
