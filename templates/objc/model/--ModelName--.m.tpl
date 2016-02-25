{{template "preHeaderComment" .}}

#import "{{.CurrentModelInfo.Name}}.h"
{{- range .CurrentModelInfo.Dependencies}}
#import "{{.}}.h"
{{- end}}

@implementation {{.CurrentModelInfo.Name}}

- (void)fillWithDictionary:(NSDictionary *)dictionary
{
// TODO: Conversion from other dictionaries and models
{{- range .CurrentModelInfo.Properties}}
    self.{{.Name}} = {{if eq .Type "bool"}}[dictionary[@"{{.Name}}"] boolValue]{{else}}dictionary[@"{{.Name}}"]{{end}};
{{- end}}
}

- (NSDictionary *)toDictionary
{
// TODO: Conversion to dictionary of other models or arrays.
    NSMutableDictionary *dictionary = [NSMutableDictionary dictionary];

{{- range .CurrentModelInfo.Properties}}
    dictionary[@"{{.Name}}"] = {{if eq .Type "bool"}}@(self.{{.Name}}){{else}}self.{{.Name}}{{end}};
{{- end}}

    return dictionary;
}

@end