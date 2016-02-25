{{template "preHeaderComment" .}}

#import "{{.CurrentModelInfo.Name}}.h"
{{- range .CurrentModelInfo.ModelDependencies}}
#import "{{.}}.h"
{{- end}}

@implementation {{.CurrentModelInfo.Name}}

- (instancetype)initWithDictionary:(NSDictionary *)dictionary
{
    if (self = [super init])
    {
        [self updateWithDictionary:dictionary];
    }
    return self;
}

- (void)updateWithDictionary:(NSDictionary *)dictionary
{
{{- range .CurrentModelInfo.Properties}}
    {{ if $.CurrentModelInfo.ModelDependencies | contains .Type -}}
        {{- if .IsArray}}
    NSMutableArray *itemsOf{{.Type}} = [NSMutableArray new];
    NSArray *itemDictionariesOf{{.Type}} = dictionary[@"{{.Name}}"];
    for (NSDictionary *itemDictionary in itemDictionariesOf{{.Type}})
    {
        [itemsOf{{.Type}} addObject:[[{{.Type}} alloc] initWithDictionary:itemDictionary]]];
    }
    self.{{.Name}} = itemsOf{{.Type}};
        {{else -}}
    self.{{.Name}} = [[{{.Type}} alloc] initWithDictionary:dictionary[@"{{.Name}}"]];
        {{- end}}
    {{- else -}}
        self.{{.Name}} = {{if eq .Type "BOOL"}}[dictionary[@"{{.Name}}"] boolValue]{{else}}dictionary[@"{{.Name}}"]{{end}};
    {{- end}}
{{- end}}
}

- (NSDictionary *)toDictionary
{
// TODO: Conversion to dictionary of other models or arrays.
    NSMutableDictionary *dictionary = [NSMutableDictionary dictionary];
{{- range .CurrentModelInfo.Properties}}
    {{if $.CurrentModelInfo.ModelDependencies | contains .Type}} Yeah! {{end}}
    dictionary[@"{{.Name}}"] = {{if eq .Type "BOOL"}}@(self.{{.Name}}){{else}}self.{{.Name}}{{end}};
{{- end}}

    return dictionary;
}

@end