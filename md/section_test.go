package md

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parse_NewSection(t *testing.T) {
	exp := `Intro to Go: Getting Started`

	t.Parallel()
	r := require.New(t)

	tags, err := ParseFile("../internal/testdata/week00/module.md")
	r.NoError(err)
	r.NotZero(tags)

	r.Contains(tags.String(), exp)
}
