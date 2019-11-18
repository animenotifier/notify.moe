package database

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/akyoto/stringutils/unsafe"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/mohae/deepcopy"
)

// Download downloads a snapshot of a database collection.
func Download(ctx aero.Context) error {
	typ := ctx.Get("type")

	if !arn.DB.HasType(typ) {
		return ctx.Error(http.StatusNotFound, "Type doesn't exist")
	}

	if arn.IsPrivateType(typ) {
		return ctx.Error(http.StatusUnauthorized, "Type is private and can not be downloaded")
	}

	// Send headers necessary for file downloads
	ctx.Response().SetHeader("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.dat"`, typ))

	// Stream data
	reader, writer := io.Pipe()
	encoder := json.NewEncoder(writer)

	go func() {
		for object := range arn.DB.All(typ) {
			idObject, hasID := object.(arn.Identifiable)

			if !hasID {
				continue
			}

			// Filter out private data
			filter, isFilter := object.(api.Filter)

			if isFilter && filter.ShouldFilter(ctx) {
				object = deepcopy.Copy(object)
				filter = object.(api.Filter)
				filter.Filter()
			}

			// Write ID
			_, err := writer.Write(unsafe.StringToBytes(idObject.GetID()))

			if err != nil {
				_ = writer.CloseWithError(err)
				return
			}

			// Write newline
			_, err = writer.Write([]byte("\n"))

			if err != nil {
				_ = writer.CloseWithError(err)
				return
			}

			// Write JSON (newline included)
			err = encoder.Encode(object)

			if err != nil {
				_ = writer.CloseWithError(err)
				return
			}
		}

		writer.Close()
	}()

	return ctx.Reader(reader)
}
