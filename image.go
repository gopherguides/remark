package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

// Image is strikethrough
type Image struct {
	*ast.Image
	Elements Elements
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
