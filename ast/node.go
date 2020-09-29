package ast

import (
	"fmt"
	"strings"
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
		case []byte:
			states = append(states, String(string(t)))
		case string:
			states = append(states, String(t))
		default:
			return nil, fmt.Errorf("expected Node got %T", i)
		}
	}

	return states, nil
}

func (nodes Nodes) String() string {
	var lines []string
	for _, n := range nodes {
		lines = append(lines, n.String())
	}
	return strings.Join(lines, "\n")
}

func (t *Nodes) Append(s Node, err error) error {
	if err != nil {
		return err
	}
	(*t) = append(*t, s)
	return nil
}
