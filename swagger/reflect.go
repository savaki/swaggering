// Package swagger ...
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
package swagger

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func inspect(t reflect.Type, tag reflect.StructTag) Property {
	if p, ok := customTypes[t]; ok {
		return p
	}

	jsonTag := tag.Get("json")
	defaultTag := tag.Get("default")
	formatTag := tag.Get("format")
	minLenTag := tag.Get("min_length")
	maxLenTag := tag.Get("max_length")
	minimumTag := tag.Get("minimum")
	exclusiveMinimumTag := tag.Get("exclusive_minimum")
	exclusiveMaximumTag := tag.Get("exclusive_maximum")
	maximumTag := tag.Get("maximum")
	patternTag := tag.Get("pattern")
	enumTag := tag.Get("enum")

	if t.Kind() == reflect.Ptr {
		if p, ok := customTypes[t.Elem()]; ok {
			p.Nullable = true
			return p
		}
	}

	p := Property{
		GoType: t,
	}

	if strings.Contains(jsonTag, ",string") {
		p.Type = "string"
		return p
	}

	var err error
	switch p.GoType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		p.Type = "integer"
		p.Format = "int32"
		if defaultTag != "" {
			p.Default, err = strconv.ParseInt(defaultTag, 10, 32)
			if err != nil {
				panic(fmt.Errorf("Failed to convert default tag value: %s", err))
			}
		}
		if minimumTag != "" {
			var min int64
			min, err = strconv.ParseInt(minimumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert minimum tag value: %s", err))
			}
			p.Minimum = &min
		}
		if maximumTag != "" {
			var max int64
			max, err = strconv.ParseInt(maximumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert maximum tag value: %s", err))
			}
			p.Maximum = &max
		}
		if exclusiveMinimumTag != "" {
			p.ExclusiveMinimum, err = strconv.ParseBool(exclusiveMinimumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_minimum tag value: %s", err))
			}
		}
		if exclusiveMaximumTag != "" {
			p.ExclusiveMaximum, err = strconv.ParseBool(exclusiveMaximumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_maximum tag value: %s", err))
			}
		}

	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		p.Type = "integer"
		p.Format = "int32"
		if defaultTag != "" {
			p.Default, err = strconv.ParseUint(defaultTag, 10, 32)
			if err != nil {
				panic(fmt.Errorf("Failed to convert default tag value: %s", err))
			}
		}
		if minimumTag != "" {
			var min int64
			min, err = strconv.ParseInt(minimumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert minimum tag value: %s", err))
			}
			p.Minimum = &min
		}
		if maximumTag != "" {
			var max int64
			max, err = strconv.ParseInt(maximumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert maximum tag value: %s", err))
			}
			p.Maximum = &max
		}
		if exclusiveMinimumTag != "" {
			p.ExclusiveMinimum, err = strconv.ParseBool(exclusiveMinimumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_minimum tag value: %s", err))
			}
		}
		if exclusiveMaximumTag != "" {
			p.ExclusiveMaximum, err = strconv.ParseBool(exclusiveMaximumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_maximum tag value: %s", err))
			}
		}

	case reflect.Int64:
		p.Type = "integer"
		p.Format = "int64"
		if defaultTag != "" {
			p.Default, err = strconv.ParseInt(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert default tag value: %s", err))
			}
		}
		if minimumTag != "" {
			var min int64
			min, err = strconv.ParseInt(minimumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert minimum tag value: %s", err))
			}
			p.Minimum = &min
		}
		if maximumTag != "" {
			var max int64
			max, err = strconv.ParseInt(maximumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert maximum tag value: %s", err))
			}
			p.Maximum = &max
		}
		if exclusiveMinimumTag != "" {
			p.ExclusiveMinimum, err = strconv.ParseBool(exclusiveMinimumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_minimum tag value: %s", err))
			}
		}
		if exclusiveMaximumTag != "" {
			p.ExclusiveMaximum, err = strconv.ParseBool(exclusiveMaximumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_maximum tag value: %s", err))
			}
		}

	case reflect.Uint64:
		p.Type = "integer"
		p.Format = "int64"
		if defaultTag != "" {
			p.Default, err = strconv.ParseUint(defaultTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert default tag value: %s", err))
			}
		}
		if minimumTag != "" {
			var min int64
			min, err = strconv.ParseInt(minimumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert minimum tag value: %s", err))
			}
			p.Minimum = &min
		}
		if maximumTag != "" {
			var max int64
			max, err = strconv.ParseInt(maximumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert maximum tag value: %s", err))
			}
			p.Maximum = &max
		}
		if exclusiveMinimumTag != "" {
			p.ExclusiveMinimum, err = strconv.ParseBool(exclusiveMinimumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_minimum tag value: %s", err))
			}
		}
		if exclusiveMaximumTag != "" {
			p.ExclusiveMaximum, err = strconv.ParseBool(exclusiveMaximumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_maximum tag value: %s", err))
			}
		}

	case reflect.Float64:
		p.Type = "number"
		p.Format = "double"
		if defaultTag != "" {
			p.Default, err = strconv.ParseFloat(defaultTag, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert default tag value: %s", err))
			}
		}
		if minimumTag != "" {
			var min int64
			min, err = strconv.ParseInt(minimumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert minimum tag value: %s", err))
			}
			p.Minimum = &min
		}
		if maximumTag != "" {
			var max int64
			max, err = strconv.ParseInt(maximumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert maximum tag value: %s", err))
			}
			p.Maximum = &max
		}
		if exclusiveMinimumTag != "" {
			p.ExclusiveMinimum, err = strconv.ParseBool(exclusiveMinimumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_minimum tag value: %s", err))
			}
		}
		if exclusiveMaximumTag != "" {
			p.ExclusiveMaximum, err = strconv.ParseBool(exclusiveMaximumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_maximum tag value: %s", err))
			}
		}

	case reflect.Float32:
		p.Type = "number"
		p.Format = "float"
		if defaultTag != "" {
			p.Default, err = strconv.ParseFloat(defaultTag, 32)
			if err != nil {
				panic(fmt.Errorf("Failed to convert default tag value: %s", err))
			}
		}
		if minimumTag != "" {
			var min int64
			min, err = strconv.ParseInt(minimumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert minimum tag value: %s", err))
			}
			p.Minimum = &min
		}
		if maximumTag != "" {
			var max int64
			max, err = strconv.ParseInt(maximumTag, 10, 64)
			if err != nil {
				panic(fmt.Errorf("Failed to convert maximum tag value: %s", err))
			}
			p.Maximum = &max
		}
		if exclusiveMinimumTag != "" {
			p.ExclusiveMinimum, err = strconv.ParseBool(exclusiveMinimumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_minimum tag value: %s", err))
			}
		}
		if exclusiveMaximumTag != "" {
			p.ExclusiveMaximum, err = strconv.ParseBool(exclusiveMaximumTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert exclusive_maximum tag value: %s", err))
			}
		}

	case reflect.Bool:
		p.Type = "boolean"
		if defaultTag != "" {
			p.Default, err = strconv.ParseBool(defaultTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert default tag value: %s", err))
			}
		}

	case reflect.String:
		p.Type = "string"
		if defaultTag != "" {
			p.Default = defaultTag
		}
		if formatTag != "" {
			splits := strings.Split(formatTag, ",")
			if splits[0] == "enum" {
				for _, eVal := range splits[1:] {
					p.Enum = append(p.Enum, strings.TrimSpace(eVal))
				}
			} else {
				p.Format = strings.TrimSpace(splits[0])
			}
		}

		if enumTag != "" {
			splits := strings.Split(enumTag, ",")
			for _, eVal := range splits {
				p.Enum = append(p.Enum, strings.TrimSpace(eVal))
			}
		}

		if minLenTag != "" {
			p.MinLength, err = strconv.Atoi(minLenTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert min tag value: %s", err))
			}
		}

		if maxLenTag != "" {
			p.MaxLength, err = strconv.Atoi(maxLenTag)
			if err != nil {
				panic(fmt.Errorf("Failed to convert max tag value: %s", err))
			}
		}

		if patternTag != "" {
			_, err := regexp.Compile(patternTag)
			if err != nil {
				panic(fmt.Errorf("Failed to compile regexp: %s", err))
			}

			p.Pattern = patternTag
		}

	case reflect.Struct:
		name := makeName(p.GoType)
		p.Ref = makeRef(name)

	case reflect.Ptr:
		p := inspect(t.Elem(), tag)
		p.Nullable = true
		return p

	case reflect.Map:
		p.Type = "object"

	case reflect.Slice:
		// For json.RawMessage
		if p.GoType.PkgPath() == "encoding/json" && p.GoType.Name() == "RawMessage" {
			p.Type = "object"
			return p
		}

		p.Type = "array"
		p.Items = &Items{}

		p.GoType = t.Elem() // dereference the slice
		switch p.GoType.Kind() {
		case reflect.Ptr:
			p.GoType = p.GoType.Elem()
			name := makeName(p.GoType)
			p.Items.Ref = makeRef(name)

		case reflect.Struct:
			name := makeName(p.GoType)
			p.Items.Ref = makeRef(name)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32:
			p.Items.Type = "integer"
			p.Items.Format = "int32"

		case reflect.Int64, reflect.Uint64:
			p.Items.Type = "integer"
			p.Items.Format = "int64"

		case reflect.Float64:
			p.Items.Type = "number"
			p.Items.Format = "double"

		case reflect.Float32:
			p.Items.Type = "number"
			p.Items.Format = "float"

		case reflect.String:
			p.Items.Type = "string"
			if formatTag != "" {
				splits := strings.Split(formatTag, ",")
				if splits[0] == "enum" {
					for _, eVal := range splits[1:] {
						p.Items.Enum = append(p.Items.Enum, strings.TrimSpace(eVal))
					}
				} else {
					p.Items.Format = strings.TrimSpace(splits[0])
				}
			}

			if enumTag != "" {
				splits := strings.Split(enumTag, ",")
				for _, eVal := range splits {
					p.Items.Enum = append(p.Enum, strings.TrimSpace(eVal))
				}
			}

			if minLenTag != "" {
				p.Items.MinLength, err = strconv.Atoi(minLenTag)
				if err != nil {
					panic(fmt.Errorf("Failed to convert min tag value: %s", err))
				}
			}

			if maxLenTag != "" {
				p.Items.MaxLength, err = strconv.Atoi(maxLenTag)
				if err != nil {
					panic(fmt.Errorf("Failed to convert max tag value: %s", err))
				}
			}

			if patternTag != "" {
				_, err := regexp.Compile(patternTag)
				if err != nil {
					panic(fmt.Errorf("Failed to compile regexp: %s", err))
				}

				p.Items.Pattern = patternTag
			}
		}
	}

	return p
}

