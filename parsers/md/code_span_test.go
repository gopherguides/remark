package md

import "testing"

func Test_CodeSpan_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "code_span")
}
