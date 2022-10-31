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

// TODO 创建分支 POST https://gitee.com/api/v5/repos/{owner}/{repo}/branches

// GetBranch gets the specified branch for a repository.
// 获取单个分支 GET https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{branch}
func (s *RepositoriesService) GetBranch(ctx context.Context, owner, repo, branch string) (*Branch, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches/%v", owner, repo, branch)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var br = new(Branch)
	resp, err := s.client.Do(ctx, req, br)
	if err != nil {
		return nil, resp, err
	}

	return br, resp, nil
}

// TODO 设置分支保护 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{branch}/protection

// TODO 取消保护分支的设置 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{branch}/protection

// TODO 分支保护策略设置 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{wildcard}/setting

// TODO 删除仓库保护分支策略 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{wildcard}/setting

// TODO 新建仓库保护分支策略 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/branches/setting/new

// todo 获取仓库的Commit评论 GET https://gitee.com/api/v5/repos/{owner}/{repo}/comments

// todo 获取单个Commit的评论 GET https://gitee.com/api/v5/repos/{owner}/{repo}/commits/{ref}/comments

// todo 获取仓库的某条Commit评论 GET https://gitee.com/api/v5/repos/{owner}/{repo}/comments/{id}

// todo 更新Commit评论 PATCH https://gitee.com/api/v5/repos/{owner}/{repo}/comments/{id}

// todo 删除Commit评论 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/comments/{id}

// todo 创建Commit评论 POST https://gitee.com/api/v5/repos/{owner}/{repo}/commits/{sha}/comments

func (c Commit) String() string {
	return Stringify(c)
}

// CommitsListOptions specifies the optional parameters to the
// RepositoriesService.ListCommits method. 这个和 一样
type CommitsListOptions struct {
	// SHA or branch to start listing Commits from. 提交起始的SHA值或者分支名. 默认: 仓库的默认分支
	SHA string `url:"sha,omitempty"`

	// Path that should be touched by the returned Commits. 包含该文件的提交
	Path string `url:"path,omitempty"`

	// Author of by which to filter Commits. 提交作者的邮箱或个人空间地址(username/login)
	Author string `url:"author,omitempty"`

	// Since when should Commits be included in the response.提交的起始时间，时间格式为 ISO 8601
	Since time.Time `url:"since,omitempty"`

	// Until when should Commits be included in the response.提交的最后时间，时间格式为 ISO 8601
	Until time.Time `url:"until,omitempty"`

	ListOptions //当前的页码, 每页的数量，最大为 100
}

