package remark

import (
	"testing"
)

func Test_CodeBlock_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "code_block")
}
