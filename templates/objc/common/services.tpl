{{define "serviceMethodName" -}}
{{$modelName := upperFirst .Model.OriginalName}}
- (void){{.CRUDMethodName}}{{if .IsArrayResponse}}{{pluralize $modelName}}{{else}}{{$modelName}}{{end}}
{{- if .SegmentParams }}With
    {{- $n := len .SegmentParams -}}
    {{range $index, $param := .SegmentParams -}}
        {{upperFirst .}}:(NSString *){{.}}{{if lt (add $index 1) $n}} {{end}}
    {{- end}}
{{- end}}
{{- end}}