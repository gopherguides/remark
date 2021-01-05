package remark

import "html/template"

type CodeSpan struct {
}

func (c CodeSpan) HTML() template.HTML {
	return ""
}
