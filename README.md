# go-api-template

Template for APIs developed in Go (Golang). Based on concepts from [DDD and Clean Architecture](https://medium.com/@mastanca/clean-architecture-ddd-a-mixed-approach-773ab4623e14).

## Dependencies
* [Gin](https://github.com/gin-gonic/gin)
* [JWT](https://github.com/dgrijalva/jwt-go)
* [Testify](https://github.com/stretchr/testify)

## Usage

Search everywhere in the project for the string ``` github.com/mastanca/go-api-template ``` and replace with your module.
Also look out for TODOs everywhere in the template

### Run

```shell script
make run
```

Will start an http server on port 8080 serving a ping under the path ````/api/path````

### Test

```shell script
make test
```