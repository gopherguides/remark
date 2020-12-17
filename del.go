package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

// Del is strikethrough
type Del struct {
	Elements Elements
}

func (del Del) String() string {
	return fmt.Sprintf("~~%s~~", del.Elements)
}

func (del Del) HTML() template.HTML {
	strike := tags.New("strike", nil)
	strike.Append(del.Elements)
	return strike.HTML()
}
