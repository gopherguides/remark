package htm

import (
	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)


func NewImage(n *html.Node) (Element, error) {
	img := &remark.Image{}

	if err := set(img, n); err != nil {
		return nil, err
	}

	return img, nil
}
