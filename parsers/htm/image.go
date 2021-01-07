package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) NewImage(n *html.Node) (Element, error) {
	img := &remark.Image{}

	if err := p.set(img, n); err != nil {
		return nil, err
	}

	return img, nil
}
