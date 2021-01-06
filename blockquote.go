package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type BlockQuote struct {
	Attributes
	Elements Elements
}

func (blockquote BlockQuote) Name() string {
	return "blockquote"
}

func (blockquote *BlockQuote) HTML() template.HTML {
	t := tags.New(blockquote.Name(), blockquote.Options())
	t.Append(blockquote.Elements.HTML())
	return t.HTML()
}
