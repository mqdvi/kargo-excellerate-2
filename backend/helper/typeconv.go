package helper

import (
	"fmt"
	"reflect"
	"strconv"
)

// BoolToString convert boolean to string
func BoolToString(value bool) string {
	return strconv.FormatBool(value)
}

// StringToBool convert string to bool
func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

// Uint64ToString convert uint64 to string
func Uint64ToString(value uint64) string {
	return strconv.FormatUint(value, 10)
}

// IntToString convert integer to string
func IntToString(value int) string {
	return strconv.Itoa(value)
}

// StringToInt convert string to integer
func StringToInt(value string) int {
	out, err := strconv.Atoi(value)
	if err != nil {
		out = 0
	}

	return out
}

// StringToUint64 convert string to uint64
func StringToUint64(value string) uint64 {
	out := StringToInt(value)

	return uint64(out)
}

// ReflectValueToString convert refect.Value to String
func ReflectValueToString(value reflect.Value) string {
	stringValue := ""

	switch value.Kind() {
	case reflect.String:
		stringValue = value.Interface().(string)
	case reflect.Bool:
		stringValue = BoolToString(value.Interface().(bool))
	case reflect.Uint64:
		stringValue = Uint64ToString(value.Interface().(uint64))
	default:
		stringValue = fmt.Sprint(value)
	}

	return stringValue
}

// ReflectInterfaceToMapStringInterface convert struct to map[string]interface
func ReflectInterfaceToMapStringInterface(data interface{}) map[string]interface{} {
	v := reflect.ValueOf(data)
	res := make(map[string]interface{})
	if data == nil {
		return nil
	}

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			res[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
	case reflect.Ptr: // Not support pointer for now
		return nil
	default:
		res["value"] = v.Interface()
	}

	return res
}
