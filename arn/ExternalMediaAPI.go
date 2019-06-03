package arn

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ api.Creatable = (*ExternalMedia)(nil)
)

// Create sets the data for new external media.
func (media *ExternalMedia) Create(ctx aero.Context) error {
	media.Service = "Youtube"
	return nil
}
