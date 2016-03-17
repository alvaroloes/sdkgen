{{template "preHeaderComment" .}}

#import "{{.CurrentModelInfo.Name}}.h"
{{ range $dep, $_ := .CurrentModelInfo.ModelDependencies}}
#import "{{$dep.Name}}.h"
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
    {{ if $.CurrentModelInfo.DependsOnModel .Type -}}
        {{- if .IsArray}}
    NSMutableArray *itemsOf{{.Type}} = [NSMutableArray new];
    NSArray *itemDictionariesOf{{.Type}} = dictionary[@"{{.Name}}"];
    for (NSDictionary *itemDictionary in itemDictionariesOf{{.Type}})
    {
        [itemsOf{{.Type}} addObject:[[{{.Type}} alloc] initWithDictionary:itemDictionary]];
    }
    self.{{.NameLabel | sanitizeProperty}} = itemsOf{{.Type}};
        {{else if .IsMap -}}
    NSMutableDictionary *dictionaryOf{{.Type}} = [NSMutableDictionary new];
    NSDictionary *rawDictionaryOf{{.Type}} = dictionary[@"{{.Name}}"];
    for (NSString *key in rawDictionaryOf{{.Type}})
    {
        dictionaryOf{{.Type}}[key] = [[{{.Type}} alloc] initWithDictionary:rawDictionaryOf{{.Type}}[key]];
    }
    self.{{.NameLabel | sanitizeProperty}} = dictionaryOf{{.Type}};
        {{else -}}
    self.{{.NameLabel | sanitizeProperty}} = [[{{.Type}} alloc] initWithDictionary:dictionary[@"{{.Name}}"]];
        {{- end}}
    {{- else -}}
        self.{{.NameLabel | sanitizeProperty}} = {{if eq .Type "BOOL"}}[dictionary[@"{{.Name}}"] boolValue]{{else}}dictionary[@"{{.Name}}"]{{end}};
    {{- end}}
{{- end}}
}

- (NSDictionary *)toDictionary
{
    NSMutableDictionary *dictionary = [NSMutableDictionary dictionary];
{{ range .CurrentModelInfo.Properties}}
    {{ if $.CurrentModelInfo.DependsOnModel .Type -}}
        {{- if .IsArray}}
    NSMutableArray *itemDictionariesOf{{.Type}} = [NSMutableArray new];
    for ({{.Type}} *item in self.{{.NameLabel | sanitizeProperty}})
    {
        [itemDictionariesOf{{.Type}} addObject:[item toDictionary]];
    }
    dictionary[@"{{.Name}}"] = itemDictionariesOf{{.Type}};
        {{- else if .IsMap -}}

    NSMutableDictionary *rawDictionaryOf{{.Type}} = [NSMutableDictionary new];
    for (NSString *key in self.{{.NameLabel | sanitizeProperty}})
    {
        {{.Type}} *item = self.{{.NameLabel | sanitizeProperty}}[key];
        rawDictionaryOf{{.Type}}[key] = [item toDictionary];
    }
    dictionary[@"{{.Name}}"] = rawDictionaryOf{{.Type}};
        {{- else -}}
    dictionary[@"{{.Name}}"] = [self.{{.NameLabel | sanitizeProperty}} toDictionary];
        {{- end}}
    {{- else -}}
        dictionary[@"{{.Name}}"] = {{if eq .Type "BOOL"}}@(self.{{.NameLabel | sanitizeProperty}}){{else}}self.{{.NameLabel | sanitizeProperty}}{{end}};
    {{- end}}
{{- end}}

    return dictionary;
}

@end