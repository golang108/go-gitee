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

import "context"

// MiscellaneousService 杂项
type MiscellaneousService service

// ListEmojis returns the emojis available to use on GitHub.
//
// 列出可使用的 Emoji GET https://gitee.com/api/v5/emojis
func (s *MiscellaneousService) ListEmojis(ctx context.Context) (map[string]string, *Response, error) {
	req, err := s.client.NewRequest("GET", "emojis", nil)
	if err != nil {
		return nil, nil, err
	}

	var emoji map[string]string
	resp, err := s.client.Do(ctx, req, &emoji)
	if err != nil {
		return nil, resp, err
	}

	return emoji, resp, nil
}
