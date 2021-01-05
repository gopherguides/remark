package htm

import (
	"reflect"

	"github.com/gopherguides/remark"
	"golang.org/x/net/html"
)

func set(el remark.Element, node *html.Node) error {
	rv := reflect.ValueOf(el).Elem()

	f := rv.FieldByName("Attributes")
	if f.IsValid() {
		if f.IsNil() {
			f.Set(reflect.MakeMap(f.Type()))
		}
		for _, at := range node.Attr {
			key := reflect.ValueOf(at.Key)
			val := reflect.ValueOf(at.Val)
			f.SetMapIndex(key, val)
		}
	}

	f = rv.FieldByName("Elements")
	if f.IsValid() {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			e, err := NewElement(c)
			if err != nil {
				return err
			}
			f.Set(reflect.Append(f, reflect.ValueOf(e)))
		}

	}
	return nil
}
