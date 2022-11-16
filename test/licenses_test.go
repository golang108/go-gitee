package test

import (
	"fmt"
	"testing"
)

func TestListlicenses(t *testing.T) {
	rr, response, err := client.Licenses.List(ctx)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetLicense(t *testing.T) {
	rr, response, err := client.Licenses.Get(ctx, "Apache-2.0")
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}
