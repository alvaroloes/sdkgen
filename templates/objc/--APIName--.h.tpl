{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}ResourceManager.h"
{{- /*TODO: #import " AuthenticationManager.h" */}}
{{- range .AllModelsInfo }}
#import "{{.Name}}.h"
{{- end}}

@interface {{.Config.APIName}} : NSObject
{{- /*TODO: @property (nonatomic, strong) {{.Config.APIPrefix}}AuthenticationManager *authenticationManager;*/}}

/**
 *  Overrides the {{.Config.APIName}} SDK base url
 */
+ (void)useBaseURLString:(NSString *)urlString;

/**
 *  Returns the ResourceManager used in all requests
 */
+ ({{.Config.APIPrefix}}ResourceManager *)resourceManager;

/**
 * Sets an error handler block that will be executed always when any error occurs.
 * It will be executed after any other error handlers attached to the performed request.
 */
+ (void)setDefaultErrorHandlerWithBlock:(void (^)(NSError *error))block;

@end
