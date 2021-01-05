package remark

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Image_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	image := &Image{
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

	exp := template.HTML(`<img class="my-class" href="http://example.com" id="my-link" target="_blank" title="Example">hi</img>`)
	act := image.HTML()
	fmt.Println(act)
	r.Equal(exp, act)
}
