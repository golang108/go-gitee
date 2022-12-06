package test

import (
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"testing"
)

func TestListOrganizations1(t *testing.T) {
	opts := &gitee.OrganizationListOptions{
		Admin: true,
	}
	orgs, response, err := client.Organizations.List(ctx, "mamh", opts)

	fmt.Println(orgs)
	fmt.Println(response)
	fmt.Println(err)

}

func TestListOrgMemberships(t *testing.T) {
	opts := &gitee.MembershipListOptions{
		Active: gitee.Bool(false),
		ListOptions: gitee.ListOptions{
			Page:    1,
			PerPage: 20,
		},
	}
	member, response, err := client.Organizations.ListOrgMemberships(ctx, opts)

	fmt.Println(member)
	fmt.Println(response)
	fmt.Println(err)
}
