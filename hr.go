package remark

import (
	"html/template"
)

type HR struct{}

func (hr HR) HTML() template.HTML {
	return template.HTML(`<hr>`)
}
