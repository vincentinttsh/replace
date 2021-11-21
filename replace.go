package replace

import (
	"reflect"
)

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
	oldField := reflect.TypeOf(old).Elem()
	newField := reflect.TypeOf(new).Elem()
	newReflect := reflect.ValueOf(new).Elem()
	for i, n := 0, newField.NumField(); i < n; i++ {
		name := newField.Field(i).Name
		_, oldHave := oldField.FieldByName(name)
		if oldHave && !newReflect.FieldByName(name).IsNil() {
			oldReflect.FieldByName(name).Set(newReflect.FieldByName(name).Elem())
		}
	}
}
