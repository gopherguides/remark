package remark

type Metadatable interface {
	Metadata() Metadata
}

type Metadata map[string]string

func (m Metadata) Name() string {
	return "metadata"
}
