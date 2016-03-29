{{template "preHeaderComment" .}}

#import "{{.Config.APIPrefix}}SerializableModelUtils.h"
#import "{{.Config.APIPrefix}}SerializableModelProtocol.h"


@implementation {{.Config.APIPrefix}}SerializableModelUtils

- (NSArray<id <{{.Config.APIPrefix}}SerializableModel>> *)parseResponse:(id)response asArrayOfModel:(Class)modelClass
{
    NSAssert([modelClass conformsToProtocol:@protocol({{.Config.APIPrefix}}SerializableModel)], @"The model class must conform {{.Config.APIPrefix}}SerializableModel protocol");
    NSAssert([response isKindOfClass:[NSArray class]], [@"Error while parsing the response as an array. It is not an array, it is a " stringByAppendingString:NSStringFromClass([response class])]);

    NSMutableArray *responseArray = [NSMutableArray new];
    for (id dictionary in response) {
        NSAssert([dictionary isKindOfClass:[NSDictionary class]], [@"Error while parsing the response as an array. Its elements must be dictionaries, but they are " stringByAppendingString:NSStringFromClass([dictionary class])]);
        [responseArray addObject:[self parseResponse:dictionary asModel:modelClass]];
    }
    return responseArray;
}

- (NSDictionary *)parseResponse:(id)response asDictionaryOfStringKeysAndValuesOfModel:(Class)modelClass
{
    NSAssert([modelClass conformsToProtocol:@protocol({{.Config.APIPrefix}}SerializableModel)], @"The model class must conform {{.Config.APIPrefix}}SerializableModel protocol");
    NSAssert([response isKindOfClass:[NSDictionary class]], [@"Error while parsing the response as a dictionary. It is not a dictionary, it is a " stringByAppendingString:NSStringFromClass([response class])]);

    NSMutableDictionary *responseDictionary = [NSMutableDictionary new];
    for (NSString *key in response) {
        id valueDictionary = response[key];
        NSAssert([valueDictionary isKindOfClass:[NSDictionary class]], [@"Error while parsing the response as a dictionary. Its values must be dictionaries, but they are " stringByAppendingString:NSStringFromClass([valueDictionary class])]);
        responseDictionary[key] = [self parseResponse:valueDictionary asModel:modelClass];
    }
    return responseDictionary;
}

- (id <{{.Config.APIPrefix}}SerializableModel>)parseResponse:(id)response asModel:(Class)modelClass
{
    NSAssert([modelClass conformsToProtocol:@protocol({{.Config.APIPrefix}}SerializableModel)], @"The model class must conform {{.Config.APIPrefix}}SerializableModel protocol");

    id<{{.Config.APIPrefix}}SerializableModel> instance = (id <{{.Config.APIPrefix}}SerializableModel>) [modelClass new];
    [self parseResponse:response updatingModel:instance];
    return instance;
}

- (void)parseResponse:(id)response updatingModel:(id <{{.Config.APIPrefix}}SerializableModel>)modelInstance
{
    NSAssert([response isKindOfClass:[NSDictionary class]], [@"Error while parsing the response as a model. It must be a dictionary, but it is a " stringByAppendingString:NSStringFromClass([response class])]);
    [modelInstance updateWithDictionary:response];
}


@end
