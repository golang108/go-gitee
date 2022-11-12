package gitee

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// RepositoriesService handles communication with the repository related
// methods of the gitee API.
//
// gitee API docs: https://gitee.com/api/v5/repos
type RepositoriesService service

type Links struct {
	Self *string `json:"self,omitempty"`
	HTML *string `json:"html,omitempty"`
}

type Permission struct {
	Pull  *bool `json:"pull,omitempty"`
	Push  *bool `json:"push,omitempty"`
	Admin *bool `json:"admin,omitempty"`
}

func (p Permission) String() string {
	return Stringify(p)
}

type Enterprise struct { // TODO "enterprise": null
}

type Program struct { // TODO "Program": []
}

type Repository struct {
	ID                  *int64      `json:"id,omitempty"`                    //"id": integer
	FullName            *string     `json:"full_name,omitempty"`             //"full_name": string
	HumanName           *string     `json:"human_name,omitempty"`            //"human_name": string
	URL                 *string     `json:"url,omitempty"`                   //"url": string
	Namespace           *Namespace  `json:"namespace,omitempty"`             //"namespace": 5 properties
	Path                *string     `json:"path,omitempty"`                  //"path": string 仓库路径
	Name                *string     `json:"name,omitempty"`                  //"name": string 仓库名称
	Owner               *User       `json:"owner,omitempty"`                 //"owner": 18 properties
	Aassigner           *User       `json:"assigner,omitempty"`              //"assigner": 18 properties
	Description         *string     `json:"description,omitempty"`           //"description": string 仓库描述
	Private             *bool       `json:"private,omitempty"`               //"private": boolean 是否私有
	Public              *bool       `json:"public,omitempty"`                //"public": boolean 是否公开
	Internal            *bool       `json:"internal,omitempty"`              //"internal": string 是否内部开源
	Fork                *bool       `json:"fork,omitempty"`                  //"fork": boolean 是否是fork仓库
	HTMLURL             *string     `json:"html_url,omitempty"`              //"html_url": string
	SSHURL              *string     `json:"ssh_url,omitempty"`               //"ssh_url": string
	ForksURL            *string     `json:"forks_url,omitempty"`             //"forks_url": string
	KeysURL             *string     `json:"keys_url,omitempty"`              //"keys_url": string
	CollaboratorsURL    *string     `json:"collaborators_url,omitempty"`     //"collaborators_url": string
	HooksURL            *string     `json:"hooks_url,omitempty"`             //"hooks_url": string
	BranchesURL         *string     `json:"branches_url,omitempty"`          //"branches_url": string
	TagsURL             *string     `json:"tags_url,omitempty"`              //"tags_url": string
	BlobsURL            *string     `json:"blobs_url,omitempty"`             //"blobs_url": string
	StargazersURL       *string     `json:"stargazers_url,omitempty"`        //"stargazers_url": string
	ContributorsURL     *string     `json:"contributors_url,omitempty"`      //"contributors_url": string
	CommitsURL          *string     `json:"commits_url,omitempty"`           //"commits_url": string
	CommentsURL         *string     `json:"comments_url,omitempty"`          //"comments_url": string
	IssueCommentURL     *string     `json:"issue_comment_url,omitempty"`     //"issue_comment_url": string
	IssueURL            *string     `json:"issues_url,omitempty"`            //"issues_url": string
	PullsURL            *string     `json:"pulls_url,omitempty"`             //"pulls_url": string
	MilestonesURL       *string     `json:"milestones_url,omitempty"`        //"milestones_url": string
	NotificationsURL    *string     `json:"notifications_url,omitempty"`     //"notifications_url": string
	LabelsURL           *string     `json:"labels_url,omitempty"`            //"labels_url": string
	ReleasesURL         *string     `json:"releases_url,omitempty"`          //"releases_url": string
	Recommend           *bool       `json:"recommend,omitempty"`             //"recommend": boolean 是否是推荐仓库
	GVP                 *bool       `json:"gvp,omitempty"`                   //"gvp": boolean 是否是 GVP 仓库
	Homepage            *string     `json:"homepage,omitempty"`              //"homepage": string 主页
	Language            *string     `json:"language,omitempty"`              //"language": string 语言
	ForksCount          *int        `json:"forks_count,omitempty"`           //"forks_count": integer 仓库fork数量
	StargazersCount     *int        `json:"stargazers_count,omitempty"`      //"stargazers_count": integer 仓库star数量
	WatchersCount       *int        `json:"watchers_count,omitempty"`        //"watchers_count": integer 仓库watch数量
	DefaultBranch       *string     `json:"default_branch,omitempty"`        //"default_branch": string 默认分支
	OpenIssuesCount     *int        `json:"open_issues_count,omitempty"`     //"open_issues_count": integer 开启的issue数量
	HasIssues           *bool       `json:"has_issues,omitempty"`            //"has_issues": boolean 是否开启issue功能
	HasWiki             *bool       `json:"has_wiki,omitempty"`              //"has_wiki": boolean 是否开启Wiki功能
	IssueComment        *bool       `json:"issue_comment,omitempty"`         //"issue_comment": boolean 是否允许用户对“关闭”状态的 Issue 进行评论
	CanComment          *bool       `json:"can_comment,omitempty"`           //"can_comment": boolean 是否允许用户对仓库进行评论
	PullRequestsEnabled *bool       `json:"pull_requests_enabled,omitempty"` //"pull_requests_enabled": boolean 是否接受 Pull Request，协作开发
	HasPage             *bool       `json:"has_page,omitempty"`              //"has_page": boolean 是否开启了 Pages
	License             *string     `json:"license,omitempty"`               //"license": string 开源许可
	Outsourced          *bool       `json:"outsourced,omitempty"`            //"outsourced": boolean 仓库类型（内部/外包）
	ProjectCreator      *string     `json:"project_creator,omitempty"`       //"project_creator": string 仓库创建者的 username
	Members             []*string   `json:"members,omitempty"`               //"members": Array[String] 仓库成员的username
	PushedAt            *Timestamp  `json:"pushed_at,omitempty"`             //"pushed_at": string 最近一次代码推送时间
	CreatedAt           *Timestamp  `json:"created_at,omitempty"`            //"created_at": string
	UpdatedAt           *Timestamp  `json:"updated_at,omitempty"`            //"updated_at": string
	Parent              *Repository `json:"parent,omitempty"`                //"parent": 69 properties
	Paas                *string     `json:"paas,omitempty"`                  //"paas": string
	Stared              *bool       `json:"stared,omitempty"`                //"stared": boolean 是否 star
	Watched             *bool       `json:"watched,omitempty"`               //"watched": boolean 是否 watch
	Permission          *Permission `json:"permission,omitempty"`            //"permission": Object 操作权限
	Relation            *string     `json:"relation,omitempty"`              //"relation": string 当前用户相对于仓库的角色
	AssigneesNumber     *int        `json:"assignees_number,omitempty"`      //"assignees_number": integer 代码审查设置，审查人数
	TestersNumber       *int        `json:"testers_number,omitempty"`        //"testers_number": integer 代码审查设置，测试人数
	Assignee            []*User     `json:"assignee,omitempty"`              //"assignee": 1 item
	Testers             []*User     `json:"testers,omitempty"`               //"testers": 1 item
	Status              *string     `json:"status,omitempty"`                //"status": string 仓库状态
	Programs            []Program   `json:"programs,omitempty"`              //"programs": 5 properties
	Enterprise          *Enterprise `json:"enterprise,omitempty"`            //"enterprise": 5 properties
	ProjectLabels       []*string   `json:"project_labels,omitempty"`        //"project_labels": 3 properties
}

