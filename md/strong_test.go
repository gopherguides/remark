package md

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Strong_String(t *testing.T) {
	t.Parallel()
	t.Skip()
	compareStringer(t, "strong")
}

func Test_Strong_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := Strong{
		Elements: Elements{
			String("future"),
		},
	}

	exp := template.HTML(`<strong>future</strong>`)
	r.Equal(exp, s.HTML())
}
