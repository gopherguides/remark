package remark

import (
	"html"
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Code struct {
	Attributes
	Language string
	Elems    Elements
	Snippets Snippets
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

	h := code.Elements().HTML()
	if s, ok := code.Attributes.Get("snippet"); ok {
		h = code.Snippets[s].HTML()
	}

	hs := html.EscapeString(string(h))
	t.Append(hs)

	pre := tags.New("pre", nil)
	pre.Append(t)
	return pre.HTML()
}
