package test

import (
	"context"
	"github.com/mamh-mixed/go-gitee/gitee"
	"golang.org/x/oauth2"
	"os"
)

var (
	client *gitee.Client
	ctx    context.Context
)

func init() {
	token := os.Getenv("GITEE_TOKEN")

	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client = gitee.NewClient(tc)
}