func defineObject(v interface{}) Object {
	var required []string

	var t reflect.Type
	switch value := v.(type) {
	case reflect.Type:
		t = value
	default:
		t = reflect.TypeOf(v)
	}

	properties := map[string]Property{}
	isArray := t.Kind() == reflect.Slice

	if isArray {
		t = t.Elem()
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		p := inspect(t, "")
		return Object{
			IsArray:  isArray,
			GoType:   t,
			Type:     p.Type,
			Format:   p.Format,
			Name:     t.Kind().String(),
			Required: required,
		}
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// skip unexported fields
		if strings.ToLower(field.Name[0:1]) == field.Name[0:1] {
			continue
		}

		// determine the json name of the field
		name := strings.TrimSpace(field.Tag.Get("json"))
		if name == "" || strings.HasPrefix(name, ",") {
			name = field.Name

		} else {
			// strip out things like , omitempty
			parts := strings.Split(name, ",")
			name = parts[0]
		}

		parts := strings.Split(name, ",") // foo,omitempty => foo
		name = parts[0]
		if name == "-" {
			// honor json ignore tag
			continue
		}

		// determine if this field is required or not
		if v := field.Tag.Get("required"); v == "true" {
			if required == nil {
				required = []string{}
			}
			required = append(required, name)
		}

		// support go-playground/validator binding tags
		if v := field.Tag.Get("binding"); v != "" {
			parts := strings.Split(v, ",") // "gt=0,dive,len=1,dive,required"
			for _, a := range parts {
				if a == "required" {
					if required == nil {
						required = []string{}
					}
					required = append(required, name)
				}
			}
		}

		p := inspect(field.Type, field.Tag)

		properties[name] = p
	}

	return Object{
		IsArray:    isArray,
		GoType:     t,
		Type:       "object",
		Name:       makeName(t),
		Required:   required,
		Properties: properties,
	}
}

func define(v interface{}) map[string]Object {
	objMap := map[string]Object{}

	obj := defineObject(v)
	objMap[obj.Name] = obj

	dirty := true

	for dirty {
		dirty = false
		for _, d := range objMap {
			for _, p := range d.Properties {
				if _, ok := customTypes[p.GoType]; ok {
					continue
				}
				if p.GoType.Kind() == reflect.Struct {
					name := makeName(p.GoType)
					if _, exists := objMap[name]; !exists {
						child := defineObject(p.GoType)
						objMap[child.Name] = child
						dirty = true
					}
				}
			}
		}
	}

	return objMap
}

// MakeSchema takes struct or pointer to a struct and returns a Schema instance suitable for use by the swagger doc
func MakeSchema(prototype interface{}) *Schema {
	schema := &Schema{
		Prototype: prototype,
	}

	obj := defineObject(prototype)
	if obj.IsArray {
		schema.Type = "array"
		schema.Items = &Items{
			Ref: makeRef(obj.Name),
		}

	} else {
		schema.Ref = makeRef(obj.Name)
	}

	return schema
}
