package gen

import (
	"text/template"
)

const suffixForInvalidVarNames = "Param"
const suffixForInvalidPropNames = "Property"

var invalidVarNames = map[string]struct{}{
	"id": {},
}

var invalidPropertyNames = map[string]struct{}{
	"description": {},
}

var objCFuncMap = template.FuncMap{
	"sanitizeVariable": func(varName string) string {
		if _, invalid := invalidVarNames[varName]; invalid {
			return varName + suffixForInvalidVarNames
		}
		return varName
	},
	"sanitizeProperty": func(propName string) string {
		if _, invalid := invalidPropertyNames[propName]; invalid {
			return propName + suffixForInvalidPropNames
		}
		return propName
	},
}
