package md

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Paragraph_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	p := Paragraph{
		Elements: Elements{
			String("angels"),
		},
	}

	exp := template.HTML(`<p>angels</p>`)
	r.Equal(exp, p.HTML())
}
