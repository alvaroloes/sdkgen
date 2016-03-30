# SDKGen
Still in an early stage of development. Be patient ;-)

## TODO
- [x] Sanitize property names too (for example "description")
- [x] Allow the specification of query parameters per each endpoint
- [x] Don't create a service class if the corresponding model doesn't have any endpoints
- [x] Don't create a model if it has no properties (This happens when a property of other model is an array of simple types)
- [x] Allow property tuning with a key=value after the colon. Something like: `"prop1: name=desiredName, type=desiredType"`
- [x] Right now, when a property value is a map, it is generated as a class. Allow it to be generated just as a map/dictionary (related with property tuning)
- [x] Allow model tuning in the same way than property tuning (taking into account the models whose name is taken from the resource endpoint)
- [ ] Allow specifying map and  rawMap/rawArray (maybe only needed "raw") type in properties and model tuning.


- [ ] Allow endpoint tuning (HTTP method -> crud method name override, resource -> model name part of service method override)
- [ ] Allow specifying Time type in properties (What format?).
- [ ] Token based authentication (think of a smart way to accomplish this. Maybe nothing is needed or simple a way to specify the headers that must be set in a general way)

- [ ] Update the readme

- [ ] Use 'RequestKind' (not relay on HTTP method, like "NeedsModelParam") in the same way as 'ResponseKind': this will allow to send different things (like an array of models to bulk update or a map)
- [ ] Support for format specifiers at the end of the endpoint (.json)? (by simply ignore them for now)
- [ ] How to detect enum values from the API spec?
- [ ] Allow flagging some query parameters as method parameters (so they'll be treated similarly as segment parameters)
- [ ] Generate string constants for the query parameter names (or something similar)
- [ ] Allow API versioning
- [ ] Allow send request with an array of objects.

- [ ] Allow non JSON responses like string or bool?
- [ ] Arrays of arrays with typed elements (not raw) are  not properly handled
- [ ] Arrays of maps with typed elements (not raw) are not properly handled
