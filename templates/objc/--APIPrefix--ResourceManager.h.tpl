{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>
#import <PromiseKit/PromiseKit.h>
#import "{{.Config.APIPrefix}}SerializableModelProtocol.h"

@interface {{.Config.APIPrefix}}ResourceManager : NSObject

@property (nonatomic, copy) NSString *baseURL;

- (instancetype)initWithBaseURL:(NSString *)baseURL;

- (AnyPromise *)getResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params
                         modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance;

- (AnyPromise *)postResourceWithURLPath:(NSString *)urlPath
                                 params:(NSDictionary *)params
                          modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance;

- (AnyPromise *)putResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params
                         modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance;

- (AnyPromise *)deleteResourceWithURLPath:(NSString *)urlPath
                                   params:(NSDictionary *)params
                            modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance;

@end
