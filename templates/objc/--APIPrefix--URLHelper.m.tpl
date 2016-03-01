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

+ (NSString *)encodeQueryStringFromDictionary:(NSDictionary *)dict
{
    NSMutableString *query = [NSMutableString stringWithString:@"?"];
    for (NSString *key in dict)
    {
        if (query.length > 1)
        {
            [query appendString:@"&"];
        }

        id value = dict[key];
        if ([value isKindOfClass:[NSArray class]])
        {
            NSMutableString *queryArrayPart = [NSMutableString new];
            for (id val in value)
            {
                if (queryArrayPart.length > 0)
                {
                    [queryArrayPart appendString:@"&"];
                }

                [queryArrayPart appendFormat:@"%@[]=%@", key, val];
            }
            [query appendString:queryArrayPart];
        }
        else
        {
            [query appendFormat:@"%@=%@", key, value];
        }
    }

    return query;
}

@end
