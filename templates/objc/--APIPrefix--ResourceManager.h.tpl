{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
#import <PromiseKit/PromiseKit.h>
#import "{{.Config.APIPrefix}}SerializableModelProtocol.h"
{{if .AuthInfo -}}
#import "{{.AuthInfo.Endpoint.ResponseModel.Name}}.h"
{{- end}}

@interface {{.Config.APIPrefix}}ResourceManager : NSObject

@property (nonatomic, copy) NSString *baseURL;

- (instancetype)initWithBaseURL:(NSString *)baseURL;
{{if .AuthInfo}}
- (void)update{{.AuthInfo.Endpoint.ResponseModel.OriginalName | upperFirst}}:({{.AuthInfo.Endpoint.ResponseModel.Name}} *){{.AuthInfo.Endpoint.ResponseModel.OriginalName | lowerFirst}};
{{end}}
- (AnyPromise *)getResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params;

- (AnyPromise *)postResourceWithURLPath:(NSString *)urlPath
                                 params:(NSDictionary *)params;

- (AnyPromise *)putResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params;

- (AnyPromise *)deleteResourceWithURLPath:(NSString *)urlPath
                                   params:(NSDictionary *)params;

@end
