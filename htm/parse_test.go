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

	tag, err := Parse([]byte(in))
	r.NoError(err)
	r.NotNil(tag)
	r.Len(tag.GetChildren(), 1)

	tag = tag.GetChildren()[0]
	r.Equal("html", tag.TagName())

	r.Len(tag.GetChildren(), 2)
	r.Equal("head", tag.GetChildren()[0].TagName())

	tag = tag.GetChildren()[1]
	r.Equal("body", tag.TagName())
	r.Len(tag.GetChildren(), 1)

	tag = tag.GetChildren()[0]
	r.Equal("div", tag.TagName())
	r.Len(tag.GetChildren(), 4)

	mdtag := tag.GetChildren()[2]
	r.Equal("metadata", mdtag.TagName())

	meta, ok := mdtag.(remark.Metadatable)
	r.True(ok)
	md := meta.Metadata()
	r.Len(md, 2)

	tag = tag.GetChildren()[3]
	r.Equal("code", tag.TagName())

	code, ok := tag.(*remark.CodeBlock)
	r.True(ok)
	r.NotNil(code)
}
