package md

import (
	"fmt"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/markbates/markdownr/ast"
	"github.com/stretchr/testify/require"
)

func Test_Parser_NewCode(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	p := NewParser(&remark.Module{})

	exp := "var a string"
	node := &ast.Code{
		Leaf: ast.Leaf{
			Literal: []byte(exp),
		},
	}

	code, err := p.NewCode(node)
	r.NoError(err)
	r.Len(code.Children, 1)

	act := code.Children[0]
	r.Equal(exp, act.String())
}

func Test_Parser_NewCodeBlock(t *testing.T) {
	t.Parallel()

	p := NewParser(&remark.Module{})

	table := []struct {
		act string
		exp string
	}{
		{act: "", exp: "plain"},
		{act: "go", exp: "go"},
		{act: "plain", exp: "plain"},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%q/%q", tt.act, tt.exp), func(t *testing.T) {
			r := require.New(t)
			exp := "var a string"
			node := &ast.CodeBlock{
				Info: []byte(tt.act),
				Leaf: ast.Leaf{
					Literal: []byte(exp),
				},
			}

			code, err := p.NewCodeBlock(node)
			r.NoError(err)
			r.Len(code.Children, 1)

			act := code.Children[0]
			r.Equal(exp, act.String())
			r.Equal(tt.exp, code.Language)
		})
	}

}
