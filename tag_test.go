package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Tags_Body(t *testing.T) {
	r := require.New(t)

	html := NewGeneric("html")
	html.Children = append(html.Children, NewGeneric("head"))
	body := NewGeneric("body")

	inc := NewGeneric("include")
	inc.Attributes["src"] = "foo.go"

	body.Children = append(body.Children, inc, String("hello"))
	html.Children = append(html.Children, body)

	b, ok := html.Children.Body()
	r.True(ok)
	r.NotNil(b)
}
