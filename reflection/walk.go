/*
Module to exercise the use of reflection across multiple object types.

Provides examples of the following:
 Reflection
 Recursion
 testify/assert

"Now that you know about reflection, do your best to avoid using it."
*/
package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			Walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			Walk(v.Interface(), fn)
		}
	case reflect.Func:
		result := val.Call(nil)

		for _, res := range result {
			Walk(res.Interface(), fn)
		}
	}
}
