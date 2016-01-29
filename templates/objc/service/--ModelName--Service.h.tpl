{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

@interface {{.CurrentModelInfo.Name}}Service : NSObject
{{/* range .CurrentModelInfo.EndpointsInfo}}
{{.}}
{{end */}}
@end