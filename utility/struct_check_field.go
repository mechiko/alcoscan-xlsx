package utility

import (
	"reflect"
	"strings"
)

func StructHasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

func FindStringInJsonTags(s interface{}, target string) (bool, string) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return false, ""
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == "" {
			// No JSON tag, skip
			continue
		}

		// Extract the field name from the JSON tag (before comma, if any)
		jsonFieldName := strings.SplitN(jsonTag, ",", 2)[0]

		if jsonFieldName == target {
			return true, field.Name // Found the target string in a JSON tag
		}
	}
	return false, "" // Target string not found in any JSON tag
}
