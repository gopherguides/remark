package htm

import (
	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewSection(n *html.Node) (*remark.Section, error) {
	s := &remark.Section{}

	if err := set(s, n); err != nil {
		return nil, err
	}

	return s, nil
}
