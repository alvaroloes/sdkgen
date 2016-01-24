package gen

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"toLower":strings.ToLower,
	"title": func(s string) string {
		return ""
	},
	"camelize": func(s string) string {
		return ""
	},
}
