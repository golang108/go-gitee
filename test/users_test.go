package test

import (
	"context"
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"golang.org/x/oauth2"
	"testing"
)

var (
	client *gitee.Client
	ctx    context.Context
)

func init() {
	token := "your gitee token"

	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client = gitee.NewClient(tc)
}

func TestGetUser(t *testing.T) {
	user, response, err := client.Users.Get(ctx, "")
	fmt.Println(user)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetUserKeys(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,  // 这个key目前只有2个，这里一页就能获取全部的了
		PerPage: 10, // perPage 表示每页的总数
	}
	for {
		keys, response, err := client.Users.GetUserKeys(ctx, opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, key := range keys {
			fmt.Println(index, len(keys), *key.ID, *key.Title)
		}
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
}

func TestGetUserKey(t *testing.T) {
	keys, response, err := client.Users.GetUserKey(ctx, 3544397)
	fmt.Println(keys)
	fmt.Println(*response)
	fmt.Println(err)
}

func TestGetUserFollowers(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    90,  // page 表示从第几页 开始，一般从第 1 页开始，然后第 2 页，然后第 3 页，到最后一页
		PerPage: 100, // perPage 表示每页的总数
	}
	for { //分页 循环 获取 所有的，这里有 几百个值的，需要循环的
		users, response, err := client.Users.GetUserFollowers(ctx, "y_project", opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, user := range users {
			fmt.Println(index, "len: ", len(users), response.NextPage, *user.ID, *user.Login, *user.Name)
		}

		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}

}

func TestGetUserFollowers1(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,
		PerPage: 10,
	}
	users, response, err := client.Users.GetUserFollowers(ctx, "", opts)
	fmt.Println(users)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetUserFollowings(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,
		PerPage: 10,
	}
	// 这里获取mamh这个账号关注了哪几个人
	users, response, err := client.Users.GetUserFollowings(ctx, "mamh", opts)
	fmt.Println(users)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetUserFollowings1(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,
		PerPage: 10,
	}
	// user参数设置空字符串，就是 当前授权账号 关注了哪几个人
	users, response, err := client.Users.GetUserFollowings(ctx, "", opts)
	fmt.Println(users)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetUserNamespaces(t *testing.T) {
	var opts = &gitee.NamespacesOptions{
		Mode: "project",
	}
	names, response, err := client.Users.GetUserNamespaces(ctx, opts)
	fmt.Println(names) //  project 类型的，只会获取2组数据
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetUserNamespace(t *testing.T) {
	var opts = &gitee.NamespaceOptions{
		Path: "mamh-java",
	}
	names, response, err := client.Users.GetUserNamespace(ctx, opts)
	fmt.Println(names)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetUserNamespace1(t *testing.T) {
	var opts = &gitee.NamespaceOptions{
		Path: "mamh",
	}
	names, response, err := client.Users.GetUserNamespace(ctx, opts)
	fmt.Println(names)
	fmt.Println(response)
	fmt.Println(err)
}
