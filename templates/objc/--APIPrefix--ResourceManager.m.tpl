{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}ResourceManager.h"
#import <AFNetworking/AFNetworking.h>

@interface {{.Config.APIPrefix}}ResourceManager()
@property (nonatomic, strong) AFHTTPSessionManager *sessionManager;
@end

@implementation {{.Config.APIPrefix}}ResourceManager

- (instancetype)initWithBaseURL:(NSString *)baseURL
{
    if (self = [super init])
    {
        _baseURL = baseURL;
        _sessionManager = [[AFHTTPSessionManager alloc] initWithBaseURL:[NSURL URLWithString:baseURL]];
        _sessionManager.responseSerializer = [AFJSONResponseSerializer serializer];
        _sessionManager.requestSerializer = [AFJSONRequestSerializer serializer];
    }
    return self;
}

- (AnyPromise *)getResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params
                         modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance
{
    {{- block "resourceManagerRequestPromiseCreation" "GET"}}
    typeof (self) __weak weakSelf = self;
    return [self doRequest:^AnyPromise * {
                 typeof (self) __strong strongSelf = weakSelf;
                 PMKResolver resolver;
                 AnyPromise *requestPromise = [[AnyPromise alloc] initWithResolver:&resolver];
                 [strongSelf.sessionManager {{.}}:@""
                                     parameters:params
                                     {{- if or (eq . "GET") (eq . "POST")}}
                                       progress:nil
                                     {{- end}}
                                        success:^(NSURLSessionDataTask * _Nonnull task, id  _Nullable responseObject) {
                                            resolver(responseObject);
                                        }
                                        failure:^(NSURLSessionDataTask * _Nullable task, NSError * _Nonnull error) {
                                            resolver(error);
                                        }];
                 return requestPromise;
             }
             modelInstance:modelInstance];
    {{- end}}
}

- (AnyPromise *)postResourceWithURLPath:(NSString *)urlPath
                                 params:(NSDictionary *)params
                          modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance
{
    {{- template "resourceManagerRequestPromiseCreation" "POST"}}
}

- (AnyPromise *)putResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params
                         modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance
{
    {{- template "resourceManagerRequestPromiseCreation" "PUT"}}
}

- (AnyPromise *)deleteResourceWithURLPath:(NSString *)urlPath
                                   params:(NSDictionary *)params
                            modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance
{
    {{- template "resourceManagerRequestPromiseCreation" "DELETE"}}
}

#pragma mark - Private methods

- (AnyPromise *)doRequest:(AnyPromise *(^)())requestBlock modelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance
{
    typeof (self) __weak weakSelf = self;
    return requestBlock()
    .then(^(id response) {
        return [weakSelf parseResponse:response withModelInstance:modelInstance];
    })
    .catch(^(NSError *error) {
        // TODO: Check unauthorized, use refresh token and retry;
        return error;
    });
}

- (id)parseResponse:(id)response withModelInstance:(id<{{.Config.APIPrefix}}SerializableModel> (^)())modelInstance
{
    // TODO
    return modelInstance();
}

@end
