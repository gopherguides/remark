package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewDocument(node *ast.Document) (*remark.Document, error) {
	m := &remark.Document{}

	var kids []ast.Node
	var head *remark.Heading
	var err error
	for _, n := range node.Children {
		switch t := n.(type) {
		case *ast.Heading:
			if t.Level > 1 {
				kids = append(kids, t)
				continue
			}

			if len(kids) > 0 {
				sec, err := p.NewSection(head, kids)
				if err != nil {
					return nil, err
				}
				m.Elements = append(m.Elements, sec)
				kids = []ast.Node{}
			}

			head, err = p.NewHeading(t)
			if err != nil {
				return nil, err
			}

		default:
			kids = append(kids, t)
		}
	}

	if len(kids) > 0 {
		sec, err := p.NewSection(head, kids)
		if err != nil {
			return nil, err
		}
		m.Elements = append(m.Elements, sec)
	}

	return m, nil
}
