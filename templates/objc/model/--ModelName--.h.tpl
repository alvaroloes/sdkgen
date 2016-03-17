{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
#import "{{.Config.APIPrefix}}SerializableModelProtocol.h"
{{ range $dep, $_ := .CurrentModelInfo.ModelDependencies}}
@class {{$dep.Name}};
{{- end}}

@interface {{.CurrentModelInfo.Name}} : NSObject <{{.Config.APIPrefix}}SerializableModel>
{{range .CurrentModelInfo.Properties -}}
@property (nonatomic) {{.TypeLabel}}{{.NameLabel | sanitizeProperty}};
{{end -}}
@end