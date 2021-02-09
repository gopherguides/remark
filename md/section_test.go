package md

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parse_NewSection(t *testing.T) {
	exp := `<html><head></head> <body></body></html> I recently had a friend come to me and mention that their website I created for them years ago needed some updating.  After looking at it, it was clear it was time for an update.  This article will walk through the design decisions and technical challenges that I needed to address in bringing their site up to some modern standards.`

	t.Parallel()
	r := require.New(t)

	tags, err := ParseFile("../internal/testdata/buffalo/module.md")
	r.NoError(err)
	r.NotZero(tags)

	r.Equal(exp, tags.Overview())
}
