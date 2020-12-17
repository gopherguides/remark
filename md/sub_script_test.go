package md

import (
	"html/template"
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/require"
)

func Test_SubScript_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "sub_script")
}

func Test_SubScript_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := SubScript{
		Subscript: &ast.Subscript{
			Leaf: ast.Leaf{
				Literal: []byte("Simon"),
			},
		},
	}

	exp := template.HTML(`<sub>Simon</sub>`)
	r.Equal(exp, s.HTML())
}
