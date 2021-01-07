package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Code struct {
	Attributes
	Language string
	Elems    Elements
}

func (code Code) Elements() Elements {
	return code.Elems
}

func (code Code) Name() string {
	return "code"
}

func (code *Code) HTML() template.HTML {
	opts := code.Options()

	lang := code.Language
	if len(lang) == 0 {
		lang = "plain"
	}
	opts["language"] = lang

	t := tags.New(code.Name(), opts)
	t.Append(code.Elements().HTML())

	pre := tags.New("pre", nil)
	pre.Append(t)
	return pre.HTML()
}
