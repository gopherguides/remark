package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewBlockQuote(n *html.Node) (Element, error) {
	blockquote := &remark.BlockQuote{}

	if err := set(blockquote, n); err != nil {
		return nil, err
	}

	return blockquote, nil
}
