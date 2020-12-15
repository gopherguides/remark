package remark

import "html/template"

type HTMer interface {
	HTML() template.HTML
}
