package restserver

import "reflect"

// Return list of all fields in a struct, recursively by embedded structs.
func StructFields(s any) []reflect.StructField {
	var fields []reflect.StructField
	structFieldsRec(s, &fields)
	return fields
}

func structFieldsRec(s any, fields *[]reflect.StructField) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			structFieldsRec(v.Field(i).Interface(), fields)
		} else {
			*fields = append(*fields, field)
		}
	}
}

func IsNil(i any) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
