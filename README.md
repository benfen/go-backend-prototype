# Go Backend Prototype

## Requires

* `go`
* `yarn`
* `docker`
* `bash`

## How to use

* Go to `web/` and run `yarn && yarn build` to create the frontend artifacts
* Run `./db.sh` from the root directory to start up a postgres docker container
* Run `go run server.go` from the root of the repo to start up the server; it will now be accessible at http://localhost:4000
