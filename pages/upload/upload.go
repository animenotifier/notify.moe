package upload

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Avatar ...
func Avatar(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	data, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Reading request body failed", err)
	}

	format, err := guessImageFormat(bytes.NewReader(data))

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Could not determine image file type", err)
	}

	fmt.Println("Avatar received!", len(data), format, user.Nick)
	// ioutil.WriteFile("avatar")
	return "ok"
}

// Guess image format from gif/jpeg/png/webp
func guessImageFormat(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	return format, err
}
