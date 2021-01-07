package md

import (
	"fmt"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/parsers/htm"
)

func htmlBody(b []byte) (remark.Element, error) {
	doc, err := htm.Parse(b)
	if err != nil {
		return nil, err
	}
	e, ok := htm.Find("body", doc)
	if !ok {
		return nil, fmt.Errorf("html body not found")
	}

	if el, ok := e.(remark.Elemental); ok {
		return el.Elements(), nil
	}
	return e, nil
}
