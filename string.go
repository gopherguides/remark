package remark

import (
	"html/template"
	"strings"
)

type String string

func (s String) Name() string {
	return ""
}

func (s String) HTML() template.HTML {
	return template.HTML(strings.TrimSpace(string(s)))
}
