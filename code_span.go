package remark

type CodeSpan struct {
	*HTMLSpan
}

func NewCodeSpan(html *HTMLSpan) (*CodeSpan, error) {
	c := &CodeSpan{
		HTMLSpan: html,
	}

	return c, nil
}
