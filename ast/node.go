package ast

import (
	"fmt"
)

type interfacer interface {
	Interface() interface{}
}

func flatten(ii []interface{}) []interface{} {
	var res []interface{}
	for _, i := range ii {
		if i == nil {
			continue
		}
		switch t := i.(type) {
		case []interface{}:
			res = append(res, flatten(t)...)
		// case Noop:
		default:
			res = append(res, i)
		}
	}

	return res
}

type Node interface {
	fmt.Stringer
}

type Nodes []Node

func NewNodes(i interface{}) (Nodes, error) {
	var states Nodes

	ii := flatten([]interface{}{i})

	for _, i := range ii {
		switch t := i.(type) {
		case Node:
			states = append(states, t)
		case []interface{}:
			st, err := NewNodes(i)
			if err != nil {
				return states, err
			}
			states = append(states, st)
		default:
			return nil, fmt.Errorf("expected Node got %T", i)
		}
	}

	return states, nil
}

func (t Nodes) String() string {
	return ""
}

func (t *Nodes) Append(s Node, err error) error {
	if err != nil {
		return err
	}
	(*t) = append(*t, s)
	return nil
}
