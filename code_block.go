package remark

import (
	"fmt"
	"html"
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
)

type CodeBlock struct {
	*Generic
	Language string
	Snippets Snippets
}

func (code CodeBlock) Name() string {
	return "code"
}

func (code *CodeBlock) HTML() template.HTML {
	opts := code.Options()

	lang := code.Language
	if len(lang) == 0 {
		lang = "plain"
	}
	opts["language"] = lang

	t := tags.New(code.Name(), opts)

	var lines []string
	for _, e := range code.Children {
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
	t.Append(strings.TrimSpace(hs))

	pre := tags.New("pre", nil)
	pre.Append(t)
	return pre.HTML()
}

func (code CodeBlock) String() string {
	return html.UnescapeString(string(code.HTML()))
}
