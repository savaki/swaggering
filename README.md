# swag

[![GoDoc](https://godoc.org/github.com/miketonks/swag?status.svg)](https://godoc.org/github.com/miketonks/swag)
[![Build Status](https://travis-ci.org/miketonks/swag.svg?branch=master)](https://travis-ci.org/miketonks/swag)

```swag``` is a lightweight library to generate swagger json for Go projects.  
 
No code generation, no framework constraints, just a simple swagger definition.

```swag``` is heavily geared towards generating REST/JSON apis.


## Installation

```bash
go get github.com/miketonks/swag
```


## Status

This package should be considered a release candidate.  No further package changes are expected at this point. 


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

Refer to the [godoc](https://godoc.org/github.com/miketonks/swag/endpoint) for a list of all the endpoint options

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

### Supported struct tags

The struct tags defined bellow apply to both **scalar** strings and **arrays**

| Tag | Description | Example | 
| ------ | ------ | ------ |
| format | Specifies the format of the string. **Supported formats:** ```uuid``` | ```format:"uuid"``` |
| min_length | Specifies the minimum length of the string | ```min_length:"1"```| 
| max_length | Specifies the maximum lenght of the string | ```max_length:"10"``` |
| enum | Specifies possible values of the string | ```enum:"Read,Write,Delete,Update"``` |
| pattern | Specifies a regular expression template for the string value | ```pattern:"^\w+$"``` |
| default | Specifies the default value of the string | ```default:"Read"```|

The struct tags defined bellow apply to **numbers** (all formats)

| Tag | Description | Example |
| ------ | ------ | ------ |
| default | Specifies the default value of the number | ```default:"1"```|
| minimum | Specifies the minimum value of the number | ```minimum:"1"```|
| maximum | Specifies the maximum value of the number | ```maximum:"50"```|
| exclusive_minimum | Excludes the minimum boundary value | ```exclusive_minimum:"true"```|
| exclusive_mamimum | Excludes the maximum boundary value | ```exclusive_maximum:"true"```|

The struct tags defined bellow apply to **booleans**
| Tag | Description | Example |
| ------ | ------ | ------ |
| default | Specifies the default value of the boolean | ```default:"true"```|

**_Note:_** Enumeration using a format tag i.e ```format:"enum,Allow,Deny"``` is now **deprecated** and soon will be removed.

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

* [http.Server](examples/builtin/main.go)
* [Echo](examples/echo/main.go)
* [Gin](examples/gin/main.go)
* [Gorilla](examples/gorilla/main.go)
* [httprouter](examples/httprouter/main.go)

