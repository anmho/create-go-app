package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

type options struct {
	appName string
	appTemplate string
}

func NewOptions() options {
	return options{
		appName: "my-go-app",
	}
}

var (
	confirmConfig bool
)

type AppTemplate string

const (
	HumaFlyIO       AppTemplate = "huma-flyio"
	ConnectCloudRun AppTemplate = "connect-cloudrun"
)

func main() {
	
	opts := NewOptions()
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the name of your app?").
				Placeholder("my-go-app").
				Value(&opts.appName),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your API Framework (API framework + PaaS)").
				Options(
					huh.NewOption("HumaRocks + fly.io", string(HumaFlyIO)),
					huh.NewOption("ConnectRPC + Google Cloud Run", string(ConnectCloudRun)),
				).
				Value(&opts.appTemplate), // store the chosen option in the "burger" variable
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Confirm configuration?").
				Description(fmt.Sprintf(
`
	Name: %s
	Template: %s`, opts.appName, opts.appTemplate)).
				Value(&confirmConfig),
		),
	)
	
	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	// generateTemplatedMain(opts)
	generateTemplatedAPI(opts)
}
