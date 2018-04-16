package utils

// YearRegistrations describes how many user registrations happened in a year.
type YearRegistrations struct {
	Total  int
	Months map[int]int
}
