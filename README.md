# SDKGen
Still in an early stage of development. Be patient ;-)

## TODO
- [ ] Don't create a service class if the corresponding model doesn't have any endpoints
- [ ] Accept different response types, not always the resource (for example a DELETE endpoint usually return nothing)
- [ ] Allow API versioning
- [ ] Allow property tuning with a key=value after the colon. Something like: `"prop1: name=desiredName, type=desiredType"`
- [ ] Right now, when a property value is a map, it is generated as a class. Allow it to be generated just as a map/dictionary
- [ ] Allow the specification of query parameters per each endpoint
- [ ] How to detect enum values from the API spec?
- [ ] Allow model tuning in the same way than property tuning
- [ ] JSON arrays of arrays may not be properly handled