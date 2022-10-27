package test

import (
	"fmt"
	"testing"
)

func TestListBranches(t *testing.T) {
	repos, response, err := client.Repositories.ListBranches(ctx, "mamh-java", "jenkins-jenkins")
	fmt.Println(repos)
	fmt.Println(response)
	fmt.Println(err)
}
