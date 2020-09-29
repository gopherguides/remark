package ast

type String string

func (t String) String() string {
	return string(t)
}
