package apitype

import (
	"reflect"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/routetests"
)

// ByType renders the api docs page for the given type.
func ByType(typeName string) func(aero.Context) error {
	return func(ctx aero.Context) error {
		t := arn.API.Type(typeName)
		fields := []*utils.APIField{}
		embedded := []string{}

		if t.Kind() == reflect.Struct {
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)

				if field.Anonymous {
					embedded = append(embedded, field.Name)
					continue
				}

				if len(field.PkgPath) > 0 {
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
		}

		route := "/api/" + strings.ToLower(typeName) + "/:id"
		examples := routetests.All()[route]

		return ctx.HTML(components.APIType(t, examples, fields, embedded))
	}
}
