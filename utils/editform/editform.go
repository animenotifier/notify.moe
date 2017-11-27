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

// Render renders a generic editing UI for any kind of datatype that has an ID.
func Render(obj interface{}, title string, user *arn.User) string {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()
	id := reflect.Indirect(v.FieldByName("ID"))
	lowerCaseTypeName := strings.ToLower(t.Name())
	endpoint := `/api/` + lowerCaseTypeName + `/` + id.String()

	var b bytes.Buffer

	b.WriteString(`<div class="widget-form">`)
	b.WriteString(`<div class="widget" data-api="` + endpoint + `">`)

	// Title
	b.WriteString(`<h1>`)
	b.WriteString(title)
	b.WriteString(`</h1>`)

	// Render the object with its fields
	RenderObject(&b, obj, "")

	// Additional buttons when logged in
	if user != nil {
		b.WriteString(`<div class="buttons">`)

		_, ok := t.FieldByName("IsDraft")

		if ok {
			isDraft := v.FieldByName("IsDraft").Interface().(bool)

			if isDraft {
				b.WriteString(`<div class="buttons"><button class="action" data-action="publish" data-trigger="click">` + utils.Icon("share-alt") + `Publish</button></div>`)
			}
		}

		if user.Role == "editor" || user.Role == "admin" {
			b.WriteString(`<button class="action" data-action="deleteObject" data-trigger="click" data-return-path="/` + lowerCaseTypeName + "s" + `" data-confirm-type="` + lowerCaseTypeName + `">` + utils.Icon("trash") + `Delete</button>`)
		}

		b.WriteString(`</div>`)
	}

	b.WriteString("</div>")
	b.WriteString("</div>")

	return b.String()
}

// RenderObject renders the UI for the object into the bytes buffer and appends an ID prefix for all API requests.
// The ID prefix should either be empty or end with a dot character.
func RenderObject(b *bytes.Buffer, obj interface{}, idPrefix string) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

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
	fieldType := field.Type.String()

	// String
	if fieldType == "string" {
		if field.Tag.Get("datalist") != "" {
			dataList := field.Tag.Get("datalist")
			values := arn.DataLists[dataList]
			b.WriteString(components.InputSelection(idPrefix+field.Name, fieldValue.String(), field.Name, field.Tag.Get("tooltip"), values))
		} else if field.Tag.Get("type") == "textarea" {
			b.WriteString(components.InputTextArea(idPrefix+field.Name, fieldValue.String(), field.Name, field.Tag.Get("tooltip")))
		} else {
			b.WriteString(components.InputText(idPrefix+field.Name, fieldValue.String(), field.Name, field.Tag.Get("tooltip")))
		}

		return
	}

	// Bool
	if fieldType == "bool" {
		if field.Name == "IsDraft" {
			return
		}

		// TODO: Render bool type
		return
	}

	// Array of strings
	if fieldType == "[]string" {
		b.WriteString(components.InputTags(idPrefix+field.Name, fieldValue.Interface().([]string), field.Name, field.Tag.Get("tooltip")))
		return
	}

	// Any kind of array
	if strings.HasPrefix(fieldType, "[]") {
		b.WriteString(`<div class="widget-section">`)
		b.WriteString(`<h3 class="widget-title">`)
		b.WriteString(field.Name)
		b.WriteString(`</h3>`)

		for sliceIndex := 0; sliceIndex < fieldValue.Len(); sliceIndex++ {
			b.WriteString(`<div class="widget-section">`)

			b.WriteString(`<div class="widget-title">`)

			// Title
			b.WriteString(strconv.Itoa(sliceIndex+1) + ". " + field.Name)
			b.WriteString(`<div class="spacer"></div>`)

			// Remove button
			b.WriteString(`<button class="action" data-action="arrayRemove" data-trigger="click" data-field="` + field.Name + `" data-index="`)
			b.WriteString(strconv.Itoa(sliceIndex))
			b.WriteString(`">` + utils.RawIcon("trash") + `</button>`)

			b.WriteString(`</div>`)

			arrayObj := fieldValue.Index(sliceIndex).Interface()
			arrayIDPrefix := fmt.Sprintf("%s[%d].", field.Name, sliceIndex)
			RenderObject(b, arrayObj, arrayIDPrefix)

			// Preview
			// elementValue := fieldValue.Index(sliceIndex)
			// RenderArrayElement(b, &elementValue)
			if fieldType == "[]*arn.ExternalMedia" {
				b.WriteString(components.ExternalMedia(fieldValue.Index(sliceIndex).Interface().(*arn.ExternalMedia)))
			}

			b.WriteString(`</div>`)
		}

		b.WriteString(`<div class="buttons">`)
		b.WriteString(`<button class="action" data-action="arrayAppend" data-trigger="click" data-field="` + field.Name + `">` + utils.Icon("plus") + `Add ` + field.Name + `</button>`)
		b.WriteString(`</div>`)

		b.WriteString(`</div>`)
		return
	}

	// Any custom field type will be recursively rendered via another RenderObject call
	b.WriteString(`<div class="widget-section">`)
	b.WriteString(`<h3 class="widget-title">` + field.Name + `</h3>`)

	// Indent the fields
	b.WriteString(`<div class="indent">`)
	RenderObject(b, fieldValue.Interface(), field.Name+".")
	b.WriteString(`</div>`)

	b.WriteString(`</div>`)
}
