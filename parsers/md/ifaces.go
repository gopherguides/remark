package md

import "github.com/olekukonko/tablewriter"

type TableWriter interface {
	WriteTable(w *tablewriter.Table) error
}
