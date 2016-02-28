{{template "preHeaderComment" .}}
{{- $model := .CurrentModelInfo}}

#import "{{$model.Name}}Service.h"
#import "{{.CurrentModelInfo.Name}}.h"
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
        segmentParams[@"{{.}}"] = {{. | sanitizeVariable | singular | camelCase}};
    {{end -}}
    NSString *urlPath = [TTURLHelper replaceSegmentParams:segmentParams inURL:@"{{.URLPath}}"];
    {{- else -}}
    NSString *urlPath = @"{{.URLPath}}";
    {{end}}
    
    [self.resourceManager {{.Method.String | lower}}ResourceWithURLPath:urlPath
                                          params:{{if .NeedsModelParam}}[{{.Model.OriginalName}} toDictionary]{{else}}nil{{end}}
                                   modelInstance:^id <TTSerializableModel>
                                   {
                                       return {{if .Method.String | eq "PUT"}}{{.Model.OriginalName}}{{else}}[{{.Model.Name}} new]{{end}};
                                   }];

    return nil;
}
{{end}}
@end