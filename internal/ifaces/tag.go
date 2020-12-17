package ifaces

import "html/template"

type Tag interface {
	Attrs() Attributes
	Body() Elements
	HTML() template.HTML
}
