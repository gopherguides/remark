package htm

import (
	"fmt"
	"os"
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

	inc := &remark.Image{
		Generic: g,
	}

	if strings.HasPrefix(src, "http") {
		return inc, nil
	}

	fp := filepath.Join(p.Root, src)
	if _, err := os.Stat(fp); err != nil {
		return nil, err
	}

	return inc, nil
}
