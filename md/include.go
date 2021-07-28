package md

import (
	"fmt"
	"path/filepath"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/px"
	"golang.org/x/net/html"
)

func (p *Parser) NewInclude(inc *html.Node) (*remark.Include, error) {
	g := remark.NewGeneric("include")

	for _, at := range inc.Attr {
		g.Set(at.Key, at.Val)
	}

	src, ok := g.Get("src")
	if !ok {
		return nil, fmt.Errorf("no src defined for include %v", inc)
	}

	dir := filepath.Dir(src)
	cab, err := p.Sub(dir)
	if err != nil {
		return nil, err
	}

	mod := &px.Module{
		Root: filepath.Join(p.Root, dir),
		Name: filepath.Base(src),
		FS:   cab,
	}

	pp := NewParser(mod)

	tags, err := pp.Parse()
	if err != nil {
		return nil, fmt.Errorf("include.go:41 [%w] %v", err, cab)
	}

	g.Children = tags
	i := &remark.Include{
		Generic: g,
	}
	// fmt.Println("TODO >> include.go:57 i", i)
	return i, nil

	// g, err := htm.NewGeneric(inc)
	// if err != nil {
	// 	return nil, err
	// }

	// i := &remark.Include{
	// 	Generic: g,
	// }

	// src, ok := i.Attributes["src"]
	// if !ok {
	// 	return nil, fmt.Errorf("no src defined for include %v", inc)
	// }

	// ext := filepath.Ext(src)
	// if ext != ".md" {
	// 	return nil, fmt.Errorf("only .md can be included %v", src)
	// }

	// root := p.Root
	// if len(filepath.Ext(root)) != 0 {
	// 	root = filepath.Dir(root)
	// }
	// fp := filepath.Join(root, src)

	// tags, err := ParseFile(fp)
	// if err != nil {
	// 	return nil, err
	// }

	// i.Children = tags
	// return i, nil

}
