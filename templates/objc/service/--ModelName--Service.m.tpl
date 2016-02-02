{{template "preHeaderComment" .}}

#import "{{.CurrentModelInfo.Name}}Service.h"
#import "{{.Config.APIPrefix}}ResourceManager.h"

@interface {{.CurrentModelInfo.Name}}Service ()
@property (nonatomic, weak) {{.Config.APIPrefix}}ResourceManager *resourceManager;
@end

@implementation {{.CurrentModelInfo.Name}}Service

+ (instancetype)serviceWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager
{
    {{.CurrentModelInfo.Name}}Service *service = [[{{.CurrentModelInfo.Name}}Service alloc] init];
    if (service != nil)
    {
        service.resourceManager = resourceManager;
    }
    return service;
}

@end