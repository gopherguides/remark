package htm

import "github.com/gopherguides/remark"

func Find(name string, elems Elements) (remark.Element, bool) {
	for _, e := range elems {
		if e.Name() == name {
			return e, true
		}
		if el, ok := e.(remark.Elemental); ok {
			if x, ok := Find(name, el.Elements()); ok {
				return x, true
			}
		}
	}
	return nil, false
}
