package htm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parser_Image(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	in := `<img style="margin-right: 10px" align="left" src="./studio3.jpg" width="300px">`

	els, err := Parse([]byte(in))
	r.NoError(err)

	img, ok := Find("img", els)
	r.True(ok)

	s, err := Print(img)
	r.NoError(err)

	exp := `<img align="left" src="./studio3.jpg" style="margin-right: 10px" width="300px"></img>`
	r.Equal(exp, s)
}
