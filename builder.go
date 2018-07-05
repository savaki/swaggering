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
package swag

import "github.com/miketonks/swag/swagger"

// Builder uses the builder pattern to generate a swagger definition
type Builder struct {
	API *swagger.API
}

// Option provides configuration options to the swagger api builder
type Option func(builder *Builder)

// Description sets info.description
func Description(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.Description = v
	}
}

// Version sets info.version
func Version(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.Version = v
	}
}

// TermsOfService sets info.termsOfService
func TermsOfService(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.TermsOfService = v
	}
}

// Title sets info.title
func Title(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.Title = v
	}
}

// ContactEmail sets info.contact.email
func ContactEmail(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.Contact = &swagger.Contact{Email: v}
	}
}

// License sets both info.license.name and info.license.url
func License(name, url string) Option {
	return func(builder *Builder) {
		builder.API.Info.License = &swagger.License{
			Name: name,
			URL:  url,
		}
	}
}

// BasePath sets basePath
func BasePath(v string) Option {
	return func(builder *Builder) {
		builder.API.BasePath = v
	}
}

// Schemes sets the scheme
func Schemes(v ...string) Option {
	return func(builder *Builder) {
		builder.API.Schemes = v
	}
}

// TagOption provides additional customizations to the #Tag option
type TagOption func(tag *swagger.Tag)

// TagDescription sets externalDocs.description on the tag field
func TagDescription(v string) TagOption {
	return func(t *swagger.Tag) {
		t.Docs.Description = v
	}
}

// TagURL sets externalDocs.url on the tag field
func TagURL(v string) TagOption {
	return func(t *swagger.Tag) {
		t.Docs.URL = v
	}
}

// Tag adds a tag to the swagger api
func Tag(name, description string, options ...TagOption) Option {
	return func(builder *Builder) {
		if builder.API.Tags == nil {
			builder.API.Tags = []swagger.Tag{}
		}

		t := swagger.Tag{
			Name:        name,
			Description: description,
		}

		for _, opt := range options {
			opt(&t)
		}

		builder.API.Tags = append(builder.API.Tags, t)
	}
}

// Host specifies the host field
func Host(v string) Option {
	return func(builder *Builder) {
		builder.API.Host = v
	}
}

// Endpoints allows the endpoints to be added dynamically to the Api
func Endpoints(endpoints ...*swagger.Endpoint) Option {
	return func(builder *Builder) {
		for _, e := range endpoints {
			builder.API.AddEndpoint(e)
		}
	}
}

// SecurityScheme creates a new security definition for the API.
func SecurityScheme(name string, options ...swagger.SecuritySchemeOption) Option {
	scheme := swagger.SecurityScheme{}

	for _, opt := range options {
		opt(&scheme)
	}

	return SecurityDefinition(name, scheme)
}

// GoogleSecurityScheme creates a new security definition for the API.
func GoogleSecurityScheme(name string, options ...swagger.GoogleSecuritySchemeOption) Option {
	scheme := swagger.GoogleSecurityScheme{}

	for _, opt := range options {
		opt(&scheme)
	}

	return SecurityDefinition(name, scheme)
}

// SecurityDefinition creates a new security definition from a given security scheme for the API.
func SecurityDefinition(name string, scheme interface{}) Option {
	return func(builder *Builder) {
		if builder.API.SecurityDefinitions == nil {
			builder.API.SecurityDefinitions = make(map[string]interface{})
		}

		builder.API.SecurityDefinitions[name] = scheme
	}
}

// Security sets a default security scheme for all endpoints in the API.
func Security(scheme string, scopes ...string) Option {
	return func(b *Builder) {
		if scopes == nil {
			scopes = []string{}
		}
		if b.API.Security == nil {
			b.API.Security = &swagger.SecurityRequirement{}
		}

		if b.API.Security.Requirements == nil {
			b.API.Security.Requirements = []map[string][]string{}
		}

		b.API.Security.Requirements = append(b.API.Security.Requirements, map[string][]string{scheme: scopes})
	}
}

// New constructs a new api builder
func New(options ...Option) *swagger.API {
	b := &Builder{
		API: &swagger.API{
			BasePath: "/",
			Swagger:  "2.0",
			Info: swagger.Info{
				Description: "Describe your API",
				Title:       "Your API Title",
			},
		},
	}

	for _, opt := range options {
		opt(b)
	}

	return b.API
}
