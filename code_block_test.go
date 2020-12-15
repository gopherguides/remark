package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CodeBlock_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "code_block")
}

func Test_CodeBlock_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	cb := CodeBlock{
		Language: "go",
		Text:     "print(1)\nprint(2)",
	}

	exp := template.HTML(`<pre><code class="go">print(1)
print(2)</code></pre>`)
	r.Equal(exp, cb.HTML())
}
