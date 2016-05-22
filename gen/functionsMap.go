package gen

import (
	"regexp"
	"strings"
	"text/template"

	"github.com/jinzhu/inflection"
	"github.com/juju/errors"
)

var camelCaseRegexp = regexp.MustCompile("[0-9A-Za-z]+")

var funcMap = template.FuncMap{
	"trim":  strings.TrimSpace,
	"lower": strings.ToLower,
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
	"dict": func(values ...interface{}) (map[string]interface{}, error) {
		if len(values)%2 != 0 {
			return nil, errors.New("invalid dict call")
		}
		dict := make(map[string]interface{}, len(values)/2)
		for i := 0; i < len(values); i += 2 {
			key, ok := values[i].(string)
			if !ok {
				return nil, errors.New("dict keys must be strings")
			}
			dict[key] = values[i+1]
		}
		return dict, nil
	},
}
