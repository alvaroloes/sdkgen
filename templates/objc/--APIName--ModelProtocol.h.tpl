{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}ResourceManager.h"

@protocol {{.Config.APIPrefix}}Model <NSObject>

/**
 * Creates an instance initialized and configured with the passed resource manager
 */
+ (instancetype)modelWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager;

@end
