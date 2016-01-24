{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}ResourceManager.h"

@protocol {{.Config.APIPrefix}}Model <NSObject>

+ (instancetype)modelWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager;

@end
