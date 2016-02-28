{{template "preHeaderComment" .}}

@class {{.Config.APIPrefix}}ResourceManager;

@protocol {{.Config.APIPrefix}}Service <NSObject>

/**
 * Creates an instance initialized and configured with the passed resource manager
 */
+ (instancetype)serviceWithResourceManager:({{.Config.APIPrefix}}ResourceManager *)resourceManager;

@end
