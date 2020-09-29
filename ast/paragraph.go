package ast

type Paragraph struct {
	Nodes Nodes
}

func (p Paragraph) String() string {
	return p.Nodes.String() + "\n"
}
