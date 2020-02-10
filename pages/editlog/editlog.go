package editlog

import (
	"github.com/aerogo/aero"
)

const (
	entriesFirstLoad = 120
	entriesPerScroll = 40
)

// Full edit log.
func Full(ctx aero.Context) error {
	return render(ctx, false)
}

// Compact edit log.
func Compact(ctx aero.Context) error {
	return render(ctx, true)
}
