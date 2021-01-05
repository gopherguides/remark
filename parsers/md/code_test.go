package md

import (
	"html/template"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_Code_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "code")
}

func Test_Code_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	c := &remark.Code{
		Text: "1+1",
	}

	exp := template.HTML(`<pre><code>1+1</code></pre>`)
	r.Equal(exp, c.HTML())
}
