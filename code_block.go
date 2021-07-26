package remark

import (
	"fmt"
	"html"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/flect"
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

func (code CodeBlock) ID() string {
	id, _ := code.Get("src")
	ext := filepath.Ext(id)
	id = strings.TrimSuffix(id, ext)
	id = flect.Dasherize(id)
	id += ext

	if s, ok := code.Attributes.Get("snippet"); ok {
		id += "#" + s
	}

	return id
}

func (code *CodeBlock) HTML() template.HTML {
	opts := code.Options()

	lang := code.Language
	if len(lang) == 0 {
		lang = "plain"
	}
	opts["language"] = lang

	var lines []string
	for _, e := range code.Children {
		switch t := e.(type) {
		case String:
			lines = append(lines, string(t))
		case fmt.Stringer:
			lines = append(lines, t.String())
		default:
			lines = append(lines, t.String())
		}
	}

	h := template.HTML(strings.Join(lines, "\n"))

	if s, ok := code.Attributes.Get("snippet"); ok {
		h = code.Snippets[s].HTML()
	}

	if id := code.ID(); len(id) > 0 {
		opts["id"] = id
	}

	t := tags.New(code.Name(), opts)
	hs := html.EscapeString(string(h))
	t.Append(strings.TrimSpace(hs))

	pre := tags.New("pre", nil)
	pre.Append(t)
	return pre.HTML()
}

func (code CodeBlock) String() string {
	return html.UnescapeString(string(code.HTML()))
}
