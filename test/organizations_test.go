package test

import (
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"testing"
)

func TestListOrganizations1(t *testing.T) {
	opts := &gitee.OrganizationListOptions{
		Admin: gitee.Bool(false),
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

func TestGetOrgMembership(t *testing.T) {
	user := ""
	org := "mamh-mixed"
	member, response, err := client.Organizations.GetOrgMembership(ctx, user, org)

	fmt.Println(member)
	fmt.Println(response)
	fmt.Println(err)

	user2 := "mamh"
	org2 := "mamh-mixed"
	member2, response2, err2 := client.Organizations.GetOrgMembership(ctx, user2, org2)

	fmt.Println(member2)
	fmt.Println(response2)
	fmt.Println(err2)
}
