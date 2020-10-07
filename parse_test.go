package remark

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parser_Parse(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	b, err := ioutil.ReadFile("./internal/md/example.md")
	r.NoError(err)

	m, err := Parse(b)
	r.NoError(err)
	fmt.Println(m)
	r.True(false)
	// r.Equal(string(b), m.String())
}
