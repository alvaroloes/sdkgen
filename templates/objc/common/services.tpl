{{define "serviceMethodName" -}}

{{$modelName := upperFirst .Model.OriginalName}}
- (void){{.CRUDMethodName}}{{if .IsArrayResponse}}{{plural $modelName}}{{else}}{{$modelName}}{{end}}
{{- if .SegmentParams }}With
    {{- $n := len .SegmentParams -}}
    {{- $first := index .SegmentParams 0 | singular | camelCase -}}
    {{$first | upperFirst}}:(NSString *){{$first -}}
    {{range $index, $param := .SegmentParams -}}
        {{if gt $index 0}} {{. | singular | camelCase}}:(NSString *){{. | singular | camelCase}}{{end}}
    {{- end}}
{{- end}}

{{- end}}