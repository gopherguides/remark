package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parser(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	f, err := os.Open("./example.md")
	r.NoError(err)
	defer f.Close()

	parser, err := ParseReader("", f, nil)
	r.NoError(err)
	r.NotNil(parser)
}
