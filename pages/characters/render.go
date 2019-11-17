package characters

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

const (
	charactersFirstLoad = 104
	charactersPerScroll = 39
)

// render renders the characters page with the given characters.
func render(ctx aero.Context, allCharacters []*arn.Character) error {
	user := arn.GetUserFromContext(ctx)
	index, _ := ctx.GetInt("index")
	tag := ctx.Get("tag")

	// Slice the part that we need
	characters := allCharacters[index:]
	maxLength := charactersFirstLoad

	if index > 0 {
		maxLength = charactersPerScroll
	}

	if len(characters) > maxLength {
		characters = characters[:maxLength]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allCharacters), maxLength, index)

	// In case we're scrolling, send Characters only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.CharactersScrollable(characters, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.Characters(characters, nextIndex, tag, user))
}
