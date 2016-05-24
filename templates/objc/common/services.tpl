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

{{define "serviceParseResponse" -}}

{{if .Endpoint.Authenticates -}}
    {{.Endpoint.ResponseModel.Name}} *{{.Endpoint.ResponseModel.OriginalName | lowerFirst}} = [{{.Config.APIPrefix}}SerializableModelUtils parseResponse:response asModel:[{{.Endpoint.ResponseModel.Name}} class]];
        [weakSelf.resourceManager update{{.Endpoint.ResponseModel.OriginalName | upperFirst}}:{{.Endpoint.ResponseModel.OriginalName | lowerFirst}}];
        return {{.Endpoint.ResponseModel.OriginalName | lowerFirst}};
{{- else if .Endpoint.IsArrayResponse -}}
    return [{{.Config.APIPrefix}}SerializableModelUtils parseResponse:response asArrayOfModel:[{{.Endpoint.ResponseModel.Name}} class]];
{{- else if .Endpoint.IsMapResponse -}}
    return [{{.Config.APIPrefix}}SerializableModelUtils parseResponse:response asDictionaryOfStringKeysAndValuesOfModel:[{{.Endpoint.ResponseModel.Name}} class]];
{{- else if .Endpoint.IsModelResponse -}}
    {{if eq .Endpoint.Method.String "PUT" | and .Endpoint.NeedsModelParam | and (eq .Endpoint.RequestModel.Name .Endpoint.ResponseModel.Name) -}}
        return [{{.Config.APIPrefix}}SerializableModelUtils parseResponse:response updatingModel:{{.Endpoint.RequestModel.OriginalName | lowerFirst}}];
    {{- else -}}
        return [{{.Config.APIPrefix}}SerializableModelUtils parseResponse:response asModel:[{{.Endpoint.ResponseModel.Name}} class]];
    {{- end}}
{{- else if .Endpoint.IsRawResponse | or .Endpoint.IsRawArrayResponse | or .Endpoint.IsRawMapResponse -}}
    return response;
{{- end}}

{{- end}}