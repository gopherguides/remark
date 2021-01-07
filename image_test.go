package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Image_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var img Image
	r.Equal("img", img.Name())
}

func Test_Image_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	img := Image{
		Attributes: Attributes{},
		Elems: Elements{
			String("hello"),
		},
	}
	img.Set("k1", "one")
	img.Set("k2", "two")
	img.Set("k3", "three")

	exp := template.HTML(`<img k1="one" k2="two" k3="three">hello</img>`)
	act := img.HTML()
	r.Equal(exp, act)
}
