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
	"bytes"
	"context"
)

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

type MarkdownRequest struct {
	Text *string `json:"text,omitempty"`
}

// Markdown renders an arbitrary Markdown document.
//
// 渲染 Markdown 文本 POST https://gitee.com/api/v5/markdown
func (s *MiscellaneousService) Markdown(ctx context.Context, text string, opts *MarkdownRequest) (string, *Response, error) {
	req, err := s.client.NewRequest("POST", "markdown", opts)
	if err != nil {
		return "", nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(ctx, req, buf)
	if err != nil {
		return "", resp, err
	}

	return buf.String(), resp, nil
}

// UserEmail represents user's email address
type UserEmail struct {
	Email *string `json:"email,omitempty"`
}

// GetEmails lists all email addresses for the authenticated user.
//
//  获取授权用户的全部邮箱 GET https://gitee.com/api/v5/emails
func (s *MiscellaneousService) GetEmail(ctx context.Context, opts *ListOptions) (*UserEmail, *Response, error) {
	u := "emails"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	emails := new(UserEmail)
	resp, err := s.client.Do(ctx, req, &emails)
	if err != nil {
		return nil, resp, err
	}

	return emails, resp, nil
}
