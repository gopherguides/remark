package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Paragraph_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const input = "Go promotes sharing memory by communicating. This chapter covers goroutines and how to synchronize communication between them. Mechanics for synchronization such as `WaitGroups` and `Mutexes` are explored along with the corresponding patterns for each."

	m := parseT(t, input)
	r.Equal(input, m.String())
}
