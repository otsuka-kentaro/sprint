package sprint

import (
	"fmt"
	"reflect"
	"strings"
)

// Sprint returns string like fmt.Sprint.
// Sprint dereferences pointer values.
func Sprint(s interface{}) string {
	return sprintValue(reflect.ValueOf(s))
}

// sprintValue handles value for iterating use.
func sprintValue(v reflect.Value) string {
	// dereference if value is a reference
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// just print if value is not a struct
	if v.Kind() != reflect.Struct {
		return fmt.Sprint(v)
	}

	// print struct
	var fields []string
	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, v.Type().Field(i).Name+": "+sprintValue(v.Field(i)))
	}
	return "{" + strings.Join(fields, ", ") + "}"
}
