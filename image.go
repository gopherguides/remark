package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type Image struct {
	*ast.Image
	Elements Elements
}

func (image Image) HTML() template.HTML {
	opts := tags.Options{
		"src": image.URL(),
	}

	t := tags.New("img", opts)
	capt := image.Caption()
	if len(capt) > 0 {
		opts["alt"] = capt
		fig := tags.New("figcaption", tags.Options{
			"body": capt,
		})
		t.AfterTag = append(t.AfterTag, fig)
	}
	t.Append(image.Elements.HTML())

	return t.HTML()
}

func (image Image) URL() string {
	return string(image.Destination)
}

func (image Image) Caption() string {
	return string(image.Title)
}

func (image Image) String() string {
	if len(image.Title) == 0 {
		return fmt.Sprintf("![%s](%s)", image.Elements, image.URL())
	}
	return fmt.Sprintf("![%s](%s %q)", image.Elements, image.URL(), image.Caption())
}

func NewImage(node *ast.Image) (*Image, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	d := &Image{
		Image:    node,
		Elements: elems,
	}

	return d, nil
}
