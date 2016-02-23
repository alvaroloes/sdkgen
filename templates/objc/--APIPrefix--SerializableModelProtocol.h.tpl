{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

@protocol {{.Config.APIPrefix}}SerializableModel <NSObject>

/**
 * Updates this model with the properties contained in the dictionary
 */
- (void)fillWithDictionary:(NSDictionary *)dictionary;

/**
 * Creates a dictionary representation of this object. It's the counterpart of fillWithDictionary
 */
- (NSDictionary *)toDictionary;

@end
