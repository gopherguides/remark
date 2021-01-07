package md

import (
	"io/ioutil"
	"path/filepath"
	"testing"

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
