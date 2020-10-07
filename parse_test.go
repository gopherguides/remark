package remark

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parser_Parse_Simple(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const input = "# H1\n\nsome **text** with `code` goes here."

	m, err := Parse([]byte(input))
	r.NoError(err)
	fmt.Printf(">>>TODO parse_test.go:19: m (%T): %v \n", m, m)
	r.True(false)
}

func Test_Parser_Parse(t *testing.T) {
	t.Skip()
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
