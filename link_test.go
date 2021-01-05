package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Link_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	link := &Link{
		Attributes: Attributes{
			"class":  "my-class",
			"href":   "http://example.com",
			"id":     "my-link",
			"target": "_blank",
			"title":  "Example",
		},
		Elements: Elements{
			String("hi"),
		},
	}

	exp := template.HTML(`<a class="my-class" href="http://example.com" id="my-link" target="_blank" title="Example">hi</a>`)
	act := link.HTML()
	r.Equal(exp, act)
}
