package remark

import "html/template"

type Snippet struct {
	Content  string
	File     string
	Language string
	Name     string
	Start    int
	End      int
}

func (snip Snippet) HTML() template.HTML {
	return template.HTML(snip.Content)
}

type Snippets map[string]Snippet
