package md

import (
	"html/template"
	"testing"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
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

		t.Run(string(tt.Image.Destination)+string(tt.Image.Title), func(st *testing.T) {
			r := require.New(st)

			img := &remark.Image{
				Attributes: remark.Attributes{
					"src":   string(tt.Image.Destination),
					"title": string(tt.Image.Title),
				},
			}

			r.Equal(tt.Exp, img.HTML())
		})
	}
}
