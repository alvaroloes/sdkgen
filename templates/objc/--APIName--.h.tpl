{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

// Services
{{- range .AllModelsInfo }}
#import "{{.Name}}Service.h"
{{- end}}

// Models
{{- range .AllModelsInfo }}
#import "{{.Name}}.h"
{{- end}}

//Protocols
#import "{{.Config.APIName}}ServiceProtocol.h"

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
 * Returns a properly initialized service of the class passed as parameter. It must
 * conform the {{.Config.APIPrefix}}Service protocol. An exception is thrown otherwise
 */
- (id<{{.Config.APIPrefix}}Service>)service:(Class<{{.Config.APIPrefix}}Service>)serviceClass;

/**
 * Sets an error handler block that will be executed always when any error occurs.
 * It will be executed after any other error handlers attached to the performed request.
 */
+ (void)setGlobalErrorHandlerWithBlock:(void (^)(NSError *error))block;

@end
