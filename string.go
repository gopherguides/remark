package remark

import "html/template"

type String string

func (s String) String() string {
	return string(s)
}

func (s String) HTML() template.HTML {
	return template.HTML(string(s))
}
