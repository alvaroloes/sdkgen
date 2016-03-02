{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

// Services
{{- range .AllModelsInfo }}
{{- if .EndpointsInfo }}
#import "{{.Name}}Service.h"
{{- end}}
{{- end}}

// Models
{{- range .AllModelsInfo }}
{{- if .Properties }}
#import "{{.Name}}.h"
{{- end}}
{{- end}}

//Protocols
#import "{{.Config.APIPrefix}}ServiceProtocol.h"

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
