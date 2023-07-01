package go_template

import (
	"strings"
)

// TODO handle this more gracefully
func SnakeToCamel(str string) string {
	chunks := strings.Split(str, "_")

	result := ""
	for _, chunk := range chunks {
		result += strings.Title(strings.ToLower(chunk))
	}

	return result
}
