package utils

import (
	"bytes"
	"html"
	"strings"
)

// RenderQuoteText renders the given quote text.
func RenderQuoteText(text string) string {
	buffer := bytes.Buffer{}
	buffer.WriteString("<p>")

	lines := strings.Split(text, "\n")

	for index, line := range lines {
		buffer.WriteString("<span class='quote-line'>")
		buffer.WriteString(html.EscapeString(line))
		buffer.WriteString("</span>")

		if index != len(lines)-1 {
			buffer.WriteString("<br>")
		}
	}

	buffer.WriteString("</p>")
	return buffer.String()
}
