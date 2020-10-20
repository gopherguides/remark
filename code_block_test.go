package remark

import (
	"testing"
)

func Test_CodeBlock_String(t *testing.T) {
	t.Parallel()
	compare(t, "code_block")
}
