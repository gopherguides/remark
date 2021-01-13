package htm

import (
	"reflect"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func (p *Parser) set(el remark.Element, node *html.Node) error {
	rv := reflect.ValueOf(el).Elem()

	f := rv.FieldByName("Attributes")

	if f.IsValid() {
		tt := reflect.New(f.Type())
		m := tt.MethodByName("Set")

		for _, at := range node.Attr {
			key := reflect.ValueOf(at.Key)
			val := reflect.ValueOf(at.Val)
			m.Call([]reflect.Value{key, val})
		}

		f.Set(reflect.Indirect(tt))
	}

	f = rv.FieldByName("Elems")
	if f.IsValid() {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			e, err := p.NewElement(c)
			if err != nil {
				return err
			}
			f.Set(reflect.Append(f, reflect.ValueOf(e)))
		}

	}

	return nil
}
