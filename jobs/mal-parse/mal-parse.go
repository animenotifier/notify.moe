package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Parsing MAL files")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Invoke via parameters
	if InvokeShellArgs() {
		return
	}

	if objectType == "all" || objectType == "anime" {
		readFiles(path.Join(arn.Root, "jobs", "mal-download", "anime"), readAnimeFile)
	}

	if objectType == "all" || objectType == "character" {
		readFiles(path.Join(arn.Root, "jobs", "mal-download", "character"), readCharacterFile)
	}
}

// Read files in a given directory and apply a function on them
func readFiles(root string, onFile func(string) error) {
	count := 0

	filepath.Walk(root, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			color.Red(err.Error())
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(name, ".html.gz") {
			return nil
		}

		count++
		err = onFile(name)

		if err != nil {
			color.Red(err.Error())
		}

		// Always continue traversing the directory
		return nil
	})

	color.Cyan("%d files found", count)
}
