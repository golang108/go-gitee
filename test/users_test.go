package test

import (
	"context"
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"golang.org/x/oauth2"
	"testing"
)

var client *gitee.Client

func TestGetUser(t *testing.T) {
	token := "your gitee token"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client = gitee.NewClient(tc)

	user, response, err := client.Users.Get(ctx, "")
	fmt.Println(user)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetUserKeys(t *testing.T) {
	token := "your gitee token"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client = gitee.NewClient(tc)

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
	token := "your gitee token"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client = gitee.NewClient(tc)

	keys, response, err := client.Users.GetUserKey(ctx, 3544397)
	fmt.Println(keys)
	fmt.Println(*response)
	fmt.Println(err)
}
