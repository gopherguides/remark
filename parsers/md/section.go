package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewSection(head *remark.Heading, children []ast.Node) (*remark.Section, error) {
	s := &remark.Section{
		Elems: remark.Elements{
			head,
		},
	}

	var kids []ast.Node
	// var err error
	for _, n := range children {
		switch t := n.(type) {
		case *ast.Heading:
			h, err := p.NewHeading(t)
			if err != nil {
				return nil, err
			}
			sec, err := p.NewSection(h, kids)
			if err != nil {
				return nil, err
			}
			s.Elems = append(s.Elements(), sec)
			kids = []ast.Node{}

		default:
			kids = append(kids, t)
		}
	}

	for _, k := range kids {
		el, err := p.NewElement(k)
		if err != nil {
			return nil, err
		}
		s.Elems = append(s.Elements(), el)
	}
	return s, nil
}

type Sections []*remark.Section
