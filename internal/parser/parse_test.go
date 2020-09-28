package parser

import (
	"testing"

	"github.com/gopherguides/remark/ast"
	"github.com/stretchr/testify/require"
)

func Test_Parser(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	res, err := Parse("", []byte(example))
	r.NoError(err)
	r.NotNil(res)

	nodes, ok := res.(ast.Nodes)
	r.True(ok)
	r.Len(res, 4)

	h, ok := nodes[0].(*ast.Header)
	r.True(ok)
	r.Equal("Header One", h.Text)
	r.Equal(1, h.Level)
	r.Equal("# Header One", h.String())

	h, ok = nodes[1].(*ast.Header)
	r.True(ok)
	r.Equal("Header Two", h.Text)
	r.Equal(2, h.Level)
	r.Equal("## Header Two", h.String())

	h, ok = nodes[3].(*ast.Header)
	r.True(ok)
	r.Equal("Header 3hree", h.Text)
	r.Equal(3, h.Level)
	r.Equal("### Header 3hree", h.String())
}

const example = `
# Header One

## Header Two

asdfasdf

### Header 3hree `
