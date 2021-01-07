package remark

type Document struct {
	Elements Elements
}

func (d Document) Name() string {
	return "document"
}
