package htm

import (
	"fmt"
	"strings"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func (p *Parser) NewMetadata(n *html.Node) (remark.Metadata, error) {
	m := remark.Metadata{}
	c := n.FirstChild
	if c == nil {
		return m, nil
	}

	lines := strings.TrimSpace(c.Data)
	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		secs := strings.Split(line, ":")
		if len(secs) < 2 {
			return nil, fmt.Errorf("invalid metadata: %q", line)
		}
		k := strings.TrimSpace(secs[0])
		v := strings.TrimSpace(secs[1])
		m[k] = v
	}

	return m, nil
}
