package utils

import "strings"

func StringToMultiline(str string, numPerLine int) string {
	if len(str) <= numPerLine {
		return str
	}
	lines := []string{}
	line := ""

	splitted := strings.Split(str, " ")
	for _, v := range splitted {
		tmp := line + " " + v
		if len(tmp) > numPerLine {
			lines = append(lines, tmp)
			line = ""
		} else {
			line = tmp
		}
	}
	lines = append(lines, line)
	return strings.Join(lines, "\n")
}
