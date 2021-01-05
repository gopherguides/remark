package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Image struct {
	Attributes Attributes
	Elements   Elements
}

func (i Image) Attrs() Attributes {
	return i.Attributes
}

func (image Image) HTML() template.HTML {
	opts := tags.Options{
		"src": image.URL(),
	}

	t := tags.New("img", opts)
	capt := image.Title()
	if len(capt) > 0 {
		if _, ok := opts["alt"]; !ok {
			opts["alt"] = capt
		}
		fig := tags.New("figcaption", tags.Options{
			"body": capt,
		})
		t.AfterTag = append(t.AfterTag, fig)
	}
	t.Append(image.Elements.HTML())

	return t.HTML()
}

func (image Image) URL() string {
	return image.Attributes["src"]
}

func (image Image) Title() string {
	return image.Attributes["title"]
}
