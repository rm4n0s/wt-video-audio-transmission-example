package main

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/rm4n0s/go-observer"
	"github.com/rm4n0s/wt-video-audio-transmission-example/app/containers"
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
	ctx := context.Background()
	router := observer.NewProperty(Route{Page: PageNone})
	a := app.New()
	w := a.NewWindow("wt-example")
	w.Resize(fyne.NewSize(640, 480))
	settingsCtn := containers.NewSettingsContainer()
	settingsCtn.Submittion.Observe().OnChange(ctx, func(value containers.Settings) {
		err := validateSettings(value)
		if err != nil {
			settingsCtn.Error.Update(err.Error())
			return
		}
		router.Update(Route{Page: PageRoomCtn, Params: map[string]string{
			"host":     value.Host,
			"username": value.Username,
		}})
	})
	roomCtn := containers.NewRoomContainer()

	router.Observe().OnChange(ctx, func(value Route) {
		switch value.Page {
		case PageSettingsCtn:
			w.SetContent(settingsCtn.Container(ctx))

		case PageRoomCtn:
			w.SetContent(roomCtn.Container(value.Params))

		}
	})

	router.Update(Route{Page: PageSettingsCtn})

	w.ShowAndRun()
}
