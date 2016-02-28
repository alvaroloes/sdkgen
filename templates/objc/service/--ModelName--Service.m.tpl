{{template "preHeaderComment" .}}
{{- $model := .CurrentModelInfo}}

#import "{{$model.Name}}Service.h"
#import "{{.Config.APIPrefix}}ResourceManager.h"
#import "{{.Config.APIPrefix}}URLHelper.h"

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
        segmentParams[@"{{.}}"] = {{. | singular | camelCase}};
    {{end -}}
    NSString *url = [TTURLHelper replaceSegmentParams:segmentParams inURL:@"{{.URLPath}}"];
    {{- else -}}
    NSString *url = @"{{.URLPath}}";
    {{end}}



    return nil;
}
{{end}}
@end