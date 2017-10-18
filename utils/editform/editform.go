package editform

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Render ...
func Render(obj interface{}, title string, user *arn.User) string {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()
	id := reflect.Indirect(v.FieldByName("ID"))
	lowerCaseTypeName := strings.ToLower(t.Name())
	endpoint := `/api/` + lowerCaseTypeName + `/` + id.String()

	var b bytes.Buffer

	b.WriteString(`<div class="widget-form">`)
	b.WriteString(`<div class="widget" data-api="` + endpoint + `">`)

	b.WriteString(`<h1>`)
	b.WriteString(title)
	b.WriteString(`</h1>`)

	RenderObject(&b, obj, "")

	if user != nil && (user.Role == "editor" || user.Role == "admin") {
		b.WriteString(`<div class="buttons">`)
		b.WriteString(`<div class="buttons"><button class="action" data-action="publish" data-trigger="click">` + utils.Icon("share-alt") + `Publish</button></div>`)
		b.WriteString(`<button class="action" data-action="deleteObject" data-trigger="click" data-return-path="/` + lowerCaseTypeName + "s" + `" data-confirm-type="` + lowerCaseTypeName + `">` + utils.Icon("trash") + `Delete</button>`)
		b.WriteString(`</div>`)
	}

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
		if field.Tag.Get("type") == "textarea" {
			b.WriteString(components.InputTextArea(idPrefix+field.Name, fieldValue.String(), field.Name, ""))
		} else {
			b.WriteString(components.InputText(idPrefix+field.Name, fieldValue.String(), field.Name, ""))
		}
	case "[]string":
		b.WriteString(components.InputTags(idPrefix+field.Name, fieldValue.Interface().([]string), field.Name, field.Tag.Get("tooltip")))
	case "bool":
		if field.Name == "IsDraft" {
			return
		}
	case "[]*arn.ExternalMedia":
		for sliceIndex := 0; sliceIndex < fieldValue.Len(); sliceIndex++ {
			b.WriteString(`<div class="widget-section">`)
			b.WriteString(`<div class="widget-title">` + strconv.Itoa(sliceIndex+1) + ". " + field.Name + `</div>`)

			arrayObj := fieldValue.Index(sliceIndex).Interface()
			arrayIDPrefix := fmt.Sprintf("%s[%d].", field.Name, sliceIndex)
			RenderObject(b, arrayObj, arrayIDPrefix)

			// Preview
			b.WriteString(components.ExternalMedia(fieldValue.Index(sliceIndex).Interface().(*arn.ExternalMedia)))

			// Remove button
			b.WriteString(`<div class="buttons"><button class="action" data-action="arrayRemove" data-trigger="click" data-field="` + field.Name + `" data-index="`)
			b.WriteString(strconv.Itoa(sliceIndex))
			b.WriteString(`">` + utils.RawIcon("trash") + `</button></div>`)

			b.WriteString(`</div>`)
		}

		b.WriteString(`<div class="buttons">`)
		b.WriteString(`<button class="action" data-action="arrayAppend" data-trigger="click" data-field="` + field.Name + `">` + utils.Icon("plus") + `Add ` + field.Name + `</button>`)
		b.WriteString(`</div>`)
	default:
		panic("No edit form implementation for " + idPrefix + field.Name + " with type " + field.Type.String())
	}
}
