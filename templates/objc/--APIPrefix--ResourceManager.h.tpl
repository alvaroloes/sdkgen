{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
#import <PromiseKit/PromiseKit.h>
#import "{{.Config.APIPrefix}}SerializableModelProtocol.h"

@interface {{.Config.APIPrefix}}ResourceManager : NSObject

@property (nonatomic, copy) NSString *baseURL;

- (instancetype)initWithBaseURL:(NSString *)baseURL;

- (AnyPromise *)getResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params;

- (AnyPromise *)postResourceWithURLPath:(NSString *)urlPath
                                 params:(NSDictionary *)params;

- (AnyPromise *)putResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params;

- (AnyPromise *)deleteResourceWithURLPath:(NSString *)urlPath
                                   params:(NSDictionary *)params;

@end
