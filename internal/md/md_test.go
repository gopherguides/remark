package md

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/stretchr/testify/require"
)

func Test_MD(t *testing.T) {
	t.Skip()
	t.Parallel()
	r := require.New(t)

	b, err := ioutil.ReadFile("./example.md")
	r.NoError(err)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	node := markdown.Parse(b, p)

	cont := node.AsContainer()
	r.Nil(cont.Parent)
	r.Nil(cont.Attribute)
	r.Len(cont.Content, 0)
	r.Len(cont.Literal, 0)
	r.Len(cont.Children, 244)
	var fn ast.NodeVisitorFunc = func(node ast.Node, entering bool) ast.WalkStatus {
		switch t := node.(type) {
		case *ast.Text:
			fmt.Println(string(t.Literal))
		}

		return ast.GoToNext
	}
	ast.Walk(node, fn)

	r.True(false)
}
