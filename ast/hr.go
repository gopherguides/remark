package ast

type HR struct {
	Text string
}

func (hr HR) String() string {
	return hr.Text + "\n"
}
