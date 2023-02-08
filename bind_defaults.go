package defaultor

import (
	"reflect"
	"strconv"
)

const defaultTagName = "default"

func BindDefault(s interface{}) {
	value := reflect.ValueOf(s).Elem()
	typ := value.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		defaultValue := field.Tag.Get(defaultTagName)
		if defaultValue != "" {
			fieldValue := value.Field(i)
			if fieldValue.IsValid() && fieldValue.CanSet() {
				switch field.Type.Kind() {
				case reflect.String:
					fieldValue.SetString(defaultValue)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					intValue, err := strconv.ParseInt(defaultValue, 10, 64)
					if err == nil {
						fieldValue.SetInt(intValue)
					}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					uintValue, err := strconv.ParseUint(defaultValue, 10, 64)
					if err == nil {
						fieldValue.SetUint(uintValue)
					}
				case reflect.Float32, reflect.Float64:
					floatValue, err := strconv.ParseFloat(defaultValue, 64)
					if err == nil {
						fieldValue.SetFloat(floatValue)
					}
				case reflect.Bool:
					boolValue, err := strconv.ParseBool(defaultValue)
					if err == nil {
						fieldValue.SetBool(boolValue)
					}
				}
			}
		}
	}
}
