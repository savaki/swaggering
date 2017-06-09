package endpoint_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/MarkSonghurst/swag"
	"github.com/MarkSonghurst/swag/endpoint"
	"github.com/MarkSonghurst/swag/swagger"
	"github.com/stretchr/testify/assert"
)

func Echo(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "hello world")
}

func TestNew(t *testing.T) {
	summary := "here's the summary"
	e := endpoint.New("get", "/", summary,
		endpoint.Handler(Echo),
	)

	assert.Equal(t, "GET", e.Method)
	assert.Equal(t, "/", e.Path)
	assert.NotNil(t, e.Handler)
	assert.Equal(t, []string{"application/json"}, e.Consumes)
	assert.Equal(t, []string{"application/json"}, e.Produces)
	assert.Equal(t, summary, e.Summary)
	assert.Equal(t, []string{}, e.Tags)
}

func TestTags(t *testing.T) {
	e := endpoint.New("get", "/", "get thing",
		endpoint.Tags("blah"),
	)

	assert.Equal(t, []string{"blah"}, e.Tags)
}

func TestDescription(t *testing.T) {
	e := endpoint.New("get", "/", "get thing",
		endpoint.Description("blah"),
	)

	assert.Equal(t, "blah", e.Description)
}

func TestOperationId(t *testing.T) {
	e := endpoint.New("get", "/", "get thing",
		endpoint.OperationID("blah"),
	)

	assert.Equal(t, "blah", e.OperationID)
}

func TestProduces(t *testing.T) {
	expected := []string{"a", "b"}
	e := endpoint.New("get", "/", "get thing",
		endpoint.Produces(expected...),
	)

	assert.Equal(t, expected, e.Produces)
}

func TestConsumes(t *testing.T) {
	expected := []string{"a", "b"}
	e := endpoint.New("get", "/", "get thing",
		endpoint.Consumes(expected...),
	)

	assert.Equal(t, expected, e.Consumes)
}

func TestPath(t *testing.T) {
	expected := swagger.Parameter{
		In:          "path",
		Name:        "id",
		Description: "the description",
		Required:    true,
		Items: swagger.Items{
			Type: "string",
		},
	}

	e := endpoint.New("get", "/", "get thing",
		endpoint.Path(expected.Name, expected.Items, nil, expected.Description, expected.Required),
	)

	assert.Equal(t, 1, len(e.Parameters))
	assert.Equal(t, expected, e.Parameters[0])
}

func TestQuery(t *testing.T) {
	expected := swagger.Parameter{
		In:          "query",
		Name:        "id",
		Description: "the description",
		Required:    true,
		Items: swagger.Items{
			Type: "string",
		},
	}

	e := endpoint.New("get", "/", "get thing",
		endpoint.Query(expected.Name, expected.Items, nil, expected.Description, expected.Required),
	)

	assert.Equal(t, 1, len(e.Parameters))
	assert.Equal(t, expected, e.Parameters[0])
}

func TestHeader(t *testing.T) {
	expected := swagger.Parameter{
		In:          "header",
		Name:        "id",
		Description: "the description",
		Required:    true,
		Items: swagger.Items{
			Type: "string",
		},
	}

	e := endpoint.New("get", "/", "get thing",
		endpoint.Header(expected.Name, expected.Items, nil, expected.Description, expected.Required),
	)

	assert.Equal(t, 1, len(e.Parameters))
	assert.Equal(t, expected, e.Parameters[0])
}

func TestFormData(t *testing.T) {
	expected := swagger.Parameter{
		In:          "formData",
		Name:        "id",
		Description: "the description",
		Required:    true,
		Items: swagger.Items{
			Type: "string",
		},
	}

	e := endpoint.New("get", "/", "get thing",
		endpoint.FormData(expected.Name, expected.Items, nil, expected.Description, expected.Required),
	)

	assert.Equal(t, 1, len(e.Parameters))
	assert.Equal(t, expected, e.Parameters[0])
}

