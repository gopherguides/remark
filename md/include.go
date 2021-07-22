package md

import (
	"fmt"
	"io/ioutil"

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

func (p *Parser) NewInclude(node *html.Node) (*remark.Include, error) {
	var src string

	g := remark.NewGeneric("include")
	for _, at := range node.Attr {
		if at.Key == "src" {
			src = at.Val
		}
		g.Attributes.Add(at.Key, at.Val)
	}

	if len(src) == 0 {
		return nil, fmt.Errorf("no src found for include %q", htm.NodeString(node))
	}

	np, err := NewParser(p.Info.Join(src), p.Client)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(np.Info.Path)
	if err != nil {
		return nil, err
	}

	tags, err := np.Parse(b)
	if err != nil {
		return nil, err
	}

	g.Children = tags
	inc := &remark.Include{
		Generic: g,
		Info:    info,
	}

	return inc, nil
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
