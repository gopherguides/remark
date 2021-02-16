package htm

import (
	"testing"

	"github.com/gopherguides/remark"

	"github.com/stretchr/testify/require"
)

func Test_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const in = `
	<div class="text-center">
			<img class="img img-fluid mb-3 mx-auto" src="legacy-site.png" title="Legacy Site"></img>
			<p>
					<a href="https://example.com" target="_blank">hi</a>
			</p>
			<details>
					keyA: valueA
					keyB: valueB
			</details>
			<code src="../internal/testdata/buffalo/src/initial/actions/app.go"></code>
	</div>`

	doc, err := Parse([]byte(in))
	r.NoError(err)
	r.NotNil(doc)
	r.Len(doc.GetChildren(), 1)

	html := doc.GetChildren()[0]
	r.Equal("html", html.TagName())

	r.Len(html.GetChildren(), 2)
	r.Equal("head", html.GetChildren()[0].TagName())

	body := html.GetChildren()[1]
	r.Equal("body", body.TagName())
	r.Len(body.GetChildren(), 1)

	div := body.GetChildren()[0]
	r.Equal("div", div.TagName())
	r.Len(div.GetChildren(), 4)

	md := div.GetChildren()[2]
	r.Equal("metadata", md.TagName())
	r.Len(md, 2)
	data, ok := md.(remark.Metadata)
	r.True(ok)
	r.Equal("valueA", data["keya"])

	ct := div.GetChildren()[3]
	r.Equal("code", ct.TagName())
	code, ok := ct.(*remark.CodeBlock)
	r.True(ok)
	r.NotNil(code)

	it, ok := doc.GetChildren().FindFirst("img")
	r.True(ok)

	img, ok := it.(*remark.Generic)
	r.True(ok)

	r.Len(img.Attributes, 3)

}
