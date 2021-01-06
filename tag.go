package remark

import "html/template"

type Tag interface {
	Element
	HTML() template.HTML
}
