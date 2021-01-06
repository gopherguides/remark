package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewHTML(n *html.Node) (Element, error) {
	html := &remark.HTML{}

	if err := set(html, n); err != nil {
		return nil, err
	}

	return html, nil
}
