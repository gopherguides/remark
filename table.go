package remark

type Table struct {
	Attributes
	Elems Elements
}

func (table Table) Elements() Elements {
	return table.Elems
}

func (table Table) Name() string {
	return "table"
}
