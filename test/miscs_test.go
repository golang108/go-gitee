package test

import (
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"testing"
)

func TestListEmojis(t *testing.T) {
	rr, response, err := client.Miscellaneous.ListEmojis(ctx)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}

func TestMarkdown(t *testing.T) {
	opts := &gitee.MarkdownRequest{
		Text: gitee.String("# xxxxxxxxxxxx test"),
	}
	rr, response, err := client.Miscellaneous.Markdown(ctx, opts)
	fmt.Println(rr)
	fmt.Println(response)
	fmt.Println(err)
}
