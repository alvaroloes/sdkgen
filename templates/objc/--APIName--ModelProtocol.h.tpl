{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}ResourceManager.h"

@protocol {{.Config.APIPrefix}}Model <NSObject>

- (instancetype)initWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager;

@end
