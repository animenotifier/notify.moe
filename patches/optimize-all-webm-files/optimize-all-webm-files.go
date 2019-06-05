package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Optimizing all webm files in notify.moe AMV directory for fast streaming & seeking")
	color.Yellow("DO NOT RUN THIS COMMAND MULTIPLE TIMES")
	defer color.Green("Finished")

	readFiles(path.Join(arn.Root, "videos", "amvs"), mkclean)
}

// Optimize a webm file
func mkclean(file string) error {
	fmt.Println(file)
	optimizedFile := file + ".optimized"

	// Run mkclean
	cmd := exec.Command(
		"mkclean",
		"--doctype", "4",
		"--keep-cues",
		"--optimize",
		file,
		optimizedFile,
	)

	// View mkclean output in terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start()

	if err != nil {
		return err
	}

	err = cmd.Wait()

	if err != nil {
		return err
	}

	// Now delete the original file and replace it with the optimized file
	err = os.Remove(file)

	if err != nil {
		return err
	}

	return os.Rename(optimizedFile, file)
}

// Read files in a given directory and apply a function on them
func readFiles(root string, onFile func(string) error) {
	err := filepath.Walk(root, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			color.Red(err.Error())
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(name, ".webm") {
			return nil
		}

		err = onFile(name)

		if err != nil {
			color.Red(err.Error())
		}

		// Always continue traversing the directory
		return nil
	})

	if err != nil {
		color.Red(err.Error())
	}
}
