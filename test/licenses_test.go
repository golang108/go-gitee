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
