package shop

import (
	"io/ioutil"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

var proAccount = ""

func init() {
	data, _ := ioutil.ReadFile("pages/shop/pro-account.md")
	proAccount = string(data)
}

// Get shop page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	return ctx.HTML(components.Shop(user, proAccount))
}
