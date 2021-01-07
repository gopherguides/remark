package htm

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const html = `<html><head></head><body><div class="text-center"><img class="img img-fluid mb-3 mx-auto" src="legacy-site.png" title="Legacy Site"></img><p><a href="https://example.com" target="_blank">hi</a></p></div></body></html>`

	doc, err := Parse([]byte(html))
	r.NoError(err)
	r.NotNil(doc)

	act := doc.HTML()
	exp := template.HTML(html)
	r.Equal(exp, act)
}
