package md

import (
	"io/ioutil"
	"path/filepath"
	"testing"

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

	sec := tags[0]
	r.Equal("section", sec.TagName())
	r.Len(sec.GetChildren(), 2)

	h1 := sec.GetChildren()[0]
	r.Equal("h1", h1.TagName())
	r.Len(h1.GetChildren(), 1)
	r.Equal("<h1>Converting a Static Site to Buffalo</h1>", h1.String())

	md, ok := sec.(*Section)
	r.True(ok)
	data := md.Metadata()
	r.Equal("true", data["overview"])
}
