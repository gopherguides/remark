package remark

type Section struct {
	Attributes
	Elems    Elements
	metadata Metadata
}

func (sec *Section) Metadata() Metadata {
	if sec.metadata == nil {
		sec.metadata = Metadata{}
	}
	return sec.metadata
}

func (s Section) Elements() Elements {
	return s.Elems
}

func (sec Section) Name() string {
	return "section"
}
