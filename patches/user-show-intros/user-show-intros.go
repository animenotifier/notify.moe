package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Showing user intros")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	data, _ := ioutil.ReadFile("bad-words-list.txt")
	fullText := string(data)
	badWords := strings.Split(fullText, ",")

	for user := range arn.StreamUsers() {
		if user.Introduction == "" {
			continue
		}

		for _, badWord := range badWords {
			pos := strings.Index(user.Introduction, badWord)

			if pos != -1 && (pos == 0 || !unicode.IsLetter(rune(user.Introduction[pos-1]))) && (pos+len(badWord) == len(user.Introduction) || !unicode.IsLetter(rune(user.Introduction[pos+len(badWord)]))) {
				color.Cyan(user.Nick)
				color.Red(badWord)
				fmt.Println(user.Introduction)
				break
			}
		}
	}
}
