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
