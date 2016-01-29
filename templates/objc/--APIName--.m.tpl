{{template "preHeaderComment" .}}

#import "{{.Config.APIName}}.h"

static NSString * const k{{.Config.APIPrefix}}BaseURL = @"TODO: This should be the default base url";

@interface {{.Config.APIName}} ()
@property (nonatomic, strong) {{.Config.APIPrefix}}ResourceManager *resourceManager;
@end

@implementation {{.Config.APIName}}

+ (instancetype)default
{
    {{- $apiVar := lowerFirst .Config.APIName}}
    static {{.Config.APIName}} *{{$apiVar}} = nil;
    static dispatch_once_t onceToken;
    dispatch_once(&onceToken, ^{
        {{$apiVar}} = [[self alloc] init];
    });
    return {{$apiVar}};
}

- (instancetype)init
{
    if (self = [super init])
    {
        _resourceManager = [[{{.Config.APIPrefix}}ResourceManager alloc] initWithBaseURL:k{{.Config.APIPrefix}}BaseURL];
    }
    return self;
}

- (void)useBaseURLString:(NSString *)baseURL
{
    self.resourceManager.baseURL = baseURL;
}

- (id<{{.Config.APIPrefix}}Service>)service:(Class)serviceClass
{
    NSAssert([serviceClass conformsToProtocol:@protocol({{.Config.APIPrefix}}Service)], @"The service class must conform {{.Config.APIPrefix}}Service protocol");
    return [serviceClass serviceWithResourceManager:self.resourceManager];
}

+ (void)setGlobalErrorHandlerWithBlock:(void (^)(NSError *))block
{
{{/*
	PMKUnhandledErrorHandler = ^void(NSError *error)
	{
		dispatch_async(dispatch_get_main_queue(), ^
		{
			block(error);
		});
	};
*/}}
}

@end
