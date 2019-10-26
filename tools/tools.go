package tools

import (
	// pack to generate our assets
	_ "github.com/aerogo/pack"

	// run as a development server that restarts on changes
	_ "github.com/aerogo/run"

	// jq as a platform-independent JSON parser to install the database
	_ "github.com/itchyny/gojq"
)
