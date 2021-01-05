package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewLink(n *html.Node) (Element, error) {
	link := &remark.Link{
		Attributes: remark.Attributes{},
	}

	if err := set(link, n); err != nil {
		return nil, err
	}

	return link, nil
}
