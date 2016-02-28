{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}URLHelper.h"

@implementation {{.Config.APIPrefix}}URLHelper

+ (NSString *)replaceSegmentParams:(NSDictionary *)params inURL:(NSString *)url
{
    if (url.length == 0)
    {
        return url;
    }

    NSMutableString *finalURL = [NSMutableString stringWithString:url];
    for (NSString *paramName in params)
    {
        NSString *paramValue = params[paramName];
        if (paramValue != nil)
        {
            NSString *wrappedParamName = [NSString stringWithFormat:@":%@", paramName];
            [finalURL replaceOccurrencesOfString:wrappedParamName
                                      withString:paramValue
                                         options:NSLiteralSearch
                                           range:NSMakeRange(0, finalURL.length)];
        }
    }

    return finalURL;
}

@end
