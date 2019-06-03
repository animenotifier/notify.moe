package osutils

import (
	"os"
)

// Exists tells you whether the given file or directory exists.
func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// // WaitExists will wait until the given file path exists.
// func WaitExists(filePath string, timeout time.Duration, pollingInterval time.Duration) {
// 	start := time.Now()

// 	for !Exists(filePath) && time.Since(start) < timeout {
// 		time.Sleep(pollingInterval)
// 	}
// }
