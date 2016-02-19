{{template "preHeaderComment" .}}
{{- $model := .CurrentModelInfo}}

#import "{{$model.Name}}Service.h"
#import "{{.Config.APIPrefix}}ResourceManager.h"

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
//TODO: The method name must include the object to update in the create/update HTTP methods
//TODO: url creation with resource object? with a simple URL? how to parse the result?
//{{.}}
}
{{end}}
@end