{{template "preHeaderComment" .}}

#import "{{.CurrentModelInfo.Name}}.h"
{{- range .CurrentModelInfo.Dependencies}}
#import "{{.}}.h"
{{- end}}

@implementation {{.CurrentModelInfo.Name}}

- (void)fillWithDictionary:(NSDictionary *)dictionary
{
// TODO: type conversions (specially for bool in the other method) nil checks?
// TODO: Conversion from other dictionaries and models
{{- range .CurrentModelInfo.Properties}}
    self.{{.Name}} = dictionary[@"{{.Name}}"];
{{- end}}
}

- (NSDictionary *)toDictionary
{
    NSMutableDictionary *dictionary = [NSMutableDictionary dictionary];
{{- range .CurrentModelInfo.Properties}}
    dictionary[@"{{.Name}}"] = self.{{.Name}};
{{- end}}
    return dictionary;
}

@end