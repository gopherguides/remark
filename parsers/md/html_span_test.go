package md

import "testing"

func Test_HTMLSpan_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "html_span")
}
