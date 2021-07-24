package htm

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func (p *Parser) NewImage(n *html.Node) (*remark.Image, error) {
	g, err := p.NewGeneric(n)
	if err != nil {
		return nil, err
	}

	src, ok := g.Attributes["src"]
	if !ok {
		return nil, fmt.Errorf("no src for img %v", n)
	}

	inc := &remark.Image{
		Generic: g,
	}

	if strings.HasPrefix(src, "http") {
		return inc, nil
	}

	sfs, ok := p.FS.(fs.StatFS)
	if !ok {
		f, err := p.Open(src)
		defer f.Close()

		if err != nil {
			return nil, err
		}
		fmt.Println("TODO >> image.go:35 f", f)
		return inc, nil
	}

	if _, err := sfs.Stat(src); err != nil {
		return nil, err
	}

	return inc, nil
}
