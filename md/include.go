package md

import (
	"fmt"
	"path/filepath"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/htm"
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
	g, err := htm.NewGeneric(inc)
	if err != nil {
		return nil, err
	}

	i := &remark.Include{
		Generic: g,
	}

	src, ok := i.Attributes["src"]
	if !ok {
		return nil, fmt.Errorf("no src defined for include %v", inc)
	}

	ext := filepath.Ext(src)
	if ext != ".md" {
		return nil, fmt.Errorf("only .md can be included %v", src)
	}

	root := p.Root
	if len(filepath.Ext(root)) != 0 {
		root = filepath.Dir(root)
	}
	fp := filepath.Join(root, src)

	tags, err := ParseFile(fp)
	if err != nil {
		return nil, err
	}

	i.Children = tags
	return i, nil

}
