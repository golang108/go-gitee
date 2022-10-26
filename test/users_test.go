package test

import (
	"context"
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"golang.org/x/oauth2"
	"testing"
)

const (
	testOrg  = "cve-manage-test"
	testRepo = "config"
)

var client *gitee.Client

func init() {
	token := "your gitee token"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client = gitee.NewClient(tc)
}

func TestInit(t *testing.T) {
	fmt.Println(client)
}
