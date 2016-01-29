{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

@interface {{.Config.APIPrefix}}ResourceManager : NSObject

@property (nonatomic, copy) NSString *baseURL;

- (instancetype)initWithBaseURL:(NSString *)baseURL;

@end
