package swagger

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Time is used to store time without date
type Time struct {
	time.Time
}

// Date is used to store date without time
type Date struct {
	time.Time
}

// UUID is used to store UUID values
type UUID string

const dtLayout = "2006-01-02"
const tmLayout = "15:04:05"

var nilTime = (time.Time{}).UnixNano()

// UnmarshalJSON handler for Date type
func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return
	}
	d.Time, err = time.Parse(dtLayout, s)
	return
}

// MarshalJSON handler for Date type
func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(dtLayout))), nil
}

// UnmarshalJSON handler for Time type
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(tmLayout, s)
	return
}

// MarshalJSON handler for Time type
func (t *Time) MarshalJSON() ([]byte, error) {
	if t.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(tmLayout))), nil
}

var customTypes map[reflect.Type]Property

func init() {
	customTypes = map[reflect.Type]Property{}

	RegisterCustomType(time.Time{}, Property{
		Type:   "string",
		Format: "date-time",
	})

	RegisterCustomType(Date{}, Property{
		Type:    "string",
		Pattern: "^\\d\\d:\\d\\d:\\d\\d$",
	})

	RegisterCustomType(Time{}, Property{
		Type:    "string",
		Pattern: "^\\d\\d-\\d\\d-\\d\\d$",
	})

	RegisterCustomType(UUID(""), Property{
		Type:   "string",
		Format: "uuid",
	})
}

// RegisterCustomType maps a reflect.Type to a pre-defined Property. This can be
// used to handle types that implement json.Marshaler or other interfaces.
// For example, a property with a Go type of time.Time would be represented as
// an object when it should be a string.
//
//    RegisterCustomType(time.Time{}, Property{
//      Type: "string",
//      Format: "date-time",
//    })
//
// Pointers to registered types will resolve to the same Property value unless
// that pointer type has also been registered as a custom type.
//
// For example: registering time.Time will also apply to *time.Time, unless
// *time.Time has also been registered.
func RegisterCustomType(v interface{}, p Property) {
	t := reflect.TypeOf(v)
	p.GoType = t
	customTypes[t] = p
}
