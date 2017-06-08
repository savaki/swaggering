# swag

[![GoDoc](https://godoc.org/github.com/savaki/swag?status.svg)](https://godoc.org/github.com/MarkSonghurst/swag)
[![Build Status](https://travis-ci.org/savaki/swag.svg?branch=master)](https://travis-ci.org/MarkSonghurst/swag)

```swag``` is a lightweight library to generate swagger JSON for Go projects.  
 
No code generation, no framework constraints, just a simple swagger definition.

This is a self-contained fork of the original ```swag``` package written by savaki (https://github.com/savaki/swag)
Thanks to savaki for starting the original package and to the other PR authors who's work I've merged in.

The original ```swag``` package was heavily geared towards generating Swagger for APIs using ```application/json``` Request bodies - this fork additionally supports Swagger generation for APIs which accept Request bodies using the ```application/x-www-form-urlencoded``` and ```multipart/form-data``` MIME types.

Please note there are some (relatively minor) package API differences from the original ```swag``` package.

## Installation

```bash
go get github.com/MarkSonghurst/swag
```


## Status

This package is currently a work in progress.


## License

At the time of writing, the original ```swag``` package contains no licence or copyright information. Out of respect to the original Author, this fork maintains that lack of clarity.
That said, as far as I'm concerned, any code I've written for this package is freely available and in the public domain.

## Concepts

```swag``` uses functional options to generate both the swagger endpoints and the swagger definition.  Where possible
```swag``` attempts to use reasonable defaults that may be overridden by the user.

### Endpoints

```swag``` provides a separate package, ```endpoint```, to generate swagger endpoints.  These endpoints can be passed
to the swagger definition generate via ```swag.Endpoints(...)```

In this simple example, we generate an endpoint to retrieve all pets.  The only required fields for an endpoint
are the method, path, and the summary.  

```go
allPets := endpoint.New("get", "/pet", "Return all the pets") 
```

However, it'll probably be useful if you include definitions of what ```GET /pet``` returns:

```go
allPets := endpoint.New("get", "/pet", "Return all the pets",
  endpoint.Response(http.StatusOk, Pet{}, "Successful operation"),
  endpoint.Response(http.StatusInternalServerError, Error{}, "Oops ... something went wrong"),
) 
```

Refer to the [godoc](https://godoc.org/github.com/MarkSonghurst/swag/endpoint) for a list of all the endpoint options

### Walk

As a convenience to users, ```*swagger.Api``` implements a ```Walk``` method to simplify traversal of all the endpoints.
See the complete example below for how ```Walk``` can be used to bind endpoints to the router.

```go
api := swag.New(
    swag.Title("Swagger Petstore"),
    swag.Endpoints(post, get),
)

// iterate over each endpoint, if we've defined a handler, we can use it to bind to the router.  We're using ```gin``
// in this example, but any web framework will do.
// 
api.Walk(func(path string, endpoint *swagger.Endpoint) {
    h := endpoint.Handler.(func(c *gin.Context))
    path = swag.ColonPath(path)
    router.Handle(endpoint.Method, path, h)
})
```

### Custom Types

For types implementing `json.Marshaler` whose JSON output does not match their Go types (such as `time.Time`),
it is possible to override the default scanning of types and define a property manually.

For example:

```go
type Person struct {
  Name string `json:"name"`
  Birthday time.Time `json:"birthday"`
}

func main() {
  RegisterCustomType(time.Time{}, Property{
    Type: "string",
    Format: "date-time",
  })
}
```

`time.Time` is automatically registered as a custom type.

## Complete Example

```go
func handlePet(w http.ResponseWriter, _ *http.Request) {
	// your code here
}

type Pet struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []string `json:"tags"`
}

func main() {
    // define our endpoints
    // 
    post := endpoint.New("post", "/pet", "Add a new pet to the store",
        endpoint.Handler(handle),
        endpoint.Description("Additional information on adding a pet to the store"),
        endpoint.Body(Pet{}, "Pet object that needs to be added to the store", true),
        endpoint.Response(http.StatusOK, Pet{}, "Successfully added pet"),
    )
    get := endpoint.New("get", "/pet/{petId}", "Find pet by ID",
        endpoint.Handler(handle),
        endpoint.Path("petId", "integer", "ID of pet to return", true),
        endpoint.Response(http.StatusOK, Pet{}, "successful operation"),
    )
    
    // define the swagger api that will contain our endpoints
    // 
    api := swag.New(
      swag.Title("Swagger Petstore"),
      swag.Endpoints(post, get),
    )
    
    // iterate over each endpoint and add them to the default server mux
    // 
    for path, endpoints := range api.Paths {
      http.Handle(path, endpoints)
    }
    
    // use the api to server the swagger.json file
    // 
    enableCors := true
    http.Handle("/swagger", api.Handler(enableCors))
    
    http.ListenAndServe(":8080", nil)
}
```

## Additional Examples

Examples for popular web frameworks can be found in the examples directory:

* [http.Server](_examples/builtin/main.go)
* [Echo](_examples/echo/main.go)
* [Gin](_examples/gin/main.go)
* [Gorilla](_examples/gorilla/main.go)
* [httprouter](_examples/httprouter/main.go)

