package replacement

import "reflect"

// Replace replace struct value by another struct
func Replace(old, new interface{}) {
	if reflect.ValueOf(old).Kind() != reflect.Ptr {
		panic("old must be a pointer to a struct")
	}
	if reflect.ValueOf(new).Kind() != reflect.Ptr {
		panic("new must be a pointer to a struct")
	}
	if reflect.ValueOf(old).Elem().Kind() != reflect.Struct {
		panic("old must be a pointer to a struct")
	}
	if reflect.ValueOf(new).Elem().Kind() != reflect.Struct {
		panic("new must be a pointer to a struct")
	}
	oldReflect := reflect.ValueOf(old).Elem()
	newReflect := reflect.ValueOf(new).Elem()
	for i, n := 0, oldReflect.NumField(); i < n; i++ {
		if !newReflect.Field(i).IsNil() {
			oldReflect.Field(i).Set(newReflect.Field(i).Elem())
		}
	}
	return
}
