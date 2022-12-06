package test

import (
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"testing"
)

func TestListOrganizations1(t *testing.T) {
	opts := &gitee.ListOptions{}
	orgs, response, err := client.Organizations.List(ctx, "mamh", opts)

	fmt.Println(orgs)
	fmt.Println(response)
	fmt.Println(err)

}
