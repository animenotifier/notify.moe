package arn

// IndexOf ...
func IndexOf(collection []string, t string) int {
	for i, v := range collection {
		if v == t {
			return i
		}
	}
	return -1
}

// Contains ...
func Contains(collection []string, t string) bool {
	return IndexOf(collection, t) >= 0
}

// func Any(collection []string, f func(string) bool) bool {
// 	for _, v := range collection {
// 		if f(v) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func All(collection []string, f func(string) bool) bool {
// 	for _, v := range collection {
// 		if !f(v) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func Filter(collection []string, f func(string) bool) []string {
// 	vsf := make([]string, 0)
// 	for _, v := range collection {
// 		if f(v) {
// 			vsf = append(vsf, v)
// 		}
// 	}
// 	return vsf
// }

// func Map(collection []string, f func(string) string) []string {
// 	vsm := make([]string, len(collection))
// 	for i, v := range collection {
// 		vsm[i] = f(v)
// 	}
// 	return vsm
// }
