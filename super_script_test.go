package remark

import (
	"html/template"
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/require"
)

func Test_SuperScript_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "super_script")
}

func Test_SuperScript_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := SuperScript{
		Superscript: &ast.Superscript{
			Leaf: ast.Leaf{
				Literal: []byte("spice rack"),
			},
		},
	}

	exp := template.HTML(`<sup>spice rack</sup>`)
	r.Equal(exp, s.HTML())
}
