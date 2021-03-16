package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CodeBlock_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var code CodeBlock
	r.Equal("code", code.Name())
}

func Test_CodeBlock_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	code := CodeBlock{
		Generic: &Generic{
			Attributes: Attributes{
				"src":     "./src/foo.go",
				"snippet": "snip",
			},
			Children: Tags{
				String("hello"),
			},
		},
	}
	code.Set("k1", "one")
	code.Set("k2", "two")
	code.Set("k3", "three")

	exp := template.HTML(`<pre><code id="src-foo.go#snip" k1="one" k2="two" k3="three" language="plain" snippet="snip" src="./src/foo.go"></code></pre>`)
	act := code.HTML()

	r.Equal(exp, act)
}
