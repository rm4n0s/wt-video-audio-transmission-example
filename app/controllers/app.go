package controllers

import (
	"context"
	"fmt"
	"os"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/rm4n0s/go-observer"
	"github.com/rm4n0s/wt-video-audio-transmission-example/app/vpclient"
)

type Page string

const (
	PageNone         = Page("none")
	PageSettingsCtrl = Page("settings")
	PageScreensCtrl  = Page("screens")
)

type Route struct {
	Page   Page
	Params map[string]string
}

type App struct {
	gtkApp *gtk.Application
	Router observer.Property[Route]
}

func NewApp(appID string) *App {
	app := &App{
		gtkApp: gtk.NewApplication(appID, gio.ApplicationFlagsNone),
		Router: observer.NewProperty(Route{Page: PageNone}),
	}
	return app

}

func (a *App) Run() {
	a.gtkApp.ConnectActivate(func() { a.activate() })
	if code := a.gtkApp.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func (a *App) activate() {
	InitCSS()
	ctx := context.Background()
	connCtx, cancelConn := context.WithCancel(ctx)
	window := gtk.NewApplicationWindow(a.gtkApp)
	window.SetTitle("wt-video-audio-transmission-example")
	settingsCtrl := NewSettingController()
	screensCtrl := NewScreenController()
	var inputClient *vpclient.VoipClient
	var outputClient *vpclient.VoipClient

	a.Router.Observe().OnChange(ctx, func(value Route) {
		switch value.Page {
		case PageSettingsCtrl:
			glib.IdleAdd(func() {
				settingsPage := settingsCtrl.Page(ctx)
				settingsCtrl.Submitted.Observe().OnChange(ctx, func(value bool) {
					err := settingsCtrl.ValidateSettings()
					if err != nil {
						settingsCtrl.HasError.Update(true)
						settingsCtrl.Error.Update(err)
						return
					}
					inputClient, outputClient, err = connectToVoip(connCtx, settingsCtrl.Host.Value(), settingsCtrl.Username.Value())
					if err != nil {
						settingsCtrl.HasError.Update(true)
						settingsCtrl.Error.Update(err)
						return
					}
					a.Router.Update(Route{Page: PageScreensCtrl, Params: map[string]string{
						"username": settingsCtrl.Username.Value(),
						"host":     settingsCtrl.Host.Value(),
					}})
				})
				window.SetChild(settingsPage)
			})
		case PageScreensCtrl:
			glib.IdleAdd(func() {
				fmt.Println(inputClient, outputClient)
				screensPage := screensCtrl.Page(value.Params["username"])
				screensCtrl.Cancelled.Observe().OnChange(ctx, func(value bool) {
					cancelConn()
					a.Router.Update(Route{Page: PageSettingsCtrl})
				})
				window.SetChild(screensPage)
			})
		}
	})

	a.Router.Update(Route{Page: PageSettingsCtrl})
	window.Show()
}
