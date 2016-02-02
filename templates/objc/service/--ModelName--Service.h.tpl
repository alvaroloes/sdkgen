{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
#import "{{.Config.APIPrefix}}ServiceProtocol.h"

@interface {{.CurrentModelInfo.Name}}Service : NSObject <{{.Config.APIPrefix}}Service>

+ (instancetype)serviceWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager;
{{/* range .CurrentModelInfo.EndpointsInfo}}
{{.}}
{{end */}}
@end