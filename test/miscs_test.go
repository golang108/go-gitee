package test

import (
	"fmt"
	"testing"
)

func TestListEmojis(t *testing.T) {
	rr, response, err := client.Miscellaneous.ListEmojis(ctx)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}
