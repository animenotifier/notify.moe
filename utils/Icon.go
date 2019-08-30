package utils

import (
	"fmt"
)

// Icon shows an icon that has a margin-right attribute.
func Icon(name string) string {
	return fmt.Sprintf("<svg-icon name='%s' class='padded-icon'></svg-icon>", name)
}

// RawIcon shows the raw icon without any additional margin.
func RawIcon(name string) string {
	return fmt.Sprintf("<svg-icon name='%s'></svg-icon>", name)
}
