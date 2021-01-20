package htm

import (
	"bytes"

	remark "github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func Parse(src []byte) (remark.Tag, error) {
	p := Parser{}
	return p.Parse(src)
}

type Parser struct {
	Root string
}

func (p *Parser) Parse(src []byte) (remark.Tag, error) {
	f := bytes.NewReader(src)

	node, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	return p.NewTag(node)
}
