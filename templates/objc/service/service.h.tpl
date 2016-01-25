/*
 * NOTE: This file has been auto-generated. Any manual changes will be overwritten
 * the next time the auto-generation is run
 */

@interface {{.CurrentModelInfo.Name}}Service : NSObject
{{range .CurrentModelInfo.EndpointsInfo}}
{{.}}
{{end}}
@end