{{template "preHeaderComment" .}}

#import "{{.Config.APIName}}.h"
{{- /*TODO: #import <AFOAuth2Manager.h>*/}}

static NSString * s{{.Config.APIPrefix}}BaseURL = nil;

@implementation {{.Config.APIName}}

+ (void) useBaseURLString:(NSString *)urlString
{
    s{{.Config.APIPrefix}}BaseURL = urlString;
}

- (instancetype)init
{
    if (self = [super init])
    {
        _resourceManager = [[{{.Config.APIPrefix}}ResourceManager alloc] initWithBaseURL:s{{.Config.APIPrefix}}BaseURL];
    }
    return self;
}

+ (instancetype)sharedInstance
{
    assert(s{{.Config.APIPrefix}}BaseURL != nil);
    {{$apiVar := lowerFirst .Config.APIName}}
    static {{.Config.APIName}} *{{$apiVar}} = nil;
    static dispatch_once_t onceToken;
    dispatch_once(&onceToken, ^{
        {{$apiVar}} = [[self alloc] init];
    });
    return {{$apiVar}};
}

- (void)setGlobalErrorHandlerWithBlock:(void (^)(NSError *))block
{
	PMKUnhandledErrorHandler = ^void(NSError *error)
	{
		dispatch_async(dispatch_get_main_queue(), ^
		{
			block(error);
		});
	};
}

@end
