{{template "preHeaderComment" .}}

#import <Foundation/Foundation.h>

@protocol {{.Config.APIPrefix}}SerializableModel;

@interface {{.Config.APIPrefix}}SerializableModelUtils : NSObject

+ (NSArray<id<{{.Config.APIPrefix}}SerializableModel>> *)parseResponse:(id)response asArrayOfModel:(Class)modelClass;
+ (NSDictionary *)parseResponse:(id)response asDictionaryOfStringKeysAndValuesOfModel:(Class)modelClass;
+ (id<TTSerializableModel>)parseResponse:(id)response asModel:(Class)modelClass;
+ (void)parseResponse:(id)response updatingModel:(id<{{.Config.APIPrefix}}SerializableModel>)modelInstance;

@end
