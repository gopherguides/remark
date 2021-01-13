package remark

type Table struct {
	Attributes
	Elems    Elements
	metadata Metadata
}

func (table *Table) Metadata() Metadata {
	if table.metadata == nil {
		table.metadata = Metadata{}
	}
	return table.metadata
}

func (table Table) Elements() Elements {
	return table.Elems
}

func (table Table) Name() string {
	return "table"
}
