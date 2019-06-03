package amvs

import "github.com/animenotifier/notify.moe/arn"

func fetchAll() []*arn.AMV {
	return arn.FilterAMVs(func(amv *arn.AMV) bool {
		return !amv.IsDraft
	})
}
