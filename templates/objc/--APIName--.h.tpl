{{template "preHeaderComment" .}}
{{ range .AllModelsInfo }}
#import "{{.Name}}.h"
{{- end}}

@interface {{.Config.APIName}} : NSObject

/**
 * Default {{.Config.APIName}} SDK instance
 */
+ (instancetype)default;

/**
 *  Overrides the {{.Config.APIName}} SDK base url
 */
- (void)useBaseURLString:(NSString *)urlString;

/**
 * Returns a properly initialized model of the class passed as parameter. It must
 * conform the {{.Config.APIPrefix}}Model protocol. An exception is thrown otherwise
 */
- (id<{{.Config.APIPrefix}}Model>)model:(Class<{{.Config.APIPrefix}}Model>)modelClass;

/**
 * Sets an error handler block that will be executed always when any error occurs.
 * It will be executed after any other error handlers attached to the performed request.
 */
+ (void)setGlobalErrorHandlerWithBlock:(void (^)(NSError *error))block;

@end
