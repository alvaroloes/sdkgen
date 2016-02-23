{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
{{- range .CurrentModelInfo.Dependencies}}
@class {{.}};
{{- end}}
#import "{{.Config.APIPrefix}}SerializableModelProtocol.h"

@interface {{.CurrentModelInfo.Name}} : NSObject <{{.Config.APIPrefix}}SerializableModel>
{{range .CurrentModelInfo.Properties -}}
@property (nonatomic) {{.Type}}{{.Name}};
{{end -}}
@end