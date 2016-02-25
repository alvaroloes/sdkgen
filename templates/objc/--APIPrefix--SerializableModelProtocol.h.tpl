{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

@protocol {{.Config.APIPrefix}}SerializableModel <NSObject>

/**
 * Initializes this model with the properties contained in the dictionary
 */
- (instancetype)initWithDictionary:(NSDictionary *)dictionary;

/**
 * Updates this model with the properties contained in the dictionary
 */
- (void)updateWithDictionary:(NSDictionary *)dictionary;

/**
 * Creates a dictionary representation of this object. It's the counterpart of fillWithDictionary
 */
- (NSDictionary *)toDictionary;

@end
