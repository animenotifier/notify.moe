package soundtrack

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Edit track.
func Edit(ctx *aero.Context) string {
	id := ctx.Get("id")
	track, err := arn.GetSoundTrack(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     track.Title,
			"og:image":     track.MainAnime().Image.Large,
			"og:url":       "https://" + ctx.App.Config.Domain + track.Link(),
			"og:site_name": "notify.moe",
			"og:type":      "music.song",
		},
	}

	return ctx.HTML(EditForm(track, "Edit soundtrack"))
}

// EditForm ...
func EditForm(obj interface{}, title string) string {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()
	id := reflect.Indirect(v.FieldByName("ID"))
	lowerCaseTypeName := strings.ToLower(t.Name())

	var b bytes.Buffer
	b.WriteString(`<div class="widget-form">`)
	b.WriteString(`<div class="widget" data-api="/api/` + lowerCaseTypeName + `/` + id.String() + `">`)
	b.WriteString(`<h1>`)
	b.WriteString(title)
	b.WriteString(`</h1>`)

	RenderObject(&b, obj, "")

	b.WriteString("</div>")
	b.WriteString("</div>")

	return b.String()
}

// RenderObject ...
func RenderObject(b *bytes.Buffer, obj interface{}, idPrefix string) {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	// Fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		RenderField(b, &v, field, idPrefix)
	}
}

// RenderField ...
func RenderField(b *bytes.Buffer, v *reflect.Value, field reflect.StructField, idPrefix string) {
	if field.Anonymous || field.Tag.Get("editable") != "true" {
		return
	}

	fieldValue := reflect.Indirect(v.FieldByName(field.Name))

	switch field.Type.String() {
	case "string":
		b.WriteString(components.InputText(idPrefix+field.Name, fieldValue.String(), field.Name, ""))
	case "[]string":
		b.WriteString(components.InputTags(idPrefix+field.Name, fieldValue.Interface().([]string), field.Name))
	case "[]*arn.ExternalMedia":
		for sliceIndex := 0; sliceIndex < fieldValue.Len(); sliceIndex++ {
			arrayObj := fieldValue.Index(sliceIndex).Interface()
			arrayIDPrefix := fmt.Sprintf("%s[%d].", field.Name, sliceIndex)
			RenderObject(b, arrayObj, arrayIDPrefix)
		}

		b.WriteString(`<div class="buttons"><button>` + utils.Icon("plus") + `Add ` + field.Name + `</button></div>`)
	default:
		panic("No edit form implementation for " + idPrefix + field.Name + " with type " + field.Type.String())
	}
}
