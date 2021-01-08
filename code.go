package remark

import (
	"fmt"
	"html"
	"html/template"
	"strings"

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

	var lines []string
	for _, e := range code.Elements() {
		switch t := e.(type) {
		case String:
			lines = append(lines, string(t))
		case fmt.Stringer:
			lines = append(lines, t.String())
		default:
			lines = append(lines, fmt.Sprintf("%s", t))
		}
	}

	h := template.HTML(strings.Join(lines, "\n"))

	if s, ok := code.Attributes.Get("snippet"); ok {
		h = code.Snippets[s].HTML()
	}

	hs := html.EscapeString(string(h))
	t.Append(hs)

	pre := tags.New("pre", nil)
	pre.Append(t)
	return pre.HTML()
}
