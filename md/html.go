package md

import (
	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/htm"
)

func (p *Parser) parseHTML(src []byte) (remark.Tag, error) {

	hp := htm.NewParser(p.Root)
	tag, err := hp.Parse(src)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
