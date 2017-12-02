package companies

import (
	"sort"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// All renders the companies page.
func All(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	companies := arn.FilterCompanies(func(company *arn.Company) bool {
		return !company.IsDraft
	})

	sort.Slice(companies, func(i, j int) bool {
		return strings.ToLower(companies[i].Name.English) < strings.ToLower(companies[j].Name.English)
	})

	groups := [][]*arn.Company{}
	currentGroupIndex := -1

	previousFirstLetter := ""

	for _, company := range companies {
		firstLetter := strings.ToLower(company.Name.English[:1])

		if firstLetter != previousFirstLetter {
			groups = append(groups, []*arn.Company{})
			currentGroupIndex++
			previousFirstLetter = firstLetter
		}

		groups[currentGroupIndex] = append(groups[currentGroupIndex], company)
	}

	return ctx.HTML(components.Companies(groups, user))
}
