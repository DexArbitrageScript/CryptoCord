package utils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func GenerateTable(header []string, rows [][]string) string {
	// calculate the maximum length of each column
	colWidths := make([]int, len(header))
	for i, h := range header {
		colWidths[i] = utf8.RuneCountInString(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			cellWidth := utf8.RuneCountInString(cell)
			if cellWidth > colWidths[i] {
				colWidths[i] = cellWidth
			}
		}
	}

	// build the horizontal line that separates the header and rows
	hLine := "+"
	for _, w := range colWidths {
		hLine += strings.Repeat("-", w+2) + "+"
	}

	// build the header row
	headerRow := "|"
	for i, h := range header {
		headerRow += fmt.Sprintf(" %-*s |", colWidths[i], h)
	}
	headerRow += "\n"

	// build the data rows
	dataRows := ""
	for _, row := range rows {
		dataRow := "|"
		for i, cell := range row {
			dataRow += fmt.Sprintf(" %-*s |", colWidths[i], cell)
		}
		dataRows += dataRow + "\n"
	}

	// combine everything into a single string
	return "```\n" + hLine + "\n" + headerRow + hLine + "\n" + dataRows + hLine + "\n```"
}

func GeneratePlusTable(header []string, rows [][]string) string {
	// calculate the maximum length of each column
	colWidths := make([]int, len(header))
	for i, h := range header {
		colWidths[i] = utf8.RuneCountInString(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			cellWidth := utf8.RuneCountInString(cell)
			if cellWidth > colWidths[i] {
				colWidths[i] = cellWidth
			}
		}
	}

	// build the horizontal line that separates the header and rows
	hLine := "+"
	for _, w := range colWidths {
		hLine += strings.Repeat("-", w+2) + "+"
	}

	// build the header row
	headerRow := ""
	for i, h := range header {
		headerRow += fmt.Sprintf("│ %-*s ", colWidths[i], h)
	}
	headerRow += "│\n"

	// build the separator row
	separatorRow := "+"
	for _, w := range colWidths {
		separatorRow += strings.Repeat("─", w+2) + "+"
	}
	separatorRow += "\n"

	// build the data rows
	dataRows := ""
	for _, row := range rows {
		dataRow := ""
		for i, cell := range row {
			dataRow += fmt.Sprintf("│ %-*s ", colWidths[i], cell)
		}
		dataRows += dataRow + "│\n"
	}

	// combine everything into a single string
	return "```\n" + separatorRow + headerRow + separatorRow + dataRows + separatorRow + "\n```"
}

func GenerateAsciiTable(header []string, rows [][]string) string {
	// calculate the maximum length of each column
	colWidths := make([]int, len(header))
	for i, h := range header {
		colWidths[i] = utf8.RuneCountInString(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			cellWidth := utf8.RuneCountInString(cell)
			if cellWidth > colWidths[i] {
				colWidths[i] = cellWidth
			}
		}
	}

	// build the header row
	headerRow := ""
	for i, h := range header {
		headerRow += fmt.Sprintf("%-*s  ", colWidths[i], h)
	}
	headerRow += "\n"

	// build the horizontal line that separates the header and rows
	hLine := ""
	for _, w := range colWidths {
		hLine += strings.Repeat("-", w+2) + " "
	}
	hLine += "\n"

	// build the data rows
	dataRows := ""
	for _, row := range rows {
		dataRow := ""
		for i, cell := range row {
			dataRow += fmt.Sprintf("%-*s  ", colWidths[i], cell)
		}
		dataRows += dataRow + "\n"
	}

	// combine everything into a single string
	return "```\n" + headerRow + hLine + dataRows + "```"
}

func DiscordDoubleTable(header1 []string, rows1 [][]string, header2 []string, rows2 [][]string) string {
	// calculate the maximum length of each column in both tables
	colWidths1 := make([]int, len(header1))
	for i, h := range header1 {
		colWidths1[i] = utf8.RuneCountInString(h)
	}
	for _, row := range rows1 {
		for i, cell := range row {
			cellWidth := utf8.RuneCountInString(cell)
			if cellWidth > colWidths1[i] {
				colWidths1[i] = cellWidth
			}
		}
	}
	colWidths2 := make([]int, len(header2))
	for i, h := range header2 {
		colWidths2[i] = utf8.RuneCountInString(h)
	}
	for _, row := range rows2 {
		for i, cell := range row {
			cellWidth := utf8.RuneCountInString(cell)
			if cellWidth > colWidths2[i] {
				colWidths2[i] = cellWidth
			}
		}
	}

	// build the horizontal lines that separate the headers and rows
	hLine1 := "+"
	for _, w := range colWidths1 {
		hLine1 += strings.Repeat("-", w+2) + "+"
	}
	hLine2 := "+"
	for _, w := range colWidths2 {
		hLine2 += strings.Repeat("-", w+2) + "+"
	}

	// build the header rows
	headerRow1 := "|"
	for i, h := range header1 {
		headerRow1 += fmt.Sprintf(" %-*s |", colWidths1[i], h)
	}
	headerRow1 += "\n" + hLine1 + "\n"
	headerRow2 := "|"
	for i, h := range header2 {
		headerRow2 += fmt.Sprintf(" %-*s |", colWidths2[i], h)
	}
	headerRow2 += "\n" + hLine2 + "\n"

	// build the data rows
	dataRows1 := ""
	for _, row := range rows1 {
		dataRow1 := "|"
		for i, cell := range row {
			dataRow1 += fmt.Sprintf(" %-*s |", colWidths1[i], cell)
		}
		dataRows1 += dataRow1 + "\n"
	}
	dataRows2 := ""
	for _, row := range rows2 {
		dataRow2 := "|"
		for i, cell := range row {
			dataRow2 += fmt.Sprintf(" %-*s |", colWidths2[i], cell)
		}
		dataRows2 += dataRow2 + "\n"
	}

	// combine everything into a single string
	return "```\n" + headerRow1 + dataRows1 + hLine1 + "\n\n" + headerRow2 + dataRows2 + hLine2 + "\n```"
}
