package remark

import "testing"

func Test_HTMLSpan_String(t *testing.T) {
	t.Parallel()
	compare(t, "html_span")
}
