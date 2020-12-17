package remark

import "html/template"

type CodeSpan struct {
}

func (c CodeSpan) String() string {
	return ""
}

func (c CodeSpan) HTML() template.HTML {
	return ""
}
