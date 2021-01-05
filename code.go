package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Code struct {
	Text string
}

func (c Code) HTML() template.HTML {
	pre := tags.New("pre", nil)
	code := tags.New("code", nil)
	code.Append(c.Text)
	pre.Append(code)

	return pre.HTML()
}
