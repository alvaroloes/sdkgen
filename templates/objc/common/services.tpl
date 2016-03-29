{{define "serviceMethodName" -}}

{{$resourceNameUpper := upperFirst .ResourceModel.OriginalName}}
// TODO <Add doc about the response type>
- (AnyPromise *){{.CRUDMethodName}}{{if .IsArrayResponse}}{{plural $resourceNameUpper}}{{else}}{{$resourceNameUpper}}{{end}}

{{- if .NeedsModelParam -}}
    :({{.RequestModel.Name}} *){{.RequestModel.OriginalName | lowerFirst}}
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