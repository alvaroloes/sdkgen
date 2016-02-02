{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
{{- range .CurrentModelInfo.Dependencies}}
@class {{.}};
{{- end}}

@interface {{.CurrentModelInfo.Name}} : NSObject
{{range .CurrentModelInfo.Properties -}}
@property (nonatomic) {{.Type}}{{.Name}};
{{end -}}
@end