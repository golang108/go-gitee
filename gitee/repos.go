package gitee

import (
	"context"
	"fmt"
)

// RepositoriesService handles communication with the repository related
// methods of the gitee API.
//
// gitee API docs: https://gitee.com/api/v5/repos
type RepositoriesService service

// Branch represents a repository branch
type Branch struct {
	Name          *string           `json:"name,omitempty"`
	Commit        *RepositoryCommit `json:"commit,omitempty"`
	Protected     *bool             `json:"protected,omitempty"`
	ProtectionUrl *string           `json:"protection_url,omitempty"`
}

func (r Branch) String() string {
	return Stringify(r)
}

// RepositoryCommit represents a commit in a repo.
type RepositoryCommit struct {
	SHA *string `json:"sha,omitempty"`
	URL *string `json:"url,omitempty"`
}

// ListBranches lists branches for the specified repository.
// 获取所有分支: GET https://gitee.com/api/v5/repos/{owner}/{repo}/branches
func (s *RepositoriesService) ListBranches(ctx context.Context, owner string, repo string) ([]*Branch, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches", owner, repo)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var branches []*Branch
	resp, err := s.client.Do(ctx, req, &branches)
	if err != nil {
		return nil, resp, err
	}

	return branches, resp, nil
}
