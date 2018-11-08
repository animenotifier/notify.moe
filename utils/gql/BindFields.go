package gql

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
)

// TAG is the tag used for json
const TAG = "json"

// BindFields can't take recursive slice type
// e.g
// type Person struct{
//	Friends []Person
// }
// it will throw panic stack-overflow
func BindFields(t reflect.Type) graphql.Fields {
	fields := make(map[string]*graphql.Field)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip private fields
		if field.Tag.Get("private") == "true" {
			continue
		}

		tag := extractTag(field.Tag)

		if tag == "-" {
			continue
		}

		fieldType := field.Type

		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		var graphType graphql.Output

		if fieldType.Kind() == reflect.Struct {
			structFields := BindFields(t.Field(i).Type)

			if tag == "" {
				fields = appendFields(fields, structFields)
				continue
			} else {
				graphType = graphql.NewObject(graphql.ObjectConfig{
					Name:   t.Name() + "_" + field.Name,
					Fields: structFields,
				})
			}
		}

		if tag == "" {
			continue
		}

		if graphType == nil {
			graphType = getGraphType(fieldType)
		}

		fields[tag] = &graphql.Field{
			Type: graphType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return extractValue(tag, p.Source), nil
			},
		}
	}
	return fields
}

func getGraphType(tipe reflect.Type) graphql.Output {
	kind := tipe.Kind()

	switch kind {
	case reflect.String:
		return graphql.String
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return graphql.Int
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		return graphql.Float
	case reflect.Bool:
		return graphql.Boolean
	case reflect.Slice:
		return getGraphList(tipe)
	}

	return graphql.String
}

func getGraphList(tipe reflect.Type) *graphql.List {
	if tipe.Kind() == reflect.Slice {
		switch tipe.Elem().Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			return graphql.NewList(graphql.Int)
		case reflect.Bool:
			return graphql.NewList(graphql.Boolean)
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			return graphql.NewList(graphql.Float)
		case reflect.String:
			return graphql.NewList(graphql.String)
		}
	}

	// finally bind object
	t := reflect.New(tipe.Elem())
	name := strings.Replace(fmt.Sprint(tipe.Elem()), ".", "_", -1)
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: BindFields(t.Elem().Type()),
	})

	return graphql.NewList(obj)
}

func appendFields(dest, origin graphql.Fields) graphql.Fields {
	for key, value := range origin {
		dest[key] = value
	}

	return dest
}

func extractValue(originTag string, obj interface{}) interface{} {
	val := reflect.Indirect(reflect.ValueOf(obj))

	for j := 0; j < val.NumField(); j++ {
		field := val.Type().Field(j)

		if field.Type.Kind() == reflect.Struct {
			res := extractValue(originTag, val.Field(j).Interface())
			if res != nil {
				return res
			}
		}

		if originTag == extractTag(field.Tag) {
			return reflect.Indirect(val.Field(j)).Interface()
		}
	}

	return nil
}

func extractTag(tag reflect.StructTag) string {
	t := tag.Get(TAG)

	if t != "" {
		t = strings.Split(t, ",")[0]
	}

	return t
}

// // BindArg is a lazy way of binding args
// func BindArg(obj interface{}, tags ...string) graphql.FieldConfigArgument {
// 	v := reflect.Indirect(reflect.ValueOf(obj))
// 	var config = make(graphql.FieldConfigArgument)

// 	for i := 0; i < v.NumField(); i++ {
// 		field := v.Type().Field(i)
// 		mytag := extractTag(field.Tag)

// 		if inArray(tags, mytag) {
// 			config[mytag] = &graphql.ArgumentConfig{
// 				Type: getGraphType(field.Type),
// 			}
// 		}
// 	}

// 	return config
// }

// func inArray(slice interface{}, item interface{}) bool {
// 	s := reflect.ValueOf(slice)

// 	if s.Kind() != reflect.Slice {
// 		panic("inArray() given a non-slice type")
// 	}

// 	for i := 0; i < s.Len(); i++ {
// 		if reflect.DeepEqual(item, s.Index(i).Interface()) {
// 			return true
// 		}
// 	}

// 	return false
// }
