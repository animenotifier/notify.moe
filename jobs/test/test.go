package main

import (
	"os/exec"
	"sync"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/akyoto/color"
)

var packages = []string{
	"github.com/animenotifier/notify.moe",
	"github.com/animenotifier/notify.moe/arn",
	"github.com/animenotifier/kitsu",
	"github.com/animenotifier/anilist",
	"github.com/animenotifier/mal",
	"github.com/animenotifier/shoboi",
	"github.com/animenotifier/twist",
	// "github.com/animenotifier/japanese",
	// "github.com/animenotifier/osu",
}

func main() {
	defer color.Green("Finished.")

	wg := sync.WaitGroup{}

	for _, pkg := range packages {
		wg.Add(1)

		go func(pkgLocal string) {
			testPackage(pkgLocal)
			wg.Done()
		}(pkg)
	}

	wg.Wait()
}

func testPackage(pkg string) {
	cmd := exec.Command("go", "test", pkg+"/...")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		panic(err)
	}

	err = cmd.Wait()

	if err != nil {
		color.Red("%s", pkg)

		// Send notification to the admin
		admin, _ := arn.GetUser("4J6qpK1ve")
		admin.SendNotification(&arn.PushNotification{
			Title:   pkg,
			Message: "Test failed",
			Link:    "https://" + pkg,
			Icon:    "https://media.notify.moe/images/brand/220.png",
			Type:    arn.NotificationTypePackageTest,
		})
		return
	}

	color.Green("%s", pkg)
}
