package middleware

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// OpenGraphContext is a context with open graph data.
type OpenGraphContext struct {
	aero.Context
	*arn.OpenGraph
}

// OpenGraph middleware modifies the context to be an OpenGraphContext.
func OpenGraph(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		openGraphCtx := &OpenGraphContext{
			Context:   ctx,
			OpenGraph: nil,
		}

		return next(openGraphCtx)
	}
}
