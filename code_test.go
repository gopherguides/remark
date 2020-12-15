package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Code_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "code")
}

func Test_Code_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	c := Code{
		Text: "1+1",
	}

	exp := template.HTML(`<pre><code>1+1</code></pre>`)
	r.Equal(exp, c.HTML())
}