type Model struct {
	String string `json:"s"`
}

func TestBody(t *testing.T) {
	expected := swagger.Parameter{
		In:          "body",
		Name:        "body",
		Description: "the description",
		Required:    true,
		Schema: &swagger.Schema{
			Ref:       "#/definitions/endpoint_testModel",
			Prototype: Model{},
		},
	}

	e := endpoint.New("get", "/", "get thing",
		endpoint.Body(Model{}, expected.Description, expected.Required),
	)

	assert.Equal(t, 1, len(e.Parameters))
	assert.Equal(t, expected, e.Parameters[0])
}

func TestResponse(t *testing.T) {
	expected := swagger.Response{
		Description: "successful",
		Schema: &swagger.Schema{
			Ref:       "#/definitions/endpoint_testModel",
			Prototype: Model{},
		},
	}

	e := endpoint.New("get", "/", "get thing",
		endpoint.Response(http.StatusOK, Model{}, "successful"),
	)

	assert.Equal(t, 1, len(e.Responses))
	assert.Equal(t, expected, e.Responses["200"])
}

func TestResponseHeader(t *testing.T) {
	expected := swagger.Response{
		Description: "successful",
		Schema: &swagger.Schema{
			Ref:       "#/definitions/endpoint_testModel",
			Prototype: Model{},
		},
		Headers: map[string]swagger.ResponseHeader{
			"X-Rate-Limit": {
				Type:        "integer",
				Format:      "int32",
				Description: "calls per hour allowed by the user",
			},
		},
	}

	e := endpoint.New("get", "/", "get thing",
		endpoint.Response(http.StatusOK, Model{}, "successful",
			endpoint.ResponseHeader("X-Rate-Limit", "integer", "int32", "calls per hour allowed by the user"),
		),
	)

	assert.Equal(t, 1, len(e.Responses))
	assert.Equal(t, expected, e.Responses["200"])
}

func TestResponseExample(t *testing.T) {
	example := Model{
		String: "alpha",
	}

	expected := swagger.Response{
		Description: "successful",
		Schema: &swagger.Schema{
			Ref:       "#/definitions/endpoint_testModel",
			Prototype: Model{},
		},
		Examples: swagger.ResponseExamples{
			"application/json": example,
		},
	}

	e := endpoint.New("get", "/", "get thing",
		endpoint.Response(http.StatusOK, Model{}, "successful",
			endpoint.ResponseExample("application/json", example),
		),
	)

	assert.Equal(t, 1, len(e.Responses))
	assert.Equal(t, expected, e.Responses["200"])
}

func TestSecurityScheme(t *testing.T) {
	api := swag.New(
		swag.SecurityScheme("basic", swagger.BasicSecurity()),
		swag.SecurityScheme("apikey", swagger.APIKeySecurity("Authorization", "header")),
	)
	assert.Len(t, api.SecurityDefinitions, 2)
	assert.Contains(t, api.SecurityDefinitions, "basic")
	assert.Contains(t, api.SecurityDefinitions, "apikey")
	assert.Equal(t, "header", api.SecurityDefinitions["apikey"].In)
}

func TestSecurity(t *testing.T) {
	e := endpoint.New("get", "/", "",
		endpoint.Handler(Echo),
		endpoint.Security("basic"),
		endpoint.Security("oauth2", "scope1", "scope2"),
	)
	assert.False(t, e.Security.DisableSecurity)
	assert.Len(t, e.Security.Requirements, 2)
	assert.Contains(t, e.Security.Requirements[0], "basic")
	assert.Contains(t, e.Security.Requirements[1], "oauth2")
	assert.Len(t, e.Security.Requirements[1]["oauth2"], 2)
}

func TestNoSecurity(t *testing.T) {
	e := endpoint.New("get", "/", "",
		endpoint.Handler(Echo),
		endpoint.NoSecurity(),
	)
	assert.True(t, e.Security.DisableSecurity)
}
