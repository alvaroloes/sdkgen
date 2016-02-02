{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
#import "{{.Config.APIPrefix}}ServiceProtocol.h"

@interface {{.CurrentModelInfo.Name}}Service : NSObject <{{.Config.APIPrefix}}Service>

+ (instancetype)serviceWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager;
{{- $model := .CurrentModelInfo}}
{{range $model.EndpointsInfo -}}
{{$modelName := upperFirst $model.OriginalName}}
- (void){{.CRUDMethodName}}{{if .IsArrayResponse}}{{pluralize $modelName}}{{else}}{{$modelName}}{{end}};
{{- end}}

@end