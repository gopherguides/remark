package md

import (
	"html/template"
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/require"
)

func Test_HTMLBlock_HTML(t *testing.T) {
	t.Skip()
	t.Parallel()
	r := require.New(t)

	const html = `<div class="main"><p id="ello"><img src="./foo.png" class="big"></img></p></div>`

	ahb := &ast.HTMLBlock{
		Leaf: ast.Leaf{
			Literal: []byte(html),
		},
	}

	hb, err := NewHTMLBlock(ahb)
	r.NoError(err)

	exp := template.HTML(html)
	r.Equal(exp, hb.HTML())
}
