package htm

import (
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_set(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	link := &remark.Link{}

	text := &html.Node{
		Type: html.TextNode,
		Data: "Hello",
	}

	node := &html.Node{
		Type: html.ElementNode,
		Data: "img",
		Attr: []html.Attribute{
			{Key: "alt", Val: "bar"},
			{Key: "class", Val: "my-image"},
			{Key: "src", Val: "foo.png"},
		},
		FirstChild: text,
	}

	err := set(link, node)
	r.NoError(err)

	r.NotEmpty(link.Attributes)
	r.Equal(link.Attributes, remark.Attributes{
		"alt":   "bar",
		"class": "my-image",
		"src":   "foo.png",
	})

	r.NotEmpty(link.Elements)

	r.Len(link.Elements, 1)
}
