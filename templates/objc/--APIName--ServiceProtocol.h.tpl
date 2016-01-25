{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}ResourceManager.h"

@protocol {{.Config.APIPrefix}}Service <NSObject>

/**
 * Creates an instance initialized and configured with the passed resource manager
 */
+ (instancetype)serviceWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager;

@end
