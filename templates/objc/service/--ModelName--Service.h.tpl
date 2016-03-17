{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
#import "{{.Config.APIPrefix}}ServiceProtocol.h"
#import <PromiseKit/PromiseKit.h>

{{ range $dep, $_ := .CurrentModelInfo.EndpointsDependencies }}
@class {{$dep.Name}};
{{- end}}

@interface {{.CurrentModelInfo.Name}}Service : NSObject <{{.Config.APIPrefix}}Service>

+ (instancetype)serviceWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager;
{{- $model := .CurrentModelInfo}}
{{range $model.EndpointsInfo -}}
{{template "serviceMethodName" .}};
{{- end}}

@end