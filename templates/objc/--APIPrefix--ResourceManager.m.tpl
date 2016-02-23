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

- (AnyPromise *)getResourceWithURLPath:(NSString *)urlPath params:(NSDictionary *)params
{
    {{- block "resourceManagerRequestPromiseCreation" "GET"}}
    PMKResolver resolver;
    AnyPromise *requestPromise = [[AnyPromise alloc] initWithResolver:&resolver];

    [self.sessionManager {{.}}:@""
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

    return [self wrapRequestPromise:requestPromise];
    {{- end}}
}

- (AnyPromise *)postResourceWithURLPath:(NSString *)urlPath params:(NSDictionary *)params
{
    {{- template "resourceManagerRequestPromiseCreation" "POST"}}
}

- (AnyPromise *)putResourceWithURLPath:(NSString *)urlPath params:(NSDictionary *)params
{
    {{- template "resourceManagerRequestPromiseCreation" "PUT"}}
}

- (AnyPromise *)deleteResourceWithURLPath:(NSString *)urlPath params:(NSDictionary *)params
{
    {{- template "resourceManagerRequestPromiseCreation" "DELETE"}}
}

#pragma mark - Private methods

- (AnyPromise *)wrapRequestPromise:(AnyPromise *)requestPromise
{
    return requestPromise;
}

@end
