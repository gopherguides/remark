package htm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_HTMLBody(t *testing.T) {
	r := require.New(t)

	src := `<include src="foo.go"></include>`
	root, err := html.Parse(strings.NewReader(src))

	r.NoError(err)
	r.NotNil(root)

	body, err := HTMLBody(root)
	r.NoError(err)
	r.NotNil(body)
	r.Equal("body", body.Data)
}
