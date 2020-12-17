package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Document_HTML(t *testing.T) {
	t.Skip()
	t.Parallel()
	r := require.New(t)

	doc, err := ParseFile("./internal/testdata/buffalo/module.md")
	r.NoError(err)

	r.NotNil(doc)
}
