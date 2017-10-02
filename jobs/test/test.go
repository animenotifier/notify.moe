package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "test", "github.com/animenotifier/notify.moe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		panic(err)
	}

	cmd.Wait()
}
