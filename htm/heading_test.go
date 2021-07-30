package htm

import (
	"fmt"
	"testing"

	"github.com/gopherguides/remark/internal/rpm"
	"github.com/stretchr/testify/require"
)

func Test_Parser_NewHeading(t *testing.T) {
	t.Parallel()

	p := testParser(t, rpm.ModuleFS(t))

	for _, name := range []string{"", "h", "h-1", "h6"} {
		t.Run(fmt.Sprintf("errors: %s", name), func(t *testing.T) {
			r := require.New(t)
			node := ElementNode(t, name)

			_, err := p.NewHeading(node)
			r.Error(err)
		})
	}

	for _, name := range HEADS.Allowed() {
		r := require.New(t)
		node := ElementNode(t, name)

		h, err := p.NewHeading(node)
		r.NoError(err)
		r.NotNil(h)
		r.Equal(name, h.Name)
	}

}
