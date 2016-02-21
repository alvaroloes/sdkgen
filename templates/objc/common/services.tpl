{{define "serviceMethodName" -}}

{{$modelNameUpper := upperFirst .Model.OriginalName}}
- (PMKPromise *){{.CRUDMethodName}}{{if .IsArrayResponse}}{{plural $modelNameUpper}}{{else}}{{$modelNameUpper}}{{end}}

{{- if .NeedsModelParam -}}
    :({{.Model.Name}} *){{.Model.OriginalName}}
    {{- if .SegmentParams}} with{{end}}
{{- end}}
{{- if .SegmentParams}}
    {{- if not .NeedsModelParam}}With{{end}}
    {{- $n := len .SegmentParams -}}
    {{- $first := index .SegmentParams 0 | singular | camelCase -}}
    {{$first | upperFirst}}:(NSString *){{$first -}}
    {{range $index, $param := .SegmentParams -}}
        {{if gt $index 0}} {{. | singular | camelCase}}:(NSString *){{. | singular | camelCase}}{{end}}
    {{- end}}
{{- end}}

{{- end}}