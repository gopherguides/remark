package parser

import (
	"fmt"
	"testing"

	"github.com/gopherguides/remark/ast"
	"github.com/stretchr/testify/require"
)

func Test_Parser(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	res, err := ParseFile("./example.md")
	r.NoError(err)
	r.NotNil(res)

	nodes, ok := res.(ast.Nodes)
	r.True(ok)
	r.Len(res, 4)

	h, ok := nodes[0].(*ast.Header)
	r.True(ok)
	r.Equal("Header One\n", h.Text.String())
	r.Equal(1, h.Level)
	r.Len(h.Nodes, 2)

	p, ok := h.Nodes[0].(*ast.Paragraph)
	r.True(ok)
	r.Contains(p.Nodes.String(), "Since its inception")

	p, ok = h.Nodes[1].(*ast.Paragraph)
	r.True(ok)
	r.Contains(p.Nodes.String(), "EDM")

	hr, ok := nodes[1].(*ast.HR)
	r.True(ok)
	r.Equal("---", hr.Text)

	h, ok = nodes[2].(*ast.Header)
	r.True(ok)
	r.Equal("Header Two\n", h.Text.String())
	r.Equal(2, h.Level)
	r.Len(h.Nodes, 1)

	p, ok = h.Nodes[0].(*ast.Paragraph)
	r.True(ok)
	r.Contains(p.Nodes.String(), "concurrent proximate events")

	h, ok = nodes[3].(*ast.Header)
	r.True(ok)
	r.Equal("Header 3hree\n", h.Text.String())
	r.Equal(3, h.Level)
	r.Len(h.Nodes, 1)

	fmt.Println("nodes", nodes)
	r.True(false)
}
