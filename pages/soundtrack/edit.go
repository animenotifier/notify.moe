package soundtrack

import (
	"bytes"
	"net/http"
	"reflect"
	"strings"

	"github.com/animenotifier/notify.moe/components"

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
			"og:title":     track.Media[0].Title,
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
	lowerCaseTypeName := strings.ToLower(t.Name())
	id := reflect.Indirect(v.FieldByName("ID"))

	var b bytes.Buffer
	b.WriteString(`<div class="widget-form">`)
	b.WriteString(`<div class="widget" data-api="/api/` + lowerCaseTypeName + `/` + id.String() + `">`)
	b.WriteString(`<h1>`)
	b.WriteString(title)
	b.WriteString(`</h1>`)

	// Fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Anonymous || field.Tag.Get("editable") != "true" {
			continue
		}

		fieldValue := reflect.Indirect(v.FieldByName(field.Name))

		switch field.Type.String() {
		case "string":
			b.WriteString(components.InputText(field.Name, fieldValue.String(), field.Name, ""))
		case "[]string":
			b.WriteString(components.InputTags(field.Name, fieldValue.Interface().([]string), field.Name))
		default:
			panic("No edit form implementation for " + field.Name + " with type " + field.Type.String())
		}
	}

	b.WriteString("</div>")
	b.WriteString("</div>")
	return b.String()
}
