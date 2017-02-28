package main

import (
	"fmt"
	"os"
	"reflect"
)

// ExpandStruct uses os.Expand to recursively expand self-references
// in struct fields based on tag names. In other words, given the
// struct
//
//     type Example struct {
//         First  string `json:"first"`
//         Second string `json:"other"`
//     }
//
//     example := &Example{
//         First:  "This is an ${other}.",
//         Second: "example",
//     }
//
// then ExpandStruct(example, "json") will return the following:
//
//     &Example{
//         First: "This is an example.",
//         Second: "example",
//     }
//
// BUG: If fields reference each other in a loop, then the expander
// will get stuck, recursing infinitely.
func ExpandStruct(i interface{}, tag string) interface{} {
	v := reflect.Indirect(reflect.ValueOf(i))
	if v.Kind() != reflect.Struct {
		panic("Expected a struct, got a " + v.Kind().String())
	}

	first := -1
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.String {
			first = i
			break
		}
	}
	if first < 0 {
		return i
	}

	r := reflect.New(v.Type()).Elem()

	var expander func(string) string
	expander = func(name string) string {
		for i := 0; i < v.NumField(); i++ {
			t := v.Type().Field(i).Tag.Get(tag)
			if t != name {
				continue
			}

			if r.Field(i).Interface() != v.Field(i).Interface() {
				switch v.Field(i).Kind() {
				case reflect.String:
					r.Field(i).SetString(os.Expand(v.Field(i).String(), expander))
				default:
					r.Field(i).Set(v.Field(i))
				}
			}

			return fmt.Sprint(r.Field(i).Interface())
		}

		return "${" + name + "}"
	}

	r.Field(first).SetString(os.Expand(v.Field(first).String(), expander))
	return r.Addr().Interface()
}
