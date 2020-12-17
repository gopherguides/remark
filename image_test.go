package remark

import (
	"html/template"
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/require"
)

func Test_Image_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "image")
}

func Test_Image_HTML(t *testing.T) {
	t.Parallel()

	table := []struct {
		Image *ast.Image
		Exp   template.HTML
	}{
		{
			Image: &ast.Image{
				Destination: []byte("./foo.png"),
			},
			Exp: `<img src="./foo.png"></img>`,
		},
		{
			Image: &ast.Image{
				Destination: []byte("./foo.png"),
				Title:       []byte("Kangaroo"),
			},
			Exp: `<img alt="Kangaroo" src="./foo.png"></img><figcaption>Kangaroo</figcaption>`,
		},
	}

	for _, tt := range table {
		tt := tt

		t.Run(string(tt.Image.Destination), func(st *testing.T) {
			r := require.New(st)

			img := Image{
				Image: tt.Image,
			}

			r.Equal(tt.Exp, img.HTML())
		})
	}
}
