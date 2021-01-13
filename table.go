package remark

import "github.com/gobuffalo/tags/v3"

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

func (g Table) String() string {
	t := tags.New(g.Name(), g.Options())
	t.Append(g.Elems)
	return t.String()
}
