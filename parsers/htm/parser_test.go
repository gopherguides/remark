package htm

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const html = `
<div class="text-center">
		<img src="legacy-site.png" title="Legacy Site" class="img img-fluid mb-3 mx-auto" />
		<p>
				<a href="https://example.com" target="_blank">Hi</a>
		</p>
</div>
`

	doc, err := Parse([]byte(html))
	r.NoError(err)
	r.NotNil(doc)

	act := doc.HTML()
	exp := template.HTML(html)
	r.Equal(exp, act)
}
