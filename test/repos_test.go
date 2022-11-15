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

package test

import (
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"testing"
)

func TestListBranches(t *testing.T) {
	branches, response, err := client.Repositories.ListBranches(ctx, "mamh-java", "jenkins-jenkins")
	for index, br := range branches {
		fmt.Println(index, *br.Name, *br.Protected, *br.ProtectionUrl)
	}
	fmt.Println(response)
	fmt.Println(err)
}

func TestCreateBranch(t *testing.T) {
	rreq := &gitee.BranchRequest{
		Refs:       gitee.String("main"), // 从已有分支 创建新的分支
		BranchName: gitee.String("master"),
	}
	branch, response, err := client.Repositories.CreateBranch(ctx, "mamh-mixed", "go-gitee", rreq)
	fmt.Println(branch)
	fmt.Println(response)
	fmt.Println(err) // 分支名已存在
}

func TestCreateBranch1(t *testing.T) {
	rreq := &gitee.BranchRequest{
		Refs:       gitee.String("7a15f560525e17bc2b58e0b6c4bff6ba82e7a557"), // 从一个 commit id 创建新的分支
		BranchName: gitee.String("master"),
	}
	branch, response, err := client.Repositories.CreateBranch(ctx, "mamh-mixed", "go-gitee", rreq)
	fmt.Println(branch)
	fmt.Println(response)
	fmt.Println(err) // 分支名已存在
}

