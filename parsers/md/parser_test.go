package md

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parser_Parse(t *testing.T) {
	// t.Skip()
	t.Parallel()
	r := require.New(t)
	b, err := ioutil.ReadFile("./internal/testdata/full.md")
	r.NoError(err)

	m, err := Parse(b)
	r.NoError(err)
	r.NotNil(m)

}
