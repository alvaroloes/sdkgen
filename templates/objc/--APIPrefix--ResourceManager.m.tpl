{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}ResourceManager.h"
#import <AFNetworking/AFNetworking.h>
{{if .AuthInfo -}}
#import <AFOAuthCredential>

static NSString *const kOAUTHCredentialIdentifier = @"{{.Config.APIPrefix}}OAUTHCredentialIdentifier";
{{end}}
@interface {{.Config.APIPrefix}}ResourceManager()
@property (nonatomic, strong) AFHTTPSessionManager *sessionManager;
@property (nonatomic, strong) AFOAuthCredential *credential;
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
        {{if .AuthInfo -}}
        _credential = [AFOAuthCredential retrieveCredentialWithIdentifier:kOAUTHCredentialIdentifier];
        if (_credential != nil)
        {
            [self.sessionManager.requestSerializer setValue:[NSString stringWithFormat:@"%@ %@", _credential.tokenType, _credential.accessToken]
                                         forHTTPHeaderField:@"Authorization"];
        }
        {{- end}}
    }
    return self;
}
{{if .AuthInfo}}
{{$modelVar := .AuthInfo.Endpoint.ResponseModel.OriginalName | lowerFirst}}
- (void)update{{.AuthInfo.Endpoint.ResponseModel.OriginalName | upperFirst}}:({{.AuthInfo.Endpoint.ResponseModel.Name}} *){{$modelVar}}
{
    self.credential = [AFOAuthCredential credentialWithOAuthToken:{{$modelVar}}.{{.AuthInfo.AccessTokenProp}}
                                                        tokenType:{{$modelVar}}.{{.AuthInfo.TokenTypeProp}}];
    {{if .AuthInfo.RefreshTokenProp -}}
    [self.credential setRefreshToken:{{$modelVar}}.{{.AuthInfo.RefreshTokenProp}}];
    {{end -}}
    [self.sessionManager.requestSerializer setValue:[NSString stringWithFormat:@"%@ %@", self.credential.tokenType, self.credential.accessToken]
                                 forHTTPHeaderField:@"Authorization"];
    [AFOAuthCredential storeCredential:self.credential withIdentifier:kOAUTHCredentialIdentifier];
}
{{end}}
- (AnyPromise *)getResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params
{
    {{- block "resourceManagerRequestPromiseCreation" "GET"}}
    typeof (self) __weak weakSelf = self;
    return [self doRequest:^AnyPromise * {
                 typeof (self) __strong strongSelf = weakSelf;
                 PMKResolver resolver;
                 AnyPromise *requestPromise = [[AnyPromise alloc] initWithResolver:&resolver];
                 [strongSelf.sessionManager {{.}}:urlPath
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
             }];
    {{- end}}
}

- (AnyPromise *)postResourceWithURLPath:(NSString *)urlPath
                                 params:(NSDictionary *)params;
{
    {{- template "resourceManagerRequestPromiseCreation" "POST"}}
}

- (AnyPromise *)putResourceWithURLPath:(NSString *)urlPath
                                params:(NSDictionary *)params;
{
    {{- template "resourceManagerRequestPromiseCreation" "PUT"}}
}

- (AnyPromise *)deleteResourceWithURLPath:(NSString *)urlPath
                                   params:(NSDictionary *)params;
{
    {{- template "resourceManagerRequestPromiseCreation" "DELETE"}}
}

#pragma mark - Private methods

- (AnyPromise *)doRequest:(AnyPromise *(^)())requestBlock
{
    {{/* typeof (self) __weak weakSelf = self; */}}
    return requestBlock()
    .then(^(id response) {
        // TODO: Log response
        return response;
    })
    .catch(^(NSError *error) {
        // TODO: Check unauthorized, use refresh token, retry, parse custom error.
        return error;
    });
}

@end
