package htm

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) NewCode(n *html.Node) (*remark.Code, error) {
	code := &remark.Code{
		Snippets: remark.Snippets{},
	}
	if err := p.set(code, n); err != nil {
		return nil, err
	}

	src, ok := code.Get("src")
	if !ok {
		return code, nil
	}

	b, err := ioutil.ReadFile(filepath.Join(p.Root, src))
	if err != nil {
		return nil, err
	}

	var lines []string
	// var snip remark.Snippet
	// var cont []string
	var cur string
	for i, line := range strings.Split(string(b), "\n") {
		sl := strings.TrimSpace(line)
		if strings.HasPrefix(sl, "// section:") {
			fmt.Printf("[// section:] has been deprecated. Use [// snippet:] instead. %s\n", sl)
			sl = strings.ReplaceAll(sl, "section:", "snippet:")
		}

		if strings.HasPrefix(sl, "// snippet:") {
			name := strings.TrimPrefix(sl, "// snippet:")
			name = strings.TrimSpace(name)
			snip, ok := code.Snippets[name]
			if ok {
				snip.End = i + 1
				cur = ""
			}
			if !ok {
				snip.File = src
				snip.Language = strings.TrimPrefix(filepath.Ext(src), ".")
				snip.Name = name
				snip.Start = i + 1
				cur = name
			}

			code.Snippets[name] = snip
			continue
		}

		if len(cur) > 0 {
			snip := code.Snippets[cur]
			snip.Content = strings.Join([]string{snip.Content, line}, "\n")
			code.Snippets[cur] = snip
			continue
		}
		lines = append(lines, line)
	}

	code.Elems = append(code.Elems, String(strings.Join(lines, "\n")))

	return code, nil
}
