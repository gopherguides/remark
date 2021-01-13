package htm

import (
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_Parser_Metadata(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	in := `
<metadata>
keyA: valueA
keyB: valueB
</metadata>
	`

	els, err := Parse([]byte(in))
	r.NoError(err)

	el, ok := Find("metadata", els)
	r.True(ok)
	r.NotNil(el)

	m, ok := el.(remark.Metadata)
	r.True(ok)

	r.Equal("valueA", m["keyA"])
	r.Equal("valueB", m["keyB"])
}

func Test_Parser_Metadata_Bad(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	in := `
<metadata>
keyA valueA
keyB:
</metadata>
	`

	_, err := Parse([]byte(in))
	r.Error(err)
}
