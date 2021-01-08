package remark

type Element interface {
	Name() string
}

type Elemental interface {
	Elements() Elements
}

type Elements []Element

func (elems Elements) Name() string {
	if len(elems) == 1 {
		return elems[0].Name()
	}
	return ""
}
