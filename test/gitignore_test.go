package test

import (
	"fmt"
	"testing"
)

func TestListgitignore(t *testing.T) {
	rr, response, err := client.Gitignores.List(ctx)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetGitignore(t *testing.T) {
	rr, response, err := client.Gitignores.Get(ctx, "C")
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}
