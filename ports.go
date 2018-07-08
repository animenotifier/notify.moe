package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

// testConnectivity will test if port 443 is accessible or not.
// If not, it will attempt to run "sudo make ports" and consume
// execution once the user has entered the password.
func testConnectivity() {
	config := tls.Config{
		InsecureSkipVerify: true,
	}

	_, err := tls.Dial("tcp", ":443", &config)

	if err != nil {
		fmt.Println("--------------------------------------------------------------------------------")
		color.Red("HTTPS port 443 is not accessible")
		fmt.Println("Running", color.YellowString("make ports"), "to be able to access", color.GreenString("https://"+app.Config.Domain))

		cmd := exec.Command("sudo", "make", "ports")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	}
}
