package reflection

import (
	"reflect"
)

func main() {

}

func walk(x interface{}, fn func(input string)) {
	val := trueValue(x)

	walkHelp := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			walkHelp(val.Index(i))
		}
	case reflect.Struct:
		for i := range val.NumField() {
			walkHelp(val.Field(i))
		}
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if i, ok := val.Recv(); ok {
				walkHelp(i)
			} else {
				break
			}
		}
	case reflect.Func:
		funResult := val.Call(nil)
		for _, value := range funResult {
			walkHelp(value)
		}
	}

}

func trueValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
