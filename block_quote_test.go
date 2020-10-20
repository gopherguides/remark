package remark

import (
	"testing"
)

func Test_BlockQuote_String(t *testing.T) {
	t.Parallel()
	compare(t, "block_quote")
}
