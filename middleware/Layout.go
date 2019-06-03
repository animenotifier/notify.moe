package middleware

import (
	"sort"
	"strings"

	"github.com/aerogo/aero"
	"github.com/akyoto/stringutils/unsafe"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Layout middleware modifies the response body
// to be wrapped around the general layout.
func Layout(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		if ctx.Request().Method() != "GET" || !strings.Contains(ctx.Request().Header("Accept"), "text/html") || strings.HasPrefix(ctx.Path(), "/_") || strings.HasPrefix(ctx.Path(), "/api/") || strings.HasPrefix(ctx.Path(), "/graphql") {
			return next(ctx)
		}

		ctx.AddModifier(func(content []byte) []byte {
			user := arn.GetUserFromContext(ctx)
			customCtx := ctx.(*OpenGraphContext)
			openGraph := customCtx.OpenGraph

			// Make output order deterministic to profit from Aero caching.
			// To do this, we need to create slices and sort the tags.
			var meta []string
			var tags []string

			if openGraph != nil {
				for name := range openGraph.Meta {
					meta = append(meta, name)
				}

				sort.Strings(meta)

				for name := range openGraph.Tags {
					tags = append(tags, name)
				}

				sort.Strings(tags)
			}

			html := components.Layout(ctx, user, openGraph, meta, tags, unsafe.BytesToString(content))
			return unsafe.StringToBytes(html)
		})

		return next(ctx)
	}
}