// ListCommits lists the commits of a repository.
// 仓库的所有提交 GET https://gitee.com/api/v5/repos/{owner}/{repo}/commits
func (s *RepositoriesService) ListCommits(ctx context.Context, owner, repo string, opts *CommitsListOptions) ([]*RepositoryCommit, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/commits", owner, repo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var commits []*RepositoryCommit
	resp, err := s.client.Do(ctx, req, &commits)
	if err != nil {
		return nil, resp, err
	}

	return commits, resp, nil
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

// todo 两个Commits之间对比的版本差异 GET https://gitee.com/api/v5/repos/{owner}/{repo}/compare/{base}...{head}

// TODO 获取仓库已部署的公钥 GET https://gitee.com/api/v5/repos/{owner}/{repo}/keys

// TODO 为仓库添加公钥 POST https://gitee.com/api/v5/repos/{owner}/{repo}/keys

// TODO 获取仓库可部署的公钥 GET https://gitee.com/api/v5/repos/{owner}/{repo}/keys/available

// TODO 启用仓库公钥  PUT https://gitee.com/api/v5/repos/{owner}/{repo}/keys/enable/{id}

// TODO 停用仓库公钥 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/keys/enable/{id}

// TODO 获取仓库的单个公钥 GET https://gitee.com/api/v5/repos/{owner}/{repo}/keys/{id}

// TODO 删除一个仓库公钥 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/keys/{id}

// TODO 获取仓库README GET https://gitee.com/api/v5/repos/{owner}/{repo}/readme

// TODO 获取仓库具体路径下的内容 GET https://gitee.com/api/v5/repos/{owner}/{repo}/contents(/{path})

// TODO 新建文件 POST https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}

// TODO 更新文件 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}

// TODO 删除文件 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}

// TODO 获取Pages信息 GET https://gitee.com/api/v5/repos/{owner}/{repo}/pages

// TODO 上传设置 Pages SSL 证书和域名 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/pages

// TODO 请求建立Pages POST https://gitee.com/api/v5/repos/{owner}/{repo}/pages/builds

// TODO 获取用户的某个仓库 GET https://gitee.com/api/v5/repos/{owner}/{repo}

// TODO 更新仓库设置 PATCH https://gitee.com/api/v5/repos/{owner}/{repo}

// TODO 删除一个仓库 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}

// TODO 修改代码审查设置 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/reviewer

// TODO 获取仓库推送规则设置 GET https://gitee.com/api/v5/repos/{owner}/{repo}/push_config

// TODO 修改仓库推送规则设置 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/push_config

// TODO 获取仓库贡献者 GET https://gitee.com/api/v5/repos/{owner}/{repo}/contributors

// TODO 列出仓库所有的tags GET https://gitee.com/api/v5/repos/{owner}/{repo}/tags

// TODO 创建一个仓库的 Tag POST https://gitee.com/api/v5/repos/{owner}/{repo}/tags

// TODO 清空一个仓库 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/clear

// TODO 获取仓库的所有成员 GET https://gitee.com/api/v5/repos/{owner}/{repo}/collaborators

// TODO 判断用户是否为仓库成员 GET https://gitee.com/api/v5/repos/{owner}/{repo}/collaborators/{username}

// TODO 添加仓库成员 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/collaborators/{username}

// TODO 移除仓库成员 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/collaborators/{username}

// TODO 查看仓库成员的权限 GET https://gitee.com/api/v5/repos/{owner}/{repo}/collaborators/{username}/permission

// TODO 查看仓库的Forks GET https://gitee.com/api/v5/repos/{owner}/{repo}/forks

// TODO Fork一个仓库 POST https://gitee.com/api/v5/repos/{owner}/{repo}/forks

// TODO 获取仓库的百度统计 key GET https://gitee.com/api/v5/repos/{owner}/{repo}/baidu_statistic_key

// TODO 设置/更新仓库的百度统计 key POST https://gitee.com/api/v5/repos/{owner}/{repo}/baidu_statistic_key

// TODO 删除仓库的百度统计 key DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/baidu_statistic_key

// TODO 获取最近30天的七日以内访问量 POST https://gitee.com/api/v5/repos/{owner}/{repo}/traffic-data

// TODO 获取仓库的所有Releases GET https://gitee.com/api/v5/repos/{owner}/{repo}/releases

// TODO 创建仓库Release POST https://gitee.com/api/v5/repos/{owner}/{repo}/releases

// TODO 获取仓库的单个Releases GET https://gitee.com/api/v5/repos/{owner}/{repo}/releases/{id}

// TODO 更新仓库Release PATCH https://gitee.com/api/v5/repos/{owner}/{repo}/releases/{id}

// TODO 删除仓库Release DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/releases/{id}

// TODO 获取仓库的最后更新的Release GET https://gitee.com/api/v5/repos/{owner}/{repo}/releases/latest

// TODO 根据Tag名称获取仓库的Release GET https://gitee.com/api/v5/repos/{owner}/{repo}/releases/tags/{tag}

// TODO 开通Gitee Go POST https://gitee.com/api/v5/repos/{owner}/{repo}/open

// TODO 列出授权用户的所有仓库 GET https://gitee.com/api/v5/user/repos

// TODO 创建一个仓库 POST https://gitee.com/api/v5/user/repos

// TODO 获取某个用户的公开仓库 GET https://gitee.com/api/v5/users/{username}/repos

// TODO 获取一个组织的仓库 GET https://gitee.com/api/v5/orgs/{org}/repos

// TODO 创建组织仓库 POST https://gitee.com/api/v5/orgs/{org}/repos

// TODO 获取企业的所有仓库 GET https://gitee.com/api/v5/enterprises/{enterprise}/repos

// TODO 创建企业仓库 POST https://gitee.com/api/v5/enterprises/{enterprise}/repos
