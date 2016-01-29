{{template "preHeaderComment" .}}

#import "{{.Config.APIName}}.h"

@interface {{.CurrentModelInfo.Name}} : NSObject
{{range .CurrentModelInfo.Properties -}}
@property (nonatomic) {{.Type}}{{.Name}};
{{end -}}
@end