package best

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxEntries = 7

// Get search page.
func Get(ctx *aero.Context) string {
	overall, err := arn.GetListOfAnimeCached("best anime overall")

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching popular anime", err)
	}

	story, err := arn.GetListOfAnimeCached("best anime story")

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching popular anime", err)
	}

	visuals, err := arn.GetListOfAnimeCached("best anime visuals")

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching popular anime", err)
	}

	soundtrack, err := arn.GetListOfAnimeCached("best anime soundtrack")

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching popular anime", err)
	}

	airing, err := arn.GetAiringAnimeCached()

	if err != nil {
		return ctx.Error(500, "Couldn't fetch airing anime", err)
	}

	return ctx.HTML(components.BestAnime(overall[:maxEntries], story[:maxEntries], visuals[:maxEntries], soundtrack[:maxEntries], airing[:maxEntries]))
}
