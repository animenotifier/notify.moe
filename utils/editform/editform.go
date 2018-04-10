package editform

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/aerogo/api"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Render renders a generic editing UI for any kind of datatype that has an ID.
func Render(obj interface{}, title string, user *arn.User) string {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()
	id := findMainID(t, v)
	lowerCaseTypeName := strings.ToLower(t.Name())
	endpoint := `/api/` + lowerCaseTypeName + `/` + id.String()

	var b bytes.Buffer

	b.WriteString(`<div class="widget-form">`)
	b.WriteString(`<div class="widget" data-api="` + endpoint + `">`)

	// Title
	b.WriteString(`<h1 class="mountable">`)
	b.WriteString(title)
	b.WriteString(`</h1>`)

	// Render the object with its fields
	RenderObject(&b, obj, "")

	// Additional buttons when logged in
	if user != nil {
		b.WriteString(`<div class="buttons">`)

		// Publish button
		_, ok := t.FieldByName("IsDraft")

		if ok {
			isDraft := v.FieldByName("IsDraft").Interface().(bool)

			if isDraft {
				b.WriteString(`<div class="buttons"><button class="mountable action" data-action="publish" data-trigger="click">` + utils.Icon("share-alt") + `Publish</button></div>`)
			}
		}

		// Delete button
		_, isDeletable := obj.(api.Deletable)

		if isDeletable && (user.Role == "editor" || user.Role == "admin") {
			returnPath := ""

			switch lowerCaseTypeName {
			case "anime":
				returnPath = "/explore"
			case "company":
				returnPath = "/companies"
			default:
				returnPath = "/" + lowerCaseTypeName + "s"
			}

			b.WriteString(`<button class="mountable action" data-action="deleteObject" data-trigger="click" data-return-path="` + returnPath + `" data-confirm-type="` + lowerCaseTypeName + `">` + utils.Icon("trash") + `Delete</button>`)
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
	fieldValue := reflect.Indirect(v.FieldByName(field.Name))

	// Embedded fields
	if field.Anonymous {
		RenderObject(b, fieldValue.Interface(), idPrefix)
		return
	}

	if field.Tag.Get("editable") != "true" {
		return
	}

	b.WriteString("<div class='mountable'>")
	defer b.WriteString("</div>")

	fieldType := field.Type.String()

	// String
	if fieldType == "string" {
		idType := field.Tag.Get("idType")

		// Try to infer the ID type by the field name
		if idType == "" {
			switch field.Name {
			case "AnimeID":
				idType = "Anime"

			case "CharacterID":
				idType = "Character"
			}
		}

		showPreview := idType != "" && fieldValue.String() != ""

		if showPreview {
			b.WriteString("<div class='widget-section-with-preview'>")
		}

		// Input field
		if field.Tag.Get("datalist") != "" {
			dataList := field.Tag.Get("datalist")
			values := arn.DataLists[dataList]
			b.WriteString(components.InputSelection(idPrefix+field.Name, fieldValue.String(), field.Name, field.Tag.Get("tooltip"), values))
		} else if field.Tag.Get("type") == "textarea" {
			b.WriteString(components.InputTextArea(idPrefix+field.Name, fieldValue.String(), field.Name, field.Tag.Get("tooltip")))
		} else {
			b.WriteString(components.InputText(idPrefix+field.Name, fieldValue.String(), field.Name, field.Tag.Get("tooltip")))
		}

		if showPreview {
			b.WriteString("<div class='widget-section-preview'>")
		}

		// Preview
		switch idType {
		case "Anime":
			animeID := fieldValue.String()
			anime, err := arn.GetAnime(animeID)

			if err == nil {
				b.WriteString(components.EditFormImagePreview(anime.Link(), anime.ImageLink("small"), true))
			}

		case "Character":
			characterID := fieldValue.String()
			character, err := arn.GetCharacter(characterID)

			if err == nil {
				b.WriteString(components.EditFormImagePreview(character.Link(), character.ImageLink("small"), false))
			}

		case "":
			break

		default:
			fmt.Println("Error: Unknown idType tag: " + idType)
		}

		// Close preview tags
		if showPreview {
			b.WriteString("</div></div>")
		}

		return
	}

	// Int
	if fieldType == "int" {
		b.WriteString(components.InputNumber(idPrefix+field.Name, float64(fieldValue.Int()), field.Name, field.Tag.Get("tooltip"), "", "", "1"))
		return
	}

	// Float
	if fieldType == "float64" {
		b.WriteString(components.InputNumber(idPrefix+field.Name, fieldValue.Float(), field.Name, field.Tag.Get("tooltip"), "", "", ""))
		return
	}

	// Bool
	if fieldType == "bool" {
		if field.Name == "IsDraft" {
			return
		}

		b.WriteString(components.InputBool(idPrefix+field.Name, fieldValue.Bool(), field.Name, field.Tag.Get("tooltip")))
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
			b.WriteString(`<button class="action" title="Delete this ` + field.Name + `" data-action="arrayRemove" data-trigger="click" data-field="` + field.Name + `" data-index="`)
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

// findMainID finds the main ID of the object.
func findMainID(t reflect.Type, v reflect.Value) reflect.Value {
	idField := v.FieldByName("ID")

	if idField.IsValid() {
		return reflect.Indirect(idField)
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("mainID") == "true" {
			return reflect.Indirect(v.Field(i))
		}
	}

	panic("Type " + t.Name() + " doesn't have a main ID!")
}
