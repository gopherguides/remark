package md

import (
	"html/template"
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/require"
)

func Test_Link_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "link")
}

func Test_Link_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	l := Link{
		Link: &ast.Link{
			Destination: []byte("http://example.com"),
			Title:       []byte("Devil Man"),
		},
		Elements: Elements{String("The Dance")},
	}

	exp := template.HTML(`<a alt="Devil Man" href="http://example.com">The Dance</a>`)
	r.Equal(exp, l.HTML())
}
