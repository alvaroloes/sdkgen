package gen

import (
	"strings"
	"text/template"

	"github.com/jinzhu/inflection"
)

var funcMap = template.FuncMap{
	"lowerFirst": func(s string) string {
		return strings.ToLower(s[:1]) + s[1:]
	},
	"upperFirst": func(s string) string {
		return strings.ToUpper(s[:1]) + s[1:]
	},
	"pluralize": func(s string) string {
		return inflection.Plural(s)
	},
}
