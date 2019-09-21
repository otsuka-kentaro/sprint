package sprint

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	nilString = "<nil>"
)

// Sprint returns string like fmt.Sprint.
// Sprint dereferences pointer values.
func Sprint(s interface{}) string {
	return sprintValue(reflect.ValueOf(s), true)
}

// sprintValue handles value for iterating use.
func sprintValue(v reflect.Value, putsNilString bool) string {
	switch v.Kind() {
	case reflect.Invalid:
		if putsNilString {
			return nilString
		}

		return ""
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String, reflect.Map,
		reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return fmt.Sprint(v)
	case reflect.Array, reflect.Slice:
		var elements []string
		for i := 0; i < v.Len(); i++ {
			print := sprintValue(v.Index(i), false)

			// if empty element
			if print == "" {
				continue
			}
			elements = append(elements, print)
		}
		// returns empty string if all field are invalid
		if len(elements) == 0 {
			return ""
		}
		return "[" + strings.Join(elements, ", ") + "]"
	case reflect.Struct:
		var fields []string
		for i := 0; i < v.NumField(); i++ {
			print := sprintValue(v.Field(i), false)

			// if empty field
			if print == "" {
				continue
			}
			fields = append(fields, v.Type().Field(i).Name+": "+print)
		}
		// returns empty string if all field are invalid
		if len(fields) == 0 {
			return ""
		}
		return "{" + strings.Join(fields, ", ") + "}"
	case reflect.Interface:
		el := v.Elem()
		if !el.IsValid() {
			if putsNilString {
				return nilString
			}

			return ""
		}

		return sprintValue(v, putsNilString)
	case reflect.Ptr:
		return sprintValue(v.Elem(), putsNilString)
	default:
		return fmt.Sprint(v)
	}
}
