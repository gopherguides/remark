package md

import (
	"html/template"
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/require"
)

func Test_List_String(t *testing.T) {
	t.Parallel()
	// compareStringer(t, "list")
}

func Test_List_OL_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	l := List{
		List: &ast.List{
			ListFlags: ast.ListTypeOrdered,
		},
		Elements: Elements{
			ListItem{Elements: Elements{String("Ding")}},
			ListItem{Elements: Elements{String("Dong")}},
		},
	}

	exp := template.HTML(`<ol><li>Ding</li><li>Dong</li></ol>`)
	r.Equal(exp, l.HTML())
}
