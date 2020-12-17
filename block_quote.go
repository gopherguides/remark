package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type BlockQuote struct {
	Elements Elements
}

func (bq BlockQuote) HTML() template.HTML {
	t := tags.New("blockquote", tags.Options{})
	t.Append(bq.Elements.HTML())
	return t.HTML()
}

func (bq BlockQuote) String() string {
	return fmt.Sprintf("> %s", bq.Elements)
}
