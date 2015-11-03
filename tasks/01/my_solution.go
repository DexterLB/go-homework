package main

import "strings"

// ExtractColumn gets specified column from each line
func ExtractColumn(logContents string, column uint8) (columnContents string) {
	columnItems := []string{}
	for _, line := range strings.Split(logContents, "\n") {
		columns := strings.Split(line, " ")
		if column == 0 && len(columns) > 1 {
			columnItems = append(columnItems, columns[0]+" "+columns[1])
		} else if column == 1 && len(columns) > 2 {
			columnItems = append(columnItems, columns[2])
		} else if column == 2 && len(columns) > 3 {
			columnItems = append(columnItems, strings.Join(columns[3:], " "))
		}
	}
	if len(columnItems) == 0 {
		return ""
	}
	return strings.Join(columnItems, "\n") + "\n"
}
