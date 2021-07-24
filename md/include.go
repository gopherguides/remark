package md

import (
	"fmt"
	"path/filepath"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/px"
	"golang.org/x/net/html"
)

// func (p *Parser) include(name string) (string, error) {
// 	fmt.Println("TODO >> include", name)
// 	base := filepath.Base(name)
// 	pn := "_" + base
// 	dir := filepath.Dir(name)
// 	fp := filepath.Join(dir, pn)
// 	fmt.Println("TODO >> include fp", fp)

// 	b, err := ioutil.ReadFile(fp)
// 	return string(b), err
// }

func (p *Parser) NewInclude(inc *html.Node) (*remark.Include, error) {
	g := remark.NewGeneric("include")

	for _, at := range inc.Attr {
		g.Set(at.Key, at.Val)
	}

	src, ok := g.Get("src")
	if !ok {
		return nil, fmt.Errorf("no src defined for include %v", inc)
	}

	cab, err := p.Sub(filepath.Dir(src))
	if err != nil {
		return nil, err
	}

	mod := &px.Module{
		ModName: filepath.Base(src),
		FS:      cab,
	}

	pp := NewParser(mod)

	tags, err := pp.Parse()
	if err != nil {
		return nil, err
	}

	g.Children = tags
	i := &remark.Include{
		Generic: g,
	}
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
