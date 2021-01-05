package htm

import (
	"fmt"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func NewLink(n *html.Node) (Element, error) {
	a := &remark.Link{}
	fmt.Printf(">>>TODO parsers/htm/link.go:10: n.Data (%T): %v \n", n.Data, n.Data)
	fmt.Printf(">>>TODO parsers/htm/link.go:10: n.Namespace (%T): %v \n", n.Namespace, n.Namespace)
	fmt.Printf(">>>TODO parsers/htm/link.go:14: n.Attr (%T): %v \n", n.Attr, n.Attr)
	return a, nil
}
