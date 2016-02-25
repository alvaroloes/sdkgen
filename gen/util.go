package gen

import (
	"regexp"
	"strings"
	"text/template"

	"github.com/jinzhu/inflection"
)

var camelCaseRegexp = regexp.MustCompile("[0-9A-Za-z]+")

var funcMap = template.FuncMap{
	"trim":strings.TrimSpace,
	"lowerFirst": func(s string) string {
		return strings.ToLower(s[:1]) + s[1:]
	},
	"upperFirst": func(s string) string {
		return strings.ToUpper(s[:1]) + s[1:]
	},
	"plural": func(s string) string {
		return inflection.Plural(s)
	},
	"singular": func(s string) string {
		return inflection.Singular(s)
	},
	"camelCase": func(s string) string {
		chunks := camelCaseRegexp.FindAllString(s, -1)
		for idx, val := range chunks {
			if idx > 0 {
				chunks[idx] = strings.Title(val)
			}
		}
		return strings.Join(chunks, "")
	},
}
