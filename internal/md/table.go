package md

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

func NewTableWriter(w io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(w)
	table.SetCenterSeparator("|")
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetAutoWrapText(false)
	return table
}
