/*
 * NOTE: This file has been auto-generated. Any manual changes will be overwritten
 * the next time the auto-generation is run
 */

@interface {{.CurrentModelInfo.Name | capitalize}} : NSObject
{{range .CurrentModelInfo.Properties -}}
@property (nonatomic, copy) {{.Type}} *{{.Name}};
{{end -}}
@end