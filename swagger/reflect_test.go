package swagger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	First string
}

type Pet struct {
	Friend      Person    `json:"friend"`
	Friends     []Person  `json:"friends"`
	Pointer     *Person   `json:"pointer" required:"true"`
	Pointers    []*Person `json:"pointers"`
	Int         int
	IntArray    []int
	String      string
	StringArray []string
}

type Empty struct {
	Nope int `json:"-"`
}

type APIResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Horse struct {
	Age int `json:"age" required:"true"`
	Breed
	Owner Person `json:"owner" required:"false"`
}

type Breed struct {
	Thoroughbred bool `json:"thoroughbred" required:"true"`
	*Origin
}

type Origin struct {
	Country string `json:"origin_country" required:"true"`
}

func TestDefine(t *testing.T) {
	v := define(Pet{})
	obj, ok := v["swaggerPet"]
	assert.True(t, ok)
	assert.False(t, obj.IsArray)
	assert.Equal(t, 8, len(obj.Properties))

	content := map[string]Object{}
	data, err := ioutil.ReadFile("testdata/pet.json")
	assert.Nil(t, err)
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&content)
	assert.Nil(t, err)
	expected := content["swaggerPet"]

	assert.Equal(t, expected.IsArray, obj.IsArray, "expected IsArray to match")
	assert.Equal(t, expected.Type, obj.Type, "expected Type to match")
	assert.Equal(t, expected.Required, obj.Required, "expected Required to match")
	assert.Equal(t, len(expected.Properties), len(obj.Properties), "expected same number of properties")

	for k, v := range obj.Properties {
		e := expected.Properties[k]
		assert.Equal(t, e.Type, v.Type, "expected %v.Type to match", k)
		assert.Equal(t, e.Description, v.Description, "expected %v.Required to match", k)
		assert.Equal(t, e.Enum, v.Enum, "expected %v.Required to match", k)
		assert.Equal(t, e.Format, v.Format, "expected %v.Required to match", k)
		assert.Equal(t, e.Ref, v.Ref, "expected %v.Required to match", k)
		assert.Equal(t, e.Example, v.Example, "expected %v.Required to match", k)
		assert.Equal(t, e.Items, v.Items, "expected %v.Required to match", k)
	}
}

func TestDefineAnonymous(t *testing.T) {
	v := define(Horse{})
	obj, ok := v["swaggerHorse"]
	assert.True(t, ok)
	assert.False(t, obj.IsArray)
	assert.Equal(t, 4, len(obj.Properties))

	content := map[string]Object{}
	data, err := ioutil.ReadFile("testdata/horse.json")
	assert.Nil(t, err)
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&content)
	assert.Nil(t, err)
	expected := content["swaggerHorse"]

	assert.Equal(t, expected.IsArray, obj.IsArray, "expected IsArray to match")
	assert.Equal(t, expected.Type, obj.Type, "expected Type to match")
	assert.Equal(t, expected.Required, obj.Required, "expected Required to match")
	assert.Equal(t, len(expected.Properties), len(obj.Properties), "expected same number of properties")

	for k, v := range obj.Properties {
		e := expected.Properties[k]
		assert.Equal(t, e.Type, v.Type, "expected %v.Type to match", k)
		assert.Equal(t, e.Description, v.Description, "expected %v.Required to match", k)
		assert.Equal(t, e.Enum, v.Enum, "expected %v.Required to match", k)
		assert.Equal(t, e.Format, v.Format, "expected %v.Required to match", k)
		assert.Equal(t, e.Ref, v.Ref, "expected %v.Required to match", k)
		assert.Equal(t, e.Example, v.Example, "expected %v.Required to match", k)
		assert.Equal(t, e.Items, v.Items, "expected %v.Required to match", k)
	}
}

func TestNotStructDefine(t *testing.T) {
	v := define(int32(1))
	obj, ok := v["int32"]
	assert.True(t, ok)
	assert.False(t, obj.IsArray)
	assert.Equal(t, "integer", obj.Type)
	assert.Equal(t, "int32", obj.Format)

	v = define(uint64(1))
	obj, ok = v["uint64"]
	assert.True(t, ok)
	assert.False(t, obj.IsArray)
	assert.Equal(t, "integer", obj.Type)
	assert.Equal(t, "int64", obj.Format)

	v = define("")
	obj, ok = v["string"]
	assert.True(t, ok)
	assert.False(t, obj.IsArray)
	assert.Equal(t, "string", obj.Type)
	assert.Equal(t, "", obj.Format)

	v = define(byte(1))
	obj, ok = v["uint8"]
	if !assert.True(t, ok) {
		fmt.Printf("%v", v)
	}
	assert.False(t, obj.IsArray)
	assert.Equal(t, "integer", obj.Type)
	assert.Equal(t, "int32", obj.Format)

	v = define([]byte{1, 2})
	obj, ok = v["uint8"]
	if !assert.True(t, ok) {
		fmt.Printf("%v", v)
	}
	assert.True(t, obj.IsArray)
	assert.Equal(t, "integer", obj.Type)
	assert.Equal(t, "int32", obj.Format)
}

func TestHonorJsonIgnore(t *testing.T) {
	v := define(Empty{})
	obj, ok := v["swaggerEmpty"]
	assert.True(t, ok)
	assert.False(t, obj.IsArray)
	assert.Equal(t, 0, len(obj.Properties), "expected zero exposed properties")
}
