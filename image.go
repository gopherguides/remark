package remark

import (
	"encoding/json"
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

func (image Image) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"caption":  image.Caption(),
		"content":  image.Elements.String(),
		"elements": image.Elements,
		"literal":  image.String(),
		"url":      image.URL(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", image): m,
	}, "", "  ")
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
