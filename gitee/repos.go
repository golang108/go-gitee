package gitee

import (
	"context"
	"fmt"
	"time"
)

// RepositoriesService handles communication with the repository related
// methods of the gitee API.
//
// gitee API docs: https://gitee.com/api/v5/repos
type RepositoriesService service

// Branch represents a repository branch
type Branch struct {
	Name          *string      `json:"name,omitempty"`
	Commit        *BasicCommit `json:"commit,omitempty"` // 这里只有 sha，和url 2个属性
	Protected     *bool        `json:"protected,omitempty"`
	ProtectionUrl *string      `json:"protection_url,omitempty"`
}

func (r Branch) String() string {
	return Stringify(r)
}

// RepositoryCommit represents a commit in a repo.
type RepositoryCommit struct {
	URL         *string `json:"url,omitempty"`
	SHA         *string `json:"sha,omitempty"`
	HTMLURL     *string `json:"html_url,omitempty"`
	CommentsURL *string `json:"comments_url,omitempty"`

	Commit *Commit `json:"commit,omitempty"` // 这个里面 反而没有 sha 和 url

	Author    *User `json:"author,omitempty"`
	Committer *User `json:"committer,omitempty"`

	Parents []*BasicCommit `json:"parents,omitempty"` // 这里只有 sha，和url 2个属性

	// Details about how many changes were made in this commit. Only filled in during GetCommit!
	Stats *CommitStats `json:"stats,omitempty"`
	// Details about which files, and how this commit touched. Only filled in during GetCommit!
	Files []*CommitFile `json:"files,omitempty"`
}

// 仓库的某个提交
type BasicCommit struct {
	SHA *string `json:"sha,omitempty"`
	URL *string `json:"url,omitempty"`
}
type Commit struct {
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
	Message   *string       `json:"message,omitempty"`
	Tree      *Tree         `json:"tree,omitempty"`
}

type Tree struct {
	URL *string `json:"url,omitempty"`
	SHA *string `json:"sha,omitempty"`
}

type CommitAuthor struct {
	Date  *time.Time `json:"date,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`
}

type CommitStats struct {
	ID        *string `json:"id,omitempty"`
	Additions *int    `json:"additions,omitempty"`
	Deletions *int    `json:"deletions,omitempty"`
	Total     *int    `json:"total,omitempty"`
}

type CommitFile struct {
	SHA      *string `json:"sha,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Status   *string `json:"status,omitempty"`

	Additions *int `json:"additions,omitempty"`
	Deletions *int `json:"deletions,omitempty"`
	Changes   *int `json:"changes,omitempty"`

	BlobURL     *string `json:"blob_url,omitempty"`
	RawURL      *string `json:"raw_url,omitempty"`
	ContentsURL *string `json:"contents_url,omitempty"`

	Patch *string `json:"patch,omitempty"`
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

func (c Commit) String() string {
	return Stringify(c)
}

// GetCommit fetches the Commit object for a given SHA.
// 仓库的某个提交: GET https://gitee.com/api/v5/repos/{owner}/{repo}/commits/{sha}
func (s *RepositoriesService) GetCommit(ctx context.Context, owner string, repo string, sha string) (*RepositoryCommit, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/commits/%v", owner, repo, sha)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := new(RepositoryCommit)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}
