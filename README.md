# SDKGen
Still in an early stage of development. Be patient ;-)

## TODO
- [x] Sanitize property names too (for example "description")
- [x] Allow the specification of query parameters per each endpoint
- [ ] Token based authentication
- [x] Don't create a service class if the corresponding model doesn't have any endpoints
- [x] Don't create a model if it has no properties (This happens when a property of other model is an array of simple types)
- [ ] Use the request type to generate de method parameter, don't rely only on HTTP Methods
- [ ] Use de response type and generate the code accordingly. don't use always the resource (for example a DELETE endpoint usually return nothing)
- [ ] Allow API versioning
- [ ] Support for format specifiers at the end of the endpoint (.json)? (by simply ignore them for now)
- [ ] Allow property tuning with a key=value after the colon. Something like: `"prop1: name=desiredName, type=desiredType"`
- [ ] Allow specifying Time type in properties (What format?).
- [ ] Allow model tuning in the same way than property tuning (taking into account the models whose name is taken from the resource endpoint)
- [ ] Right now, when a property value is a map, it is generated as a class. Allow it to be generated just as a map/dictionary
- [ ] How to detect enum values from the API spec?
- [ ] JSON arrays of arrays may not be properly handled
- [ ] Allow flagging some query parameters as method parameters (so they'll be treated similarly as segment parameters)
- [ ] Generate string constants for the query parameter names (or something similar)