package htm

import (
	"fmt"
	"path/filepath"
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

	fsrc := filepath.Join(p.Root, src)
	img := &remark.Image{
		Generic: g,
	}
	defer img.Set("src", fsrc)

	if strings.HasPrefix(src, "http") {
		return img, nil
	}

	if _, err := p.Stat(src); err != nil {
		return nil, err
	}

	return img, nil

}
