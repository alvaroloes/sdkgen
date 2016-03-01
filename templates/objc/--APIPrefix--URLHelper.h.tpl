{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

@interface {{.Config.APIPrefix}}URLHelper : NSObject

+ (NSString *)replaceSegmentParams:(NSDictionary *)params inURL:(NSString *)url;

+ (NSString *)encodeQueryStringFromDictionary:(NSDictionary *)dict;

@end
