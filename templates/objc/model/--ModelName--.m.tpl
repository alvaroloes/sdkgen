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
        [itemsOf{{.Type}} addObject:[[{{.Type}} alloc] initWithDictionary:itemDictionary]];
    }
    self.{{.Name | sanitizeProperty}} = itemsOf{{.Type}};
        {{else if .IsMap -}}
    NSMutableDictionary *dictionaryOf{{.Type}} = [NSMutableDictionary new];
    NSDictionary *rawDictionaryOf{{.Type}} = dictionary[@"{{.Name}}"];
    for (NSString *key in rawDictionaryOf{{.Type}})
    {
        dictionaryOf{{.Type}}[key] = [[{{.Type}} alloc] initWithDictionary:rawDictionaryOf{{.Type}}[key]];
    }
    self.{{.Name | sanitizeProperty}} = dictionaryOf{{.Type}};
        {{else -}}
    self.{{.Name | sanitizeProperty}} = [[{{.Type}} alloc] initWithDictionary:dictionary[@"{{.Name}}"]];
        {{- end}}
    {{- else -}}
        self.{{.Name | sanitizeProperty}} = {{if eq .Type "BOOL"}}[dictionary[@"{{.Name}}"] boolValue]{{else}}dictionary[@"{{.Name}}"]{{end}};
    {{- end}}
{{- end}}
}

- (NSDictionary *)toDictionary
{
    NSMutableDictionary *dictionary = [NSMutableDictionary dictionary];
{{ range .CurrentModelInfo.Properties}}
    {{ if $.CurrentModelInfo.ModelDependencies | contains .Type -}}
        {{- if .IsArray}}
    NSMutableArray *itemDictionariesOf{{.Type}} = [NSMutableArray new];
    for ({{.Type}} *item in self.{{.Name | sanitizeProperty}})
    {
        [itemDictionariesOf{{.Type}} addObject:[item toDictionary]];
    }
    dictionary[@"{{.Name}}"] = itemDictionariesOf{{.Type}};
        {{- else if .IsMap -}}

    NSMutableDictionary *rawDictionaryOf{{.Type}} = [NSMutableDictionary new];
    for (NSString *key in self.{{.Name | sanitizeProperty}})
    {
        {{.Type}} *item = self.{{.Name | sanitizeProperty}}[key];
        rawDictionaryOf{{.Type}}[key] = [item toDictionary];
    }
    dictionary[@"{{.Name}}"] = rawDictionaryOf{{.Type}};
        {{- else -}}
    dictionary[@"{{.Name}}"] = [self.{{.Name | sanitizeProperty}} toDictionary];
        {{- end}}
    {{- else -}}
        dictionary[@"{{.Name}}"] = {{if eq .Type "BOOL"}}@(self.{{.Name | sanitizeProperty}}){{else}}self.{{.Name | sanitizeProperty}}{{end}};
    {{- end}}
{{- end}}

    return dictionary;
}

@end