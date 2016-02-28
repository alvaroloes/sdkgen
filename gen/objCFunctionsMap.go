package gen

import (
	"strings"
	"text/template"
)

const prefixForInvalidVarNames = "param"

var invalidVarNames = map[string]struct{}{
	"id": {},
}

var objCFuncMap = template.FuncMap{
	"sanitizeVariable": func(varName string) string {
		if _, invalid := invalidVarNames[varName]; invalid {
			return prefixForInvalidVarNames + strings.Title(varName)
		}
		return varName
	},
}
