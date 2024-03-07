package controllers

import (
	"fmt"
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/rm4n0s/wt-video-audio-transmission-example/app/components"
)

type App struct {
	app *gtk.Application
}

func NewApp(appID string) *App {
	app := gtk.NewApplication(appID, gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })
	return &App{
		app: app,
	}

}

func (a *App) Run() {
	if code := a.app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	components.InitCSS()

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("wt-video-audio-transmission-example")

	// step 1 add settings for the connection
	submit := func(settings components.Settings) error {
		fmt.Println(settings)
		return fmt.Errorf("problem")
	}
	next := func() {}
	settingsPage := components.ConnectionSettingsComponent(submit, next)
	window.SetChild(settingsPage)
	window.Show()
}
