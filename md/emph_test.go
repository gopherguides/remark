package md

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Emph_string(t *testing.T) {
	t.Parallel()
	compareStringer(t, "emph")
}

func Test_Emph_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	e := Emph{
		Elements: Elements{
			String("Steve"),
		},
	}

	exp := template.HTML(`<em>Steve</em>`)
	r.Equal(exp, e.HTML())
}
