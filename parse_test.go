package remark

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parser_Parse(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	m, err := ParseFile("./internal/md/example.md")
	r.NoError(err)
	fmt.Println(m)
	r.True(false)
}
