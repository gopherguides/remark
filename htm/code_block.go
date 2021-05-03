package htm

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gopherguides/remark"

	"golang.org/x/net/html"
)

func (p *Parser) NewCodeBlock(n *html.Node) (*remark.CodeBlock, error) {
	g, err := p.NewGeneric(n)
	if err != nil {
		return nil, err
	}

	code := &remark.CodeBlock{
		Generic:  g,
		Snippets: remark.Snippets{},
	}

	if s, ok := code.Get("section"); ok {
		// fmt.Printf("[<code section>] is deprecated. Use [<code snippet>] instead. %s\n", s)
		code.Set("snippet", s)
	}

	src, ok := code.Get("src")
	if !ok {
		return code, nil
	}

	code.Language = strings.TrimPrefix(filepath.Ext(src), ".")

	b, err := ioutil.ReadFile(filepath.Join(p.Root, src))
	if err != nil {
		return nil, err
	}

	var lines []string

	open := map[string]remark.Snippet{}

	for i, line := range strings.Split(string(b), "\n") {
		sl := strings.TrimSpace(line)
		if strings.HasPrefix(sl, "// section:") {
			sl = strings.ReplaceAll(sl, "section:", "snippet:")
		}

		if strings.HasPrefix(sl, "// snippet:") {
			name := strings.TrimPrefix(sl, "// snippet:")
			name = strings.TrimSpace(name)
			snip, ok := open[name]
			if ok {
				snip.End = i + 1
				code.Snippets[name] = snip
				delete(open, name)
			}
			if !ok {
				snip.File = src
				snip.Language = code.Language
				snip.Name = name
				snip.Start = i + 1
				open[name] = snip
			}

			code.Snippets[name] = snip
			continue
		}

		for k, snip := range open {
			snip.Content = strings.Join([]string{snip.Content, line}, "\n")
			open[k] = snip
		}

		lines = append(lines, line)
	}

	code.Children = append(code.Children, remark.String(strings.Join(lines, "\n")))

	return code, nil
}
