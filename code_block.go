package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type CodeBlock struct {
	Language string
	Text     string
}

func (cb CodeBlock) HTML() template.HTML {
	pre := tags.New("pre", nil)
	code := tags.New("code", tags.Options{
		"class": cb.Language,
		"body":  cb.Text,
	})
	pre.Append(code)
	return pre.HTML()
}

type CodeBlocks []CodeBlocks
