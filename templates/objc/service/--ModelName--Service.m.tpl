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

    {{if .Authenticates}}typeof(self) __weak weakSelf = self;{{end}}
    return [self.resourceManager {{.Method.String | lower}}ResourceWithURLPath:urlPath
                                                 params:{{if .NeedsModelParam -}}
                                                            [{{.RequestModel.OriginalName | lowerFirst}} toDictionary]
                                                        {{- else -}}
                                                            {{if .URLQueryParams }}query{{else}}nil{{end}}
                                                        {{- end}}]
    {{- if .HasResponse}}
    .then(^(id response) {
        {{template "serviceParseResponse" dict "Endpoint" . "Config" $.Config}}
    });
    {{- else -}}
    ;
    {{- end}}
}
{{end}}
@end