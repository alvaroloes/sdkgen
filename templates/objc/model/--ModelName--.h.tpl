{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
{{- range .CurrentModelInfo.ModelDependencies}}
@class {{.}};
{{- end}}
#import "{{.Config.APIPrefix}}SerializableModelProtocol.h"

@interface {{.CurrentModelInfo.Name}} : NSObject <{{.Config.APIPrefix}}SerializableModel>
{{range .CurrentModelInfo.Properties -}}
@property (nonatomic) {{.TypeLabel}}{{.Name}};
{{end -}}
@end