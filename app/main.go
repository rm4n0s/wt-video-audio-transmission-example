package main

import (
	"context"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/rm4n0s/go-observer"
	"github.com/rm4n0s/wt-video-audio-transmission-example/app/controllers"
)

type Page string

const (
	PageNone        = Page("none")
	PageSettingsCtn = Page("settings")
	PageRoomCtn     = Page("room")
)

type Route struct {
	Page   Page
	Params map[string]string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx := context.Background()
	router := observer.NewProperty(Route{Page: PageNone})
	a := app.New()

	w := a.NewWindow("wt-example")
	w.Resize(fyne.NewSize(640, 480))
	settingsCtrl := controllers.NewSettingsController()
	roomCtrl := controllers.NewRoomController()
	router.Observe().OnChange(ctx, func(value Route) {
		switch value.Page {
		case PageSettingsCtn:
			w.SetContent(settingsCtrl.Container(ctx))

		case PageRoomCtn:
			w.SetContent(roomCtrl.Container(value.Params))

		}
	})
	settingsCtrl.Submittion.Observe().OnChange(ctx, func(value controllers.Settings) {
		err := settingsCtrl.ValidateSettings(value)
		if err != nil {
			settingsCtrl.Error.Update(err.Error())
			return
		}
		err = roomCtrl.InitConns(value.Host, value.Username)
		if err != nil {
			settingsCtrl.Error.Update(err.Error())
			return
		}
		router.Update(Route{Page: PageRoomCtn, Params: map[string]string{
			"host":     value.Host,
			"username": value.Username,
		}})
	})

	roomCtrl.Closed.Observe().OnChange(ctx, func(value bool) {
		router.Update(Route{Page: PageSettingsCtn})
	})

	router.Update(Route{Page: PageSettingsCtn})

	w.ShowAndRun()
}
