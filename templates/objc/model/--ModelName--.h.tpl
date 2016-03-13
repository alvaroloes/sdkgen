{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
#import "{{.Config.APIPrefix}}SerializableModelProtocol.h"
{{ range .CurrentModelInfo.ModelDependencies}}
@class {{.}};
{{- end}}

@interface {{.CurrentModelInfo.Name}} : NSObject <{{.Config.APIPrefix}}SerializableModel>
{{range .CurrentModelInfo.Properties -}}
@property (nonatomic) {{.TypeLabel}}{{.NameLabel | sanitizeProperty}};
{{end -}}
@end