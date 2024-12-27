package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

type options struct {
	AppName string
	ModuleName string
	AppTemplate string
}

func NewOptions() options {
	return options{
		AppName: "my-go-app",
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
				Value(&opts.AppName),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your API Framework (API framework + PaaS)").
				Options(
					huh.NewOption("ConnectRPC + Google Cloud Run", string(ConnectCloudRun)),
					huh.NewOption("HumaRocks + fly.io", string(HumaFlyIO)),
				).
				Value(&opts.AppTemplate),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Confirm configuration?").
				Description(fmt.Sprintf(
`
	Name: %s
	Template: %s`, opts.AppName, opts.AppTemplate)).
				Value(&confirmConfig),
		),
	)
	
	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	opts.ModuleName=opts.AppName

	err := generateTemplatedAPI(opts)
	if err != nil {
		panic(err)
	}
}
