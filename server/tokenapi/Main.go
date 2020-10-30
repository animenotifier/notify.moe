package tokenapi

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/log"
	"github.com/akyoto/uuid"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/tidwall/gjson"

	"net/http"
)

type TokenRequest struct {
	Token uuid.UUID
	User  arn.User

	Json gjson.Result
}

func Main(app *aero.Application, authLog *log.Log) {
	app.Post("/tokenapi", func(ctx aero.Context) error {
		response, err := ctx.Request().Body().String()
		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Couldn't get body", err)
		}
		if !gjson.Valid(response) {
			return ctx.Error(http.StatusBadRequest, "Couldn't parse JSON")
		}

		request := &TokenRequest{}
		request.Json = gjson.Parse(response)

		token, err := uuid.Parse(request.Json.Get("token").String())
		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Couldn't parse token", err)
		}
		request.Token = token

		action := request.Json.Get("action").String()
		if action == "UpdateAnime" {
			parameters := &AnimeParameters{}

			// @TODO

			AnimeUpdate(request, parameters)
		}

		return ctx.Error(http.StatusAccepted, "Processed.")
	})
}
