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

	JSON gjson.Result
}

func Main(app *aero.Application, Log *log.Log) {
	app.Post("/tokenapi", func(ctx aero.Context) error {
		response, err := ctx.Request().Body().String()
		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Couldn't get body", err)
		}
		if !gjson.Valid(response) {
			return ctx.Error(http.StatusBadRequest, "Couldn't parse JSON")
		}

		request := &TokenRequest{}
		request.JSON = gjson.Parse(response)

		token, err := uuid.Parse(request.JSON.Get("token").String())
		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Couldn't parse token", err)
		}
		request.Token = token
		request.User = *GetUserFromToken(token)

		action := request.JSON.Get("action").String()
		if action == "UpdateAnime" {
			err := AnimeUpdate(request)
			if err != nil {
				ctx.Error(http.StatusBadRequest, "Error updating anime:", err)
			}
		}

		return ctx.Error(http.StatusAccepted, "Processed request for", request.User.CleanNick())
	})
}
