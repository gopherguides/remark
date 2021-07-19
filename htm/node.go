package htm

import (
	"encoding/json"

	"golang.org/x/net/html"
)

func NodeString(n *html.Node) string {
	ats := map[string]string{}
	mm := map[string]interface{}{
		"data":  n.Data,
		"attrs": ats,
	}

	if len(n.Namespace) > 0 {
		mm["namespace"] = n.Namespace
	}

	for _, at := range n.Attr {
		ats[at.Key] = at.Val
	}
	b, _ := json.Marshal(mm)
	return string(b)
}
