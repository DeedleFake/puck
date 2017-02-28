package util

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
func ExpandStruct(i interface{}, tag string) interface{} {
	v := reflect.Indirect(reflect.ValueOf(i))
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("Expected a struct, got a %v", v.Kind()))
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

	var stack []string
	var expander func(string) string
	expander = func(name string) string {
		for _, prev := range stack {
			if prev == name {
				panic(fmt.Errorf("Recursive loop detected while expanding %q", name))
			}
		}

		stack = append(stack, name)
		defer func() {
			stack = stack[:len(stack)-1]
		}()

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
