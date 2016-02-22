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
    return nil;
}

- (AnyPromise *)postResourceWithURLPath:(NSString *)urlPath params:(NSDictionary *)params
{
    return nil;
}

- (AnyPromise *)putResourceWithURLPath:(NSString *)urlPath params:(NSDictionary *)params
{
    return nil;
}

- (AnyPromise *)deleteResourceWithURLPath:(NSString *)urlPath params:(NSDictionary *)params
{
    return nil;
}

@end
