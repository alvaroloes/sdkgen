package gen

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"lowerFirst":func(s string) string {
		return strings.ToLower(s[:1]) + s[1:]
	},
}
