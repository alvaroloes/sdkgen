{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}ResourceManager.h"

@interface {{.Config.APIPrefix}}ResourceManager()

@end

@implementation {{.Config.APIPrefix}}ResourceManager

- (instancetype)initWithBaseURL:(NSString *)baseURL
{
    if (self = [super init])
    {
        _baseURL = baseURL;
    }
    return self;
}
@end
