package gen

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"capitalize": capitalize,
	"camelize":   camelize,
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

func camelize(s string) string {
	// TODO
	return s
}
