# Update API contracts

When adding a new version of an API contract to this directory:

1. Remove the `servers` section (by default pointing to the public API endpoint of Trexis)
2. Update the `.env` file with the path the the updated version
3. Run `task generate` to re-generate the contracts

New APIs will be available in the mock-server, but will respond with `method not implemented`.

In order to implement the new API:

* Update the generated `server/api_<<new>>_service.go` to implement the logic
* Create a new entity using `go run entgo.io/ent/cmd/ent init <<new>>` to implement the datamodel
* Create `generator/<<new>>.go` to implement the data generator logic
* Update `generator/generator.go` to wire the data generator into the overall model
