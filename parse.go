package remark

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gopherguides/remark/internal/md"
)

type Parser struct {
	*parser.Parser
}

func (p *Parser) Parse(src []byte) (*Module, error) {
	if p.Parser == nil {
		p.Parser = md.NewParser()
	}

	node := markdown.Parse(src, p.Parser)

	doc, ok := node.(*ast.Document)
	if !ok {
		return nil, fmt.Errorf("expected a document, got %T", node)
	}

	m, err := NewModule(doc)
	if err != nil {
		return nil, err
	}

	// var el Element
	// current := &Section{}
	// var sub *Section
	var fn ast.NodeVisitorFunc = func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}

		switch t := node.(type) {
		case *ast.Heading:
			if t.Level <= 1 {
				return ast.SkipChildren
			}
		default:
			fmt.Printf(">>>TODO parse.go:48: t (%T): %v \n", t, t)
		}

		// if !entering {
		// 	return ast.GoToNext
		// }
		//
		// el, err := NewElement(node)
		// if err != nil {
		// 	return ast.Terminate
		// }
		//
		// switch t := el.(type) {
		// case *Heading:
		// 	if entering {
		// 		switch t.Level {
		// 		case 0, 1:
		// 			t.Level = 1
		// 			if current.Heading == nil {
		// 				m.Sections = append(m.Sections, current)
		// 			}
		// 			current, err = NewSection(t)
		// 			if err != nil {
		// 				return ast.Terminate
		// 			}
		// 		}
		// 	}
		// 	if !entering {
		// 		switch t.Level {
		// 		case 0, 1:
		// 			panic(current)
		// 			m.Sections = append(m.Sections, current)
		// 		}
		// 	}
		// }
		//
		// switch t := el.(type) {
		// case *Heading:
		// 	switch t.Level {
		// 	case 0, 1:
		// 		m.Sections = append(m.Sections, current)
		//
		// 		t.Level = 1
		// 		current, err = NewSection(t)
		// 		if err != nil {
		// 			return ast.Terminate
		// 		}
		// 	default:
		// 			if sub != nil {
		// 					current.Sections = append()
		// 			}
		// 		sub, err = NewSection(t)
		// 		if err != nil {
		// 			return ast.Terminate
		// 		}
		// 		current.SubSections = append(current.SubSections, s)
		// 		current = s
		// 	}
		// case *Paragraph:
		// 	fmt.Printf(">>>TODO parse.go:66: t (%T): %v \n", t, t)
		// 	if len(t.String()) != 0 {
		// 		current.Elements = append(current.Elements, t)
		// 	}
		// 	return ast.SkipChildren
		// default:
		// 	current.Elements = append(current.Elements, t)
		// }

		return ast.GoToNext
	}

	ast.Walk(node, fn)

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (p *Parser) ParseFile(filename string) (*Module, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return p.Parse(b)
}

func (p *Parser) ParseReader(r io.Reader) (*Module, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return p.Parse(b)
}

func Parse(src []byte) (*Module, error) {
	p := &Parser{}
	return p.Parse(src)
}

func ParseFile(filename string) (*Module, error) {
	p := &Parser{}
	return p.ParseFile(filename)
}

func ParseReader(r io.Reader) (*Module, error) {
	p := &Parser{}
	return p.ParseReader(r)
}
