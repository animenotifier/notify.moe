package upload

import (
	"bytes"
	"fmt"
	"image"
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

	// Decode
	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid image format", err)
	}

	fmt.Println("Avatar received!", len(data), format, img.Bounds().Dx(), img.Bounds().Dy(), user.Nick)
	// ioutil.WriteFile("avatar")
	return "ok"
}
