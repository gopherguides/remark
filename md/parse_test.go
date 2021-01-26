package md

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	fp := "../internal/testdata/buffalo/module.md"
	b, err := ioutil.ReadFile(fp)
	r.NoError(err)

	tags, err := Parse(filepath.Dir(fp), b)
	r.NoError(err)
	r.NotNil(tags)
	for _, t := range tags {
		r.Equal("section", t.TagName())
	}

	r.Len(tags, 20)
	r.Contains(tags[19].String(), `https://twitter.com/corylanou`)

	sec, ok := tags[0].(*Section)
	r.True(ok)
	r.Equal("section", sec.TagName())
	r.Len(sec.GetChildren(), 2)
	r.Equal("Converting a Static Site to Buffalo", sec.Title)

	h1 := sec.GetChildren()[0]
	r.Equal("h1", h1.TagName())
	r.Len(h1.GetChildren(), 1)
	r.Equal("<h1>Converting a Static Site to Buffalo</h1>", h1.String())

	p, ok := sec.GetChildren()[1].(*remark.Generic)
	r.True(ok)
	r.Equal("p", p.TagName())
	r.Len(p.GetChildren(), 3)

	code, ok := p.GetChildren()[1].(*remark.Code)
	r.True(ok)
	r.Equal("code", code.TagName())

	exp := template.HTML(`<code language="plain">code</code>`)
	r.Equal(exp, code.HTML())
}
