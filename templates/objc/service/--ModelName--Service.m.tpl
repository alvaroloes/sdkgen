{{template "preHeaderComment" .}}
{{- $model := .CurrentModelInfo}}

#import "{{$model.Name}}Service.h"
{{ range $dep, $_ := .CurrentModelInfo.EndpointsDependencies }}
#import "{{$dep.Name}}.h"
{{- end}}
#import "{{.Config.APIPrefix}}ResourceManager.h"
#import "{{.Config.APIPrefix}}URLHelper.h"
#import "{{.Config.APIPrefix}}SerializableModelUtils.h"

@interface {{$model.Name}}Service ()
@property (nonatomic, weak) {{.Config.APIPrefix}}ResourceManager *resourceManager;
@end

@implementation {{$model.Name}}Service

+ (instancetype)serviceWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager
{
    {{$model.Name}}Service *service = [[{{$model.Name}}Service alloc] init];
    if (service != nil)
    {
        service.resourceManager = resourceManager;
    }
    return service;
}
{{range $model.EndpointsInfo -}}
{{template "serviceMethodName" .}}
{
    {{if .SegmentParams -}}
    NSMutableDictionary *segmentParams = [NSMutableDictionary new];
    {{range .SegmentParams -}}
        segmentParams[@"{{.}}"] = {{. | sanitizeVariable | singular | camelCase}};
    {{end -}}
    NSString *urlPath = [{{$.Config.APIPrefix}}URLHelper replaceSegmentParams:segmentParams inURL:@"{{.URLPath}}"];
    {{- else -}}
    NSString *urlPath = @"{{.URLPath}}";
    {{- end}}

    {{- if .URLQueryParams | and .NeedsModelParam }}
    urlPath = [urlPath stringByAppendingString:[{{$.Config.APIPrefix}}URLHelper encodeQueryStringFromDictionary:query]];
    {{- end}}

    return [self.resourceManager {{.Method.String | lower}}ResourceWithURLPath:urlPath
                                                 params:{{if .NeedsModelParam -}}
                                                            [{{.RequestModel.OriginalName | lowerFirst}} toDictionary]
                                                        {{- else -}}
                                                            {{if .URLQueryParams }}query{{else}}nil{{end}}
                                                        {{- end}}]{{if .HasResponse}}
    .then(^(id response) {
        {{if .IsArrayResponse -}}
            return [{{$.Config.APIPrefix}}SerializableModelUtils parseResponse:response asArrayOfModel:[{{.ResponseModel.Name}} class]];
        {{else if .IsMapResponse -}}
            return [{{$.Config.APIPrefix}}SerializableModelUtils parseResponse:response asDictionaryOfStringKeysAndValuesOfModel:[{{.ResponseModel.Name}} class]];
        {{else if .IsModelResponse -}}
            {{if eq .Method.String "PUT" | and .NeedsModelParam | and (eq .RequestModel.Name .ResponseModel.Name) -}}
                return [{{$.Config.APIPrefix}}SerializableModelUtils parseResponse:response updatingModel:{{.RequestModel.OriginalName | lowerFirst}}];
            {{- else -}}
                return [{{$.Config.APIPrefix}}SerializableModelUtils parseResponse:response asModel:[{{.ResponseModel.Name}} class]];
            {{- end}}
        {{else if .IsRawResponse | or .IsRawArrayResponse | or .IsRawMapResponse -}}
            return response;
        {{- end}}
    });
    {{- else}};{{end}}
}
{{end}}
@end