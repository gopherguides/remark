package rpm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ModuleFS(t *testing.T) {
	r := require.New(t)

	cab := ModuleFS(t)
	r.NotNil(cab)
	r.Len(cab, 3)
}
