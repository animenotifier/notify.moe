package companies

import (
	"sort"
	"strings"
	"unicode"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// All renders an index of all companies.
func All(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

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

		if !unicode.IsLetter([]rune(firstLetter)[0]) {
			continue
		}

		if firstLetter != previousFirstLetter {
			groups = append(groups, []*arn.Company{})
			currentGroupIndex++
			previousFirstLetter = firstLetter
		}

		groups[currentGroupIndex] = append(groups[currentGroupIndex], company)
	}

	return ctx.HTML(components.CompaniesIndex(groups, user))
}
