package md

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/htm"
	"golang.org/x/net/html"
)

func (p *Parser) include(name string) (string, error) {
	fmt.Println("TODO >> include", name)
	base := filepath.Base(name)
	pn := "_" + base
	dir := filepath.Dir(name)
	fp := filepath.Join(dir, pn)
	fmt.Println("TODO >> include fp", fp)

	b, err := ioutil.ReadFile(fp)
	return string(b), err
}

func (p *Parser) NewInclude(n *html.Node) (*remark.Include, error) {
	body, err := htm.HTMLBody(n)
	if err != nil {
		return nil, err
	}
	inc := body.FirstChild
	if inc == nil {
		return nil, fmt.Errorf("could not find include tag in node %v", n)
	}
	g := htm.NewGeneric(inc)
	return &remark.Include{
		Generic: g,
	}, nil
	// g := remark.NewGeneric(n)

	// i := &remark.Include{
	// 	Generic: g,
	// }

	// src, ok := i.Attributes["src"]
	// fmt.Println("TODO > src", src)
	// if !ok {
	// 	return nil, fmt.Errorf("no src found for include %v", n)
	// }

	// // tags, err := remark.ParseFile(src)
	// // if err != nil {
	// // 	return nil, err
	// // }
	// // i.Children = tags
	// fmt.Println("TODO >> i", i)
	// return i, nil
}
