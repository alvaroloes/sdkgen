{{define "serviceMethodName" -}}

{{$modelNameUpper := upperFirst .Model.OriginalName}}
- (AnyPromise *){{.CRUDMethodName}}{{if .IsArrayResponse}}{{plural $modelNameUpper}}{{else}}{{$modelNameUpper}}{{end}}

{{- if .NeedsModelParam -}}
    :({{.Model.Name}} *){{.Model.OriginalName}}
{{- end}}
{{- if .SegmentParams}}
    {{- if .NeedsModelParam}} with{{else}}With{{end}}
    {{- $n := len .SegmentParams -}}
    {{- $first := index .SegmentParams 0 | singular | camelCase -}}
    {{$first | upperFirst}}:(NSString *){{$first | sanitizeVariable -}}
    {{range $index, $param := .SegmentParams -}}
        {{if gt $index 0}} {{. | singular | camelCase}}:(NSString *){{. | sanitizeVariable | singular | camelCase}}{{end}}
    {{- end}}
{{- end}}
{{- if .URLQueryParams }}
    {{- if .SegmentParams }} query
    {{- else if .NeedsModelParam}} withQuery
    {{- else}}WithQuery
    {{- end -}}
    :(NSDictionary *)query
{{- end}}
{{- end}}