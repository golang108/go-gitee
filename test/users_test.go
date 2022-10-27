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
		Page:    1,
		PerPage: 10,
	}
	keys, response, err := client.Users.GetUserKeys(ctx, opts)
	fmt.Println(keys[0].String())
	fmt.Println(*response)
	fmt.Println(err)
}

func TestGetUserKey(t *testing.T) {
	keys, response, err := client.Users.GetUserKey(ctx, 3544397)
	fmt.Println(keys)
	fmt.Println(*response)
	fmt.Println(err)
}

func TestGetUserFollowers(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,
		PerPage: 10,
	}
	users, response, err := client.Users.GetUserFollowers(ctx, "y_project", opts)
	fmt.Println(users)
	fmt.Println(response)
	fmt.Println(err)
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
