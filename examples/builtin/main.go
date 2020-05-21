// Copyright 2017 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"io"
	"net/http"
	"time"

	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

func handle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, req.Method+" - Insert your code here")
}

// Category example from the swagger pet store
type Category struct {
	ID   int64  `json:"category"`
	Name string `json:"name"`
}

// Pet example from the swagger pet store
type Pet struct {
	ID        int64      `json:"id"`
	Category  Category   `json:"category"`
	Name      string     `json:"name"`
	PhotoUrls []string   `json:"photoUrls"`
	Tags      []string   `json:"tags"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func main() {
	post := endpoint.New("post", "/pet", "Add a new pet to the store",
		endpoint.Handler(handle),
		endpoint.Description("Additional information on adding a pet to the store"),
		endpoint.Body(Pet{}, "Pet object that needs to be added to the store", true),
		endpoint.Response(http.StatusOK, Pet{}, "Successfully added pet"),
		endpoint.Security("petstore_auth", "read:pets", "write:pets"),
	)
	get := endpoint.New("get", "/pet/{petId}", "Find pet by ID",
		endpoint.Handler(handle),
		endpoint.Path("petId", "integer", "ID of pet to return", true),
		endpoint.Response(http.StatusOK, Pet{}, "successful operation"),
		endpoint.Security("petstore_auth", "read:pets"),
	)

	api := swag.New(
		swag.Endpoints(post, get),
		swag.Security("petstore_auth", "read:pets"),
		swag.SecurityScheme("petstore_auth",
			swagger.OAuth2Security("accessCode", "http://example.com/oauth/authorize", "http://example.com/oauth/token"),
			swagger.OAuth2Scope("write:pets", "modify pets in your account"),
			swagger.OAuth2Scope("read:pets", "read your pets"),
		),
	)

	for path, endpoints := range api.Paths {
		http.Handle(path, endpoints)
	}

	enableCors := true
	http.Handle("/swagger", api.Handler(enableCors))

	http.ListenAndServe(":8080", nil)
}
