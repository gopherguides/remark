package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewDocument(n *html.Node) (*remark.Document, error) {
	d := &remark.Document{}

	sec := &remark.Section{}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		el, err := NewElement(c)
		if err != nil {
			return nil, err
		}
		if s, ok := el.(*remark.Section); ok {
			d.Sections = append(d.Sections, sec)
			sec = s
		}
		sec.Elements = append(sec.Elements, el)
	}
	d.Sections = append(d.Sections, sec)

	return d, nil
}
