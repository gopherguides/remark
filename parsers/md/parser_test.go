package md

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/parsers/htm"
	"github.com/stretchr/testify/require"
)

func Test_Parser_Parse(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	fp := "../internal/testdata/buffalo/module.md"
	b, err := ioutil.ReadFile(fp)
	r.NoError(err)

	p := &Parser{
		Root: filepath.Dir(fp),
	}
	m, err := p.Parse(b)
	r.NoError(err)
	r.NotNil(m)

	// fmt.Println(m.HTML())
	// r.False(true)
}

func Test_Parser_Parse_Metadata(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	in := `
# Foo

<metadata>
keyA: valueA
keyB: valueB
</metadata>
	`

	els, err := Parse([]byte(in))
	r.NoError(err)
	fmt.Printf(">>>TODO parsers/md/parser_test.go:48: els (%T): %v \n", els, els)
	el, ok := htm.Find("metadata", els)
	r.True(ok)
	r.NotNil(el)
	fmt.Printf(">>>TODO parsers/md/parser_test.go:52: el (%T): %v \n", el, el)
	m, ok := el.(remark.Metadata)
	r.True(ok)

	r.Equal("valueA", m["keyA"])
	r.Equal("valueB", m["keyB"])
}
