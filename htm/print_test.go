package htm

import (
	"path"
	"testing"

	"github.com/gobuffalo/tags/v3"
	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_Printer(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	input := `
	<h1>
		<strong>words</strong>
	</h1>
	<p>
		<a href="http://example.com">
			<img src="./assets/vscode-trust.png" />
		</a>
		<code>foo</code>
	</p>`

	parser, err := NewParser("../internal/testdata/week00/module.md", nil)
	r.NoError(err)

	doc, err := parser.Parse([]byte(input))
	r.NoError(err)
	r.NotNil(doc)

	p := &Printer{}

	act, err := p.Print(doc)
	r.NoError(err)

	// fmt.Println(act)

	exp := `<html><head></head> <body><h1><strong>words</strong></h1> <p><a href="http://example.com"><img src="./assets/vscode-trust.png"></img></a> <pre><code language="plain">foo</code></pre></p></body></html>`
	r.Equal(exp, act)

	p.Set("img", func(tag remark.Tag) (string, error) {
		tg := tags.New("div", tags.Options{
			"class": "img-wrapper",
		})
		tg.Append(tags.New(tag.TagName(), tag.Options()))
		if kids := tag.GetChildren(); len(kids) > 0 {
			s, err := p.Print(kids...)
			if err != nil {
				return "", err
			}
			tg.Append(s)
		}
		return tg.String(), nil
	})

	p.Decorate("img", func(tag remark.Tag) (remark.Tag, error) {
		attrs := tag.Attrs()
		src, _ := attrs.Get("src")
		attrs.Set("src", "https://"+path.Join("example.com", src))
		attrs.Set("width", "100px")
		return tag, nil
	})

	p.Decorate("a", func(tag remark.Tag) (remark.Tag, error) {
		attrs := tag.Attrs()
		attrs.Set("target", "_blank")
		return tag, nil
	})

	p.Set("code", func(tag remark.Tag) (string, error) {
		return "BAR", nil
	})

	act, err = p.Print(doc)
	r.NoError(err)

	exp = `<html><head></head> <body><h1><strong>words</strong></h1> <p><a href="http://example.com" target="_blank"><div class="img-wrapper"><img src="https://example.com/assets/vscode-trust.png" width="100px" /></div></a> BAR</p></body></html>`
	r.Equal(exp, act)

}
