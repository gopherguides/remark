package htm

import (
	"fmt"
	"strings"

	remark "github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) parseMetadata(node *html.Node) (*remark.Generic, error) {
	node.Data = "metadata"
	tag, err := p.NewGeneric(node)
	if err != nil {
		return nil, err
	}

	tag.Name = "metadata"

	var lines []string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		tg, err := p.NewTag(c)
		if err != nil {
			return nil, err
		}
		if tg == nil {
			continue
		}
		txter, ok := tg.(remark.Textable)
		if !ok {
			continue
		}
		for _, line := range strings.Split(txter.Text(), "\n") {
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				lines = append(lines, line)
			}
		}
	}

	md := remark.Metadata{}
	for _, line := range lines {
		ss := strings.Split(line, ":")
		if len(ss) < 2 {
			return nil, fmt.Errorf("invalid metadata %q", line)
		}
		k := strings.TrimSpace(strings.ToLower(ss[0]))
		v := strings.TrimSpace(ss[1])
		md[k] = v
	}
	tag.Data = md
	return tag, nil
}
