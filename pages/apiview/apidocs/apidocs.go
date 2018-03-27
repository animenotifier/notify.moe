package apidocs

import (
	"reflect"
	"unicode"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// ByType renders the api docs page for the given type.
func ByType(typeName string) func(*aero.Context) string {
	return func(ctx *aero.Context) string {
		t := arn.API.Type(typeName)
		fields := []*utils.APIField{}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			if field.Anonymous || !unicode.IsUpper(rune(field.Name[0])) {
				continue
			}

			typeName := ""

			switch field.Type.Kind() {
			case reflect.Ptr:
				typeName = field.Type.Elem().Name()

			case reflect.Slice:
				sliceElementType := field.Type.Elem()

				if sliceElementType.Kind() == reflect.Ptr {
					sliceElementType = sliceElementType.Elem()
				}

				typeName = sliceElementType.Name() + "[]"

			default:
				typeName = field.Type.Name()
			}

			fields = append(fields, &utils.APIField{
				Name: field.Name,
				JSON: field.Tag.Get("json"),
				Type: typeName,
			})
		}

		return ctx.HTML(components.APIDocs(t, fields))
	}
}
