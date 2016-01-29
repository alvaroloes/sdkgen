{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

@interface {{.CurrentModelInfo.Name}} : NSObject
{{/* range .CurrentModelInfo.Properties -}}
@property (nonatomic, copy) {{.Type}}{{.Name}};
{{end - */}}
@end