func TestGetBranch(t *testing.T) {
	branch, response, err := client.Repositories.GetBranch(ctx, "mamh-mixed", "go-gitee", "main")
	fmt.Println(branch)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetCommit(t *testing.T) {
	commit, response, err := client.Repositories.GetCommit(ctx, "mamh-mixed", "go-gitee", "8896821c53eda6698ef5c75ba5182e547e8476f1")

	fmt.Println(commit, response, err)

}

func TestListCommits(t *testing.T) {
	var opts = &gitee.CommitsListOptions{}
	for {
		commits, response, err := client.Repositories.ListCommits(ctx, "mamh-mixed", "go-gitee", opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, commit := range commits {
			fmt.Println(index, len(commits), *commit.Commit.Message, *commit.SHA)
		}
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}

}

func TestListComments(t *testing.T) {
	var opts = &gitee.CommentsListOptions{
		Order: "desc",
	}
	for {
		comments, response, err := client.Repositories.ListComments(ctx, "mamh-mixed", "go-gitee", opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, comment := range comments {
			fmt.Println(index, len(comments),
				*comment.ID,
				*comment.User.Name,
				*comment.CreatedAt,
				*comment.Body,
			)
		}
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
}

func TestListCommitComments(t *testing.T) {
	var opts = &gitee.ListOptions{}
	ref := "c764302e6da151e08608c08ab30e986b04b9064b"
	for {
		comments, response, err := client.Repositories.ListCommitComments(ctx, "mamh-mixed", "go-gitee", ref, opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, comment := range comments {
			fmt.Println(index, len(comments),
				*comment.ID,
				*comment.User.Name,
				*comment.CreatedAt,
				*comment.Body,
			)
		}
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
}

func TestGetComment(t *testing.T) {
	comment, response, err := client.Repositories.GetComment(ctx, "mamh-mixed", "go-gitee", 14339904)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response,
		*comment.ID,
		*comment.User.Name,
		*comment.CreatedAt,
		*comment.Body,
	)
}

func TestDeleteComment(t *testing.T) {
	response, err := client.Repositories.DeleteComment(ctx, "mamh-mixed", "go-gitee", 14339904)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}

func TestCreateComment(t *testing.T) {
	creq := &gitee.CommentRequest{
		Body:     gitee.String("body for comment test/repos_test.go, 4 hang"),
		Path:     gitee.String("test/repos_test.go"),
		Position: gitee.Int64(4),
	}
	ref := "c764302e6da151e08608c08ab30e986b04b9064b"

	comment, response, err := client.Repositories.CreateComment(ctx, "mamh-mixed", "go-gitee", ref, creq)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response,
		*comment.ID,
		*comment.User.Name,
		*comment.CreatedAt,
		*comment.Body,
	)

}

func TestUpdateComment(t *testing.T) {
	creq := &gitee.CommentRequest{
		Body: gitee.String("update for id 14340395 comment, \n 0 3 14340395 系统提示 2022-11-12 19:17:47 +0800 CST body for comment test/repos_test.go, 18 hang\n"),
	}
	id := int64(14340395)

	comment, response, err := client.Repositories.UpdateComment(ctx, "mamh-mixed", "go-gitee", id, creq)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response,
		*comment.ID,
		*comment.User.Name,
		*comment.CreatedAt,
		*comment.Body,
	)

}

func TestCreateKey(t *testing.T) {
	kreq := &gitee.KeyRequest{
		Key:   gitee.String("ssh-rsa"),
		Title: gitee.String("public key title"),
	}

	key, response, err := client.Repositories.CreateKey(ctx, "mamh-mixed", "go-gitee", kreq)
	if err != nil {
		fmt.Println(err) // 1. 指纹生成失败, 当前公钥是无效的. 2. 当前仓库已经启用此公钥
		return
	}
	fmt.Println(response)
	fmt.Println(key)
}

func TestListKeys(t *testing.T) {
	opts := &gitee.ListOptions{}
	keys, response, err := client.Repositories.ListKeys(ctx, "mamh-mixed", "go-gitee", opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
	fmt.Println(keys)

}

func TestListAvailableKeys(t *testing.T) {
	opts := &gitee.ListOptions{}
	keys, response, err := client.Repositories.ListAvailableKeys(ctx, "mamh-mixed", "go-gitee", opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
	fmt.Println(keys)

}

func TestEnableKey(t *testing.T) {
	id := int64(3585098)
	response, err := client.Repositories.EnableKey(ctx, "mamh-mixed", "go-gitee", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response) // 启动用成功是否 返回 都是  404 Not Found,  "message": "Deploy Key"
}

func TestDisableKey(t *testing.T) {
	id := int64(3585098)
	response, err := client.Repositories.DisableKey(ctx, "mamh-mixed", "go-gitee", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response) // 启动用成功是否 返回 都是  404 Not Found,  "message": "Deploy Key"
}

func TestGetKey(t *testing.T) {
	id := int64(3584973)
	key, response, err := client.Repositories.GetKey(ctx, "mamh-mixed", "go-gitee", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
	fmt.Println(key)

}

func TestDeleteKey(t *testing.T) {
	id := int64(3584973)
	response, err := client.Repositories.DeleteKey(ctx, "mamh-mixed", "go-gitee", id)
	if err != nil {
		fmt.Println(err) // 没有相关公钥
		return
	}
	fmt.Println(response)
}

func TestGetReadme(t *testing.T) {
	opts := &gitee.RepositoryContentGetOptions{
		Ref: "main", // 分支、tag或commit。默认: 仓库的默认分支(通常是master)
	}
	readme, response, err := client.Repositories.GetReadme(ctx, "mamh-mixed", "go-gitee", opts)
	if err != nil {
		fmt.Println(err) // 404 Not Found, map[], Commit
		return
	}
	fmt.Println(response)
	fmt.Println(readme)

}

func TestGetContents1(t *testing.T) {
	opts := &gitee.RepositoryContentGetOptions{
		Ref: "main", // 分支、tag或commit。默认: 仓库的默认分支(通常是master)
	}
	filepath := "gitee" // 获取仓库下面的一个目录的内容, 返回值 第一个就会是nil
	fileContent, directoryContent, response, err := client.Repositories.GetContents(ctx, "mamh-mixed", "go-gitee", filepath, opts)
	if err != nil {
		fmt.Println(err) // 404 Not Found, map[], Commit
		return
	}
	fmt.Println(response)
	fmt.Println(fileContent)
	fmt.Println(directoryContent)

}

func TestGetContents2(t *testing.T) {
	opts := &gitee.RepositoryContentGetOptions{
		Ref: "main", // 分支、tag或commit。默认: 仓库的默认分支(通常是master)
	}
	filepath := "gitee/repos.go" // 获取仓库下某个文件，这时候第二个参数就会是 []
	fileContent, directoryContent, response, err := client.Repositories.GetContents(ctx, "mamh-mixed", "go-gitee", filepath, opts)
	if err != nil {
		fmt.Println(err) // 404 Not Found, map[], Commit
		return
	}
	fmt.Println(response)
	fmt.Println(fileContent)
	fmt.Println(directoryContent)

}

func TestCreateFile(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc-test"
	path := "go.mod"

	opts := &gitee.RepositoryContentFileOptions{
		Message: gitee.String("go.mod: add go.mod file"),
		Content: []byte("bW9kdWxlIGdpdGh1Yi5jb20vbWFnZXNmYy9tYWdlc2ZjLXRlc3QKCmdvIDEuMTY="),
		Branch:  gitee.String("master"), // 分支名称。默认为仓库对默认分支
	}

	file, response, err := client.Repositories.CreateFile(ctx, owner, repo, path, opts)
	fmt.Println(file, response, err) // A file with this name already exists

}

func TestUpdateFile(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc-test"
	path := "go.mod"

	opts := &gitee.RepositoryContentFileOptions{
		Message: gitee.String("go.mod: update go.mod file"),
		Content: []byte("指定的文件 的=="),
		Branch:  gitee.String("master"),                                   // 分支名称。默认为仓库对默认分支
		SHA:     gitee.String("7a680025f5489e7e6279ee169384a958c7edac61"), // 指定的文件 的 sha 可以和这里的不一样.文件的 Blob SHA，可通过 [获取仓库具体路径下的内容] API 获取
	}

	file, response, err := client.Repositories.UpdateFile(ctx, owner, repo, path, opts)
	fmt.Println(file, response, err) // A file with this name already exists

}

func TestDeleteFile(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc-test"
	path := "go.mod"

	opts := &gitee.RepositoryContentFileOptions{
		Message: gitee.String("go.mod: delete go.mod file"),
		Branch:  gitee.String("master"),                                   // 分支名称。默认为仓库对默认分支
		SHA:     gitee.String("6b53779e48a646f8d24aecb4bc788c5237718c22"), // 可通过 [获取仓库具体路径下的内容] API 获取
	}

	file, response, err := client.Repositories.DeleteFile(ctx, owner, repo, path, opts)
	fmt.Println(file, response, err) // A file with this name already exists

}

func TestList(t *testing.T) {
	opts := &gitee.RepositoryListOptions{}
	repository, response, err := client.Repositories.List(ctx, "", opts)

	fmt.Println("Namespace", repository[0].Namespace)
	fmt.Println("Owner", repository[0].Owner)
	fmt.Println("Aassigner", repository[0].Aassigner)
	fmt.Println("Parent", repository[0].Parent)
	fmt.Println("Permission", repository[0].Permission)
	fmt.Println("Assignee", repository[0].Assignee[0])
	fmt.Println("Testers", repository[0].Testers[0])
	fmt.Println("Programs", repository[0].Programs)
	fmt.Println("Enterprise", repository[0].Enterprise)
	fmt.Println("ProjectLabels", repository[0].ProjectLabels)

	fmt.Println("response", response)
	fmt.Println("err", err)

}

func TestList1(t *testing.T) {
	opts := &gitee.RepositoryListOptions{}
	repository, response, err := client.Repositories.List(ctx, "elunez", opts)

	fmt.Println("repository", repository)

	fmt.Println("response", response)
	fmt.Println("err", err)

}

func TestListOrganizations(t *testing.T) {
	opts := &gitee.RepositoryListOptions{}
	repository, response, err := client.Repositories.ListOrganizations(ctx, "mamh-mixed", opts)

	fmt.Println("repository", repository)
	fmt.Println("response", response)
	fmt.Println("err", err)

}

func TestListEnterprises(t *testing.T) {
	opts := &gitee.RepositoryListOptions{}
	repository, response, err := client.Repositories.ListEnterprises(ctx, "magesfc", opts)

	fmt.Println("repository", repository)
	fmt.Println("response", response)
	fmt.Println("err", err)

}

func TestCompareCommits(t *testing.T) {
	owner := "mamh-mixed"
	repo := "go-gitee"
	base := "7a15f560525e17bc2b58e0b6c4bff6ba82e7a557"
	head := "d6b2bbde37dd77d7d9bc75363ff2b02ddad3ddaa"
	comp, response, err := client.Repositories.CompareCommits(ctx, owner, repo, base, head)
	for i, commit := range comp.Commits {
		fmt.Println(i, *commit.SHA)
	}
	fmt.Println("base: ", *comp.MergeBaseCommit.SHA)
	fmt.Println(response)
	fmt.Println(err)
}

func TestCreate(t *testing.T) {
	opts := &gitee.CreateRepositoryRequest{
		Name:        gitee.String("repo_name2"), // 仓库名称
		Path:        gitee.String("repo_Path2"), //路径 (请注意：仓库路径即仓库访问 URL 地址，更改仓库路径将导致原克隆地址不可用)
		Description: gitee.String("repo_Description"),
		Homepage:    gitee.String("xxxxxxxxx"),
	}
	repository, response, err := client.Repositories.Create(ctx, opts)
	fmt.Println("repository", repository)
	fmt.Println("response", response)
	fmt.Println("err", err)
}

func TestCreateOrgRepository(t *testing.T) {
	opts := &gitee.CreateOrgRepositoryRequest{
		CreateRepositoryRequest: &gitee.CreateRepositoryRequest{
			Name:        gitee.String("repo_name1"), // 仓库名称
			Path:        gitee.String("repo_Path1"), //路径 (请注意：仓库路径即仓库访问 URL 地址，更改仓库路径将导致原克隆地址不可用)
			Description: gitee.String("repo_Description"),
		},
		Public: gitee.Int(1), //仓库开源类型。0(私有), 1(外部开源), 2(内部开源)，注：与private互斥，以public为主。
	}
	repository, response, err := client.Repositories.CreateOrgRepository(ctx, "mamh-mixed", opts)
	fmt.Println("repository", repository)
	fmt.Println("response", response)
	fmt.Println("err", err)
}

func TestCreateEntRepository(t *testing.T) {
	opts := &gitee.CreateEntRepositoryRequest{
		CreateRepositoryRequest: &gitee.CreateRepositoryRequest{
			Name:        gitee.String("magesfc仓库名称"), // 仓库名称
			Path:        gitee.String("repo_Path"),   //路径 (请注意：仓库路径即仓库访问 URL 地址，更改仓库路径将导致原克隆地址不可用)
			Description: gitee.String("repo_Description"),
		},
		Outsourced:     gitee.Bool(true),
		ProjectCreator: gitee.String("mamh"),
		Members:        gitee.String("mamh"),
	}
	// 匿名字段 或者 这样来赋值
	// opts.Name = gitee.String("magesfc仓库名称")
	//opts.Path =  gitee.String("repo_Path")
	//opts.Description = gitee.String("repo_Description")

	repository, response, err := client.Repositories.CreateEntRepository(ctx, "magesfc", opts)
	fmt.Println("repository", repository)

	fmt.Println("response", response)
	fmt.Println("err", err)
}

func TestUpdateBranchProtection(t *testing.T) {
	owner := "mamh-mixed"
	repo := "go-gitee"
	branch := "main"

	protection, response, err := client.Repositories.UpdateBranchProtection(ctx, owner, repo, branch)
	//先执行 UpdateBranchProtection 创建一个名称是 main 的规则。
	//然后再执行 CreateBranchWildcardProtection 会报错，重名了.反过来的话是可以的。
	// 并且还会覆盖掉。先前创建 的报 “标准模式，作用于 0 个分支 规则没有生效？” -> “标准模式，作用于 1 个分支”
	fmt.Println(protection)

	fmt.Println(response)
	fmt.Println(err)

}

func TestRemoveBranchProtection(t *testing.T) {
	owner := "mamh-mixed"
	repo := "go-gitee"
	branch := "main"

	response, err := client.Repositories.RemoveBranchProtection(ctx, owner, repo, branch)
	// RemoveBranchWildcardProtection 用这个反而可以删掉。这接口弄的真乱！

	fmt.Println(response)
	fmt.Println(err) //"message": "Operation is not allowed"
	// "message": "404 Not Found"

}

func TestUpdateBranchWildcardProtection(t *testing.T) {
	owner := "mamh-mixed"
	repo := "go-gitee"
	wildcard := "main_wildcard"
	pr := &gitee.ProtectionRequest{
		NewWildcard: gitee.String("main_new_wildcard1"), // wildcard -> new_wildcard 重命名
		Pusher:      gitee.String("admin"),
		Merger:      gitee.String("admin"),
	}
	prot, response, err := client.Repositories.UpdateBranchWildcardProtection(ctx, owner, repo, wildcard, pr)

	// 400 Bad Request, map[], 分支/通配符已经被使用
	fmt.Println(prot) // gitee.ProtectionSetting{ID:1707231, ProjectID:25870304, Wildcard:"main_new_wildcard1"}

	fmt.Println(response)
	fmt.Println(err)

}

func TestCreateBranchWildcardProtection(t *testing.T) {
	owner := "mamh-mixed"
	repo := "go-gitee"

	pr := &gitee.ProtectionRequest{
		Wildcard: gitee.String("main"), // 注意这里 是 Wildcard 字段, 新建的时候 用的这个字段名，更新的时候用的NewWildcard
		Pusher:   gitee.String("admin"),
		Merger:   gitee.String("admin"),
	}
	prot, response, err := client.Repositories.CreateBranchWildcardProtection(ctx, owner, repo, pr)

	// 400 Bad Request, map[], 分支/通配符已经被使用
	fmt.Println(prot) // gitee.ProtectionSetting{ID:1707235, ProjectID:25870304, Wildcard:"main_wildcard"}

	fmt.Println(response)
	fmt.Println(err)

}

func TestRemoveBranchWildcardProtection(t *testing.T) {
	owner := "mamh-mixed"
	repo := "go-gitee"
	branch := "main"

	response, err := client.Repositories.RemoveBranchWildcardProtection(ctx, owner, repo, branch)
	// UpdateBranchProtection 创建的 用 RemoveBranchProtection 不能删掉
	// 用这个 RemoveBranchWildcardProtection 反而可以删掉。这接口弄的真乱！

	fmt.Println(response)
	fmt.Println(err) // "message": "Wildcard Not Found"

}

func TestGetPagesInfo(t *testing.T) {
	owner := "oschina"
	repo := "git-osc"

	info, response, err := client.Repositories.GetPagesInfo(ctx, owner, repo)

	fmt.Println(*info.URL)
	fmt.Println(*info.Status)

	fmt.Println(response)
	fmt.Println(err)
}

func TestGetRepository(t *testing.T) {
	owner := "oschina"
	repo := "git-osc"

	repository, response, err := client.Repositories.Get(ctx, owner, repo)

	fmt.Println(repository)

	fmt.Println(response)
	fmt.Println(err)
}

func TestEditRepository(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc-test"

	erq := &gitee.EditRepositoryRequest{
		Name:        "new-repo_name",
		Description: "new description",
	}
	repository, response, err := client.Repositories.Edit(ctx, owner, repo, erq)

	fmt.Println(repository)

	fmt.Println(response)
	fmt.Println(err)
}

func TestDeletetRepository(t *testing.T) {
	owner := "magesfc"
	repo := "sample_repository"

	response, err := client.Repositories.Delete(ctx, owner, repo)
	fmt.Println(response)
	fmt.Println(err) // Not Found Project
}

func TestGetPushConfig(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc"
	pushConfig, response, err := client.Repositories.GetPushConfig(ctx, owner, repo)

	fmt.Println(pushConfig)
	fmt.Println(response)
	fmt.Println(err)

}

func TestUpdatePushConfig(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc"
	pushConfig, response, err := client.Repositories.GetPushConfig(ctx, owner, repo)

	fmt.Println(pushConfig)
	fmt.Println(response)
	fmt.Println(err)

	pushConfig.ExceptManager = gitee.Bool(false)

	newConfig, response, err := client.Repositories.UpdatePushConfig(ctx, owner, repo, pushConfig)
	fmt.Println(newConfig)
	fmt.Println(response)
	fmt.Println(err)
}

func TestListContributors(t *testing.T) {
	owner := "log4j"
	repo := "pig"
	opts := &gitee.ListContributorsOptions{}
	contributors, response, err := client.Repositories.ListContributors(ctx, owner, repo, opts)

	fmt.Println(contributors)
	fmt.Println(len(contributors))
	fmt.Println(response)
	fmt.Println(err)
}

func TestListTags(t *testing.T) {
	owner := "log4j"
	repo := "pig"
	opts := &gitee.ListOptions{}

	tags, response, err := client.Repositories.ListTags(ctx, owner, repo, opts)
	fmt.Println(tags)
	fmt.Println(len(tags))
	fmt.Println(response)
	fmt.Println(err)
}

func TestCreateTag(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc"
	ctq := &gitee.CreateTagRequest{
		Refs:       "master",
		TagName:    "v0.0.1",
		TagMessage: "test create tags",
	}
	tag, response, err := client.Repositories.CreateTag(ctx, owner, repo, ctq)

	fmt.Println(tag)
	fmt.Println(response)
	fmt.Println(err)

}

func TestClear(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc"
	response, err := client.Repositories.Clear(ctx, owner, repo)

	fmt.Println(response)
	fmt.Println(err)
}

func TestListCollaborators(t *testing.T) {
	owner := "log4j"
	repo := "pig"
	opts := &gitee.ListOptions{}
	collaborators, response, err := client.Repositories.ListCollaborators(ctx, owner, repo, opts)

	fmt.Println(collaborators)
	fmt.Println(len(collaborators))
	fmt.Println(response)
	fmt.Println(err)
}

func TestIsCollaborator(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc"
	me := "mamh"
	isCollaborator, response, err := client.Repositories.IsCollaborator(ctx, owner, repo, me)

	fmt.Println(isCollaborator)
	fmt.Println(response)
	fmt.Println(err)
}

func TestAddCollaborator(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc"
	me := "mamh"
	acreq := &gitee.AddCollaboratorRequest{
		Permission: "admin",
	}
	ci, response, err := client.Repositories.AddCollaborator(ctx, owner, repo, me, acreq)

	fmt.Println(ci)
	fmt.Println(response)
	fmt.Println(err) //404 Not Found, map[], User Not Found
}

func TestRemoveCollaborator(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc"
	me := "mamh"

	response, err := client.Repositories.RemoveCollaborator(ctx, owner, repo, me)
	fmt.Println(response)
	fmt.Println(err) //400 Bad Request, map[], Owner cannot withdraw from a project
	//404 Not Found, map[], User
}

func TestGetPermissionLevel(t *testing.T) {
	owner := "magesfc"
	repo := "magesfc"
	me := "mamh"

	level, response, err := client.Repositories.GetPermissionLevel(ctx, owner, repo, me)
	fmt.Println(level)
	fmt.Println(response)
	fmt.Println(err)
}

func TestListForks(t *testing.T) {
	owner := "y_project"
	repo := "RuoYi"
	opts := &gitee.RepositoryListForksOptions{}
	f, response, err := client.Repositories.ListForks(ctx, owner, repo, opts)
	fmt.Println(f)
	fmt.Println(response)
	fmt.Println(err)
}

func TestCreateFork(t *testing.T) {
	owner := "y_project"
	repo := "RuoYi"
	opts := &gitee.RepositoryCreateForkOptions{
		Organization: "magesfc",
		Name:         "ruoyi",
		Path:         "ruoyi_git",
	}
	f, response, err := client.Repositories.CreateFork(ctx, owner, repo, opts)
	fmt.Println(f)
	fmt.Println(response)
	fmt.Println(err) // 403 Forbidden, map[], 已经存在同名的仓库（忽略大小写），Fork 失败
}

func TestListTraffic(t *testing.T) {
	owner := "mamh-mixed"
	repo := "go-gitee"
	opts := &gitee.TrafficDataRequest{
		StartDay: "2022-11-01",
		EndDay:   "2022-11-11",
	}
	tr, response, err := client.Repositories.ListTraffic(ctx, owner, repo, opts)
	fmt.Println(tr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestListReleases(t *testing.T) {
	owner := "magesfc"
	repo := "ruoyi_git"
	opts := &gitee.RepositoryReleaseListOptions{}
	rr, response, err := client.Repositories.ListReleases(ctx, owner, repo, opts)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestCreateRelease(t *testing.T) {
	owner := "magesfc"
	repo := "ruoyi_git"
	releaseReq := &gitee.CreateRepositoryReleaseRequest{
		TagName:         "v2.3.4",
		TargetCommitish: "master",
		Name:            "Release 名称",
		Body:            "Release 描述",
		Prerelease:      true,
	}
	rr, response, err := client.Repositories.CreateRelease(ctx, owner, repo, releaseReq)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetRelease(t *testing.T) {
	owner := "magesfc"
	repo := "ruoyi_git"
	id := int64(264806)
	rr, response, err := client.Repositories.GetRelease(ctx, owner, repo, id)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestEditRelease(t *testing.T) {
	owner := "magesfc"
	repo := "ruoyi_git"
	id := int64(264863)
	releaseReq := &gitee.EditReleaseRequest{
		RepositoryReleaseRequest: &gitee.RepositoryReleaseRequest{
			TagName:    gitee.String("v3.3.4"),
			Name:       gitee.String("Release new 名称"),
			Body:       gitee.String("Release new 描述"),
			Prerelease: gitee.Bool(true),
		},
	}
	rr, response, err := client.Repositories.EditRelease(ctx, owner, repo, id, releaseReq)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestDeleteRelease(t *testing.T) {
	owner := "magesfc"
	repo := "ruoyi_git"
	id := int64(264806)

	response, err := client.Repositories.DeleteRelease(ctx, owner, repo, id)
	fmt.Println(response)
	fmt.Println(err) //404 Not Found, map[], 404 Not Found
}

func TestGetLatestRelease(t *testing.T) {
	owner := "magesfc"
	repo := "ruoyi_git"
	rr, response, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetReleaseByTag(t *testing.T) {
	owner := "magesfc"
	repo := "ruoyi_git"
	tag := "v10.10.10"
	rr, response, err := client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestCreateOpenGo(t *testing.T) {
	owner := "magesfc"
	repo := "ruoyi_git"

	response, err := client.Repositories.CreateOpenGo(ctx, owner, repo)
	fmt.Println(response)
	fmt.Println(err)
}
