package remark

import (
	"html"
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Code struct {
	*Generic
}

func (code Code) Name() string {
	return "code"
}

func (code *Code) HTML() template.HTML {
	opts := code.Options()

	opts["language"] = "plain"

	t := tags.New(code.Name(), opts)
	t.Append(code.GetChildren().String())

	return t.HTML()
}

func (code Code) String() string {
	return html.UnescapeString(string(code.HTML()))
}