func (r Repository) String() string {
	return Stringify(r)
}

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

type BranchRequest struct { // 暂时命名 为 不带 Create 前缀的吧
	Refs       *string `json:"refs,omitempty"`        // 起点名称, 默认：master
	BranchName *string `json:"branch_name,omitempty"` // 新创建的分支名称
}

// 创建分支 POST https://gitee.com/api/v5/repos/{owner}/{repo}/branches
func (s *RepositoriesService) CreateBranch(ctx context.Context, owner string, repo string, rreq *BranchRequest) (*Branch, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches", owner, repo)
	req, err := s.client.NewRequest("POST", u, rreq)
	if err != nil {
		return nil, nil, err
	}

	b := new(Branch)
	resp, err := s.client.Do(ctx, req, b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, nil
}

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

type ProtectionSetting struct {
	ID        *int64  `json:"id,omitempty"`
	ProjectID *int64  `json:"project_id,omitempty"`
	Wildcard  *string `json:"wildcard,omitempty"` // wildcard name
	Strict    *bool   `json:"strict,omitempty"`   // 是否严格检查
	Pusher    *string `json:"pusher,omitempty"`   // admin: 仓库管理员, none: 禁止任何人合并; 用户: 个人的地址path(多个用户用 ';' 隔开)
	Merger    *string `json:"merger,omitempty"`   // admin: 仓库管理员, none: 禁止任何人合并; 用户: 个人的地址path(多个用户用 ';' 隔开)
	//  contexts 还有个这个鸟属性，着文档写的真TMD垃圾
}

func (p ProtectionSetting) String() string {
	return Stringify(p)
}

// Protection represents a repository branch's protection.
type Protection struct {
	Name          *string           `json:"name,omitempty"` // branch name
	Commit        *RepositoryCommit `json:"commit,omitempty"`
	Protected     *bool             `json:"protected,omitempty"`
	ProtectionURL *string           `json:"protection_url,omitempty"`
	Links         *Links            `json:"_links,omitempty"`
}

func (p Protection) String() string {
	return Stringify(p)
}

type ProtectionRequest struct {
	Wildcard    *string `json:"wildcard,omitempty"`     // 分支/通配符
	NewWildcard *string `json:"new_wildcard,omitempty"` // 新分支/通配符(为空不修改)
	Pusher      *string `json:"pusher,omitempty"`       // admin: 仓库管理员, none: 禁止任何人合并; 用户: 个人的地址path(多个用户用 ';' 隔开)
	Merger      *string `json:"merger,omitempty"`       // admin: 仓库管理员, none: 禁止任何人合并; 用户: 个人的地址path(多个用户用 ';' 隔开)
}

// UpdateBranchProtection updates the protection of a given branch.
// 这个接口 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{branch}/protection 设置的 保护
//       和 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/branches/setting/new 设置的保护不是一个东西，
// owner 仓库所属空间地址(企业、组织或个人的地址path)
// repo 仓库路径(path)
// branch 分支名称
//     设置分支保护 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{branch}/protection
func (s *RepositoriesService) UpdateBranchProtection(ctx context.Context, owner, repo, branch string) (*Protection, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches/%v/protection", owner, repo, branch)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, nil, err
	}

	p := new(Protection)
	resp, err := s.client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// UpdateBranchWildcardProtection updates the protection of a given branch. 这个接口是更改已有的 分支保护测试，
// owner 仓库所属空间地址(企业、组织或个人的地址path)
// repo 仓库路径(path)
// wildcard 分支/通配符
// preq 分支保护策略设置 结构体，里面用到 new_wildcard，pusher，merger 字段 和 CreateBranchWildcardProtection 公用一个
// wildcard 设置分支/通配符.  感觉就是这个分支规则 起的一个名称,附带着通配的作用.
// new_wildcard 这次操作不是更新吗，如果填写这个字段，就是重新起个名称,附带着通配的作用.
// 分支保护策略设置 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{wildcard}/setting
func (s *RepositoriesService) UpdateBranchWildcardProtection(ctx context.Context, owner, repo, wildcard string,
	preq *ProtectionRequest) (*ProtectionSetting, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches/%v/setting", owner, repo, wildcard)
	req, err := s.client.NewRequest("PUT", u, preq)
	if err != nil {
		return nil, nil, err
	}

	p := new(ProtectionSetting) // 返回值不一样 ProtectionSetting 和 Protection
	resp, err := s.client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// RemoveBranchProtection removes the protection of a given branch.
//
//  取消保护分支的设置 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{branch}/protection
func (s *RepositoriesService) RemoveBranchProtection(ctx context.Context, owner, repo, branch string) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches/%v/protection", owner, repo, branch)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// RemoveBranchWildcardProtection removes the protection of a given wildcard.
//  删除仓库保护分支策略 DELETE https://gitee.com/api/v5/repos/{owner}/{repo}/branches/{wildcard}/setting
func (s *RepositoriesService) RemoveBranchWildcardProtection(ctx context.Context, owner, repo, wildcard string) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches/%v/setting", owner, repo, wildcard)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// CreateBranchWildcardProtection create the protection of a given branch
// owner 仓库所属空间地址(企业、组织或个人的地址path), 这个接口是新建一个 新的
// repo  仓库路径(path)
// preq   结构体，里面用到 wildcard，pusher，merger 字段.
// wildcard 设置分支/通配符.  感觉就是这个分支规则 起的一个名称,附带着通配的作用.
// 例如：设置为“master”，则对名称为“master”的分支生效；设置为“*-stable“ 或 ”release*“，则对名称符合此通配符的所有保护分支生效
// 新建仓库保护分支策略 PUT https://gitee.com/api/v5/repos/{owner}/{repo}/branches/setting/new
func (s *RepositoriesService) CreateBranchWildcardProtection(ctx context.Context, owner, repo string,
	preq *ProtectionRequest) (*ProtectionSetting, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/branches/setting/new", owner, repo)
	req, err := s.client.NewRequest("PUT", u, preq)
	if err != nil {
		return nil, nil, err
	}

	p := new(ProtectionSetting)
	resp, err := s.client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// RepositoryComment represents a comment for a commit, file, or line in a repository.
type RepositoryComment struct {
	ID          *int64         `json:"id,omitempty"`
	InReplyToID *int64         `json:"in_reply_to_id,omitempty"`
	Body        *string        `json:"body"`
	Source      *string        `json:"source,omitempty"`
	User        *User          `json:"user,omitempty"` // User-mutable fields
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	Target      *CommentTarget `json:"target,omitempty"` // TODO这个获取的都是null，不知道是干什么用的
}

type CommentTarget struct {
	Issue       *string `json:"issue,omitempty"`
	PullRequest *string `json:"pull_request,omitempty"`
}

func (r RepositoryComment) String() string {
	return Stringify(r)
}

type CommentsListOptions struct {
	Order string `url:"order,omitempty"` //排序顺序: asc(default),desc

	ListOptions //当前的页码, 每页的数量，最大为 100
}

// ListComments lists all the comments for the repository.
//
// 获取仓库的Commit评论 GET https://gitee.com/api/v5/repos/{owner}/{repo}/comments
func (s *RepositoriesService) ListComments(ctx context.Context, owner, repo string, opts *CommentsListOptions) ([]*RepositoryComment, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/comments", owner, repo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var comments []*RepositoryComment
	resp, err := s.client.Do(ctx, req, &comments)
	if err != nil {
		return nil, resp, err
	}

	return comments, resp, nil
}

// ListCommitComments lists all the comments for a given commit SHA.
// owner 仓库所属空间地址(企业、组织或个人的地址path)
// repo 仓库路径(path)
// ref* Commit的Reference
// 获取单个Commit的评论 GET https://gitee.com/api/v5/repos/{owner}/{repo}/commits/{ref}/comments
func (s *RepositoriesService) ListCommitComments(ctx context.Context, owner, repo, ref string, opts *ListOptions) ([]*RepositoryComment, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/commits/%v/comments", owner, repo, ref)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var comments []*RepositoryComment
	resp, err := s.client.Do(ctx, req, &comments)
	if err != nil {
		return nil, resp, err
	}

	return comments, resp, nil
}

// GetComment gets a single comment from a repository.
// 获取仓库的某条Commit评论 GET https://gitee.com/api/v5/repos/{owner}/{repo}/comments/{id}
func (s *RepositoriesService) GetComment(ctx context.Context, owner, repo string, id int64) (*RepositoryComment, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/comments/%v", owner, repo, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := new(RepositoryComment)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

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

type CommitsComparison struct {
	BaseCommit      *RepositoryCommit   `json:"base_commit,omitempty"`
	MergeBaseCommit *RepositoryCommit   `json:"merge_base_commit,omitempty"`
	Files           []*CommitFile       `json:"files,omitempty"`
	Commits         []*RepositoryCommit `json:"commits,omitempty"`
}

func (c CommitsComparison) String() string {
	return Stringify(c)
}

// CompareCommits compares a range of commits with each other.
//
// 两个Commits之间对比的版本差异 GET https://gitee.com/api/v5/repos/{owner}/{repo}/compare/{base}...{head}
func (s *RepositoriesService) CompareCommits(ctx context.Context, owner, repo string, base, head string) (*CommitsComparison, *Response, error) {
	escapedBase := url.QueryEscape(base)
	escapedHead := url.QueryEscape(head)

	u := fmt.Sprintf("repos/%v/%v/compare/%v...%v", owner, repo, escapedBase, escapedHead)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	comp := new(CommitsComparison)
	resp, err := s.client.Do(ctx, req, comp)
	if err != nil {
		return nil, resp, err
	}

	return comp, resp, nil
}

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

type RepositoryCreateOptions struct {
	Name              *string `json:"name,omitempty"`               // 仓库名称
	Path              *string `json:"path,omitempty"`               // 仓库路径
	Description       *string `json:"description,omitempty"`        //仓库描述
	Homepage          *string `json:"homepage,omitempty"`           //主页(eg: https://gitee.com) 一个有效的http链接
	GitignoreTemplate *string `json:"gitignore_template,omitempty"` //Git Ignore模版
	LicenseTemplate   *string `json:"license_template,omitempty"`   // License模版
	Private           *bool   `json:"private,omitempty"`            //目前仅支持私有
	HasIssues         *bool   `json:"has_issues,omitempty"`         //允许提Issue与否。默认: 允许(true)
	HasWiki           *bool   `json:"has_wiki,omitempty"`           //提供Wiki与否。默认: 提供(true)
	CanComment        *bool   `json:"can_comment,omitempty"`        //允许用户对仓库进行评论。默认： 允许(true)
	AutoInit          *bool   `json:"auto_init,omitempty"`          //值为true时则会用README初始化仓库。默认: 不初始化(false)

	//创建企业仓库 POST https://gitee.com/api/v5/enterprises/{enterprise}/repos
	Enterprise     *string `json:"enterprise,omitempty"`      //企业的路径(path/login)  # 必填项
	Outsourced     *bool   `json:"outsourced,omitempty"`      //值为true值为外包仓库, false值为内部仓库。默认: 内部仓库(false)
	ProjectCreator *string `json:"project_creator,omitempty"` //负责人的username
	Members        *string `json:"members,omitempty"`         //用逗号分开的仓库成员。如: member1,member2
}

// Create a new repository. If an organization is specified, the new
// repository will be created under that org. If the empty string is
// specified, it will be created for the authenticated user.
// 创建一个仓库 POST https://gitee.com/api/v5/user/repos
// 创建组织仓库 POST https://gitee.com/api/v5/orgs/{org}/repos
// 创建企业仓库 POST https://gitee.com/api/v5/enterprises/{enterprise}/repos
func (s *RepositoriesService) Create(ctx context.Context, orgOrenterprise string, opt *RepositoryCreateOptions) (*Repository, *Response, error) {
	var u string
	if orgOrenterprise != "" {
		u = fmt.Sprintf("orgs/%v/repos", orgOrenterprise)
		if opt.Enterprise != nil {
			u = fmt.Sprintf("enterprises/%v/repos", orgOrenterprise)
		}
	} else {
		u = "user/repos"
	}

	req, err := s.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	r := new(Repository)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

type RepositoryListOptions struct {
	// Visibility of repositories to list. Can be one of all, public, or private.
	// Default: all  公开(public)、私有(private)或者所有(all)，默认: 所有(all)
	Visibility string `url:"visibility,omitempty"`

	// List repos of given affiliation[s].
	// Comma-separated list of values. Can include:
	// * owner: Repositories that are owned by the authenticated user.
	// * collaborator: Repositories that the user has been added to as a
	//   collaborator.
	// * organization_member: Repositories that the user has access to through
	//   being a member of an organization. This includes every repository on
	//   every team that the user is on.
	// Default: owner,collaborator,organization_member
	// owner(授权用户拥有的仓库)、collaborator(授权用户为仓库成员)、
	// organization_member(授权用户为仓库所在组织并有访问仓库权限)、
	// enterprise_member(授权用户所在企业并有访问仓库权限)、
	// admin(所有有权限的，包括所管理的组织中所有仓库、所管理的企业的所有仓库)。
	// 可以用逗号分隔符组合。如: owner, organization_member 或 owner, collaborator, organization_member
	Affiliation string `url:"affiliation,omitempty"`

	// Type of repositories to list.
	// Can be one of all, owner, public, private, member. Default: all
	// Will cause a 422 error if used in the same request as visibility or
	// affiliation.
	// 筛选用户仓库: 其创建(owner)、个人(personal)、其为成员(member)、公开(public)、私有(private)，
	//不能与 visibility 或 affiliation 参数一并使用，否则会报 422 错误
	Type string `url:"type,omitempty"`

	// How to sort the repository list. Can be one of created, updated, pushed,
	// full_name. Default: full_name
	// 排序方式: 创建时间(created)，更新时间(updated)，最后推送时间(pushed)，
	// 仓库所属与名称(full_name)。默认: full_name
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort repositories. Can be one of asc or desc.
	// Default: when using full_name: asc; otherwise desc
	// 如果sort参数为full_name，用升序(asc)。否则降序(desc)
	Direction string `url:"direction,omitempty"`

	// 搜索关键字
	Q string `url:"q,omitempty"`

	ListOptions
}

// List the repositories for a user. Passing the empty string will list
// repositories for the authenticated user.
// 列出授权用户的所有仓库 GET https://gitee.com/api/v5/user/repos
// 获取某个用户的公开仓库 GET https://gitee.com/api/v5/users/{username}/repos
func (s *RepositoriesService) List(ctx context.Context, user string, opts *RepositoryListOptions) ([]*Repository, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/repos", user)
	} else {
		u = "user/repos"
	}

	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var repos []*Repository
	resp, err := s.client.Do(ctx, req, &repos)
	if err != nil {
		return nil, resp, err
	}

	return repos, resp, nil

}

// 获取一个组织的仓库 GET https://gitee.com/api/v5/orgs/{org}/repos
func (s *RepositoriesService) ListOrgs(ctx context.Context, org string, opts *RepositoryListOptions) ([]*Repository, *Response, error) {
	var u string
	if org != "" {
		u = fmt.Sprintf("orgs/%v/repos", org)
	} else {
		return nil, nil, fmt.Errorf("org is empty")
	}

	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var repos []*Repository
	resp, err := s.client.Do(ctx, req, &repos)
	if err != nil {
		return nil, resp, err
	}

	return repos, resp, nil
}

// 获取企业的所有仓库 GET https://gitee.com/api/v5/enterprises/{enterprise}/repos
func (s *RepositoriesService) ListEnterprises(ctx context.Context, enterprise string, opts *RepositoryListOptions) ([]*Repository, *Response, error) {
	var u string
	if enterprise != "" {
		u = fmt.Sprintf("enterprises/%v/repos", enterprise)
	} else {
		return nil, nil, fmt.Errorf("enterprise is empty")
	}

	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var repos []*Repository
	resp, err := s.client.Do(ctx, req, &repos)
	if err != nil {
		return nil, resp, err
	}

	return repos, resp, nil
}
