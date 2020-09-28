package ast

type Paragraph struct {
	Text string
}

func (p Paragraph) String() string {
	return p.Text
}
