package companies

import "github.com/animenotifier/notify.moe/arn"

// fetchAll returns all companies
func fetchAll() []*arn.Company {
	return arn.FilterCompanies(func(company *arn.Company) bool {
		return !company.IsDraft
	})
}
