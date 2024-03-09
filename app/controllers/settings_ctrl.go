package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/rm4n0s/go-observer"
	"github.com/rm4n0s/wt-video-audio-transmission-example/utils"
)

type SettingsController struct {
	Host      observer.Property[string]
	Username  observer.Property[string]
	Submitted observer.Property[bool]
	Error     observer.Property[error]
	HasError  observer.Property[bool]
}

func NewSettingController() *SettingsController {
	var err error
	return &SettingsController{
		Host:      observer.NewProperty("https://localhost:4433"),
		Username:  observer.NewProperty(""),
		Submitted: observer.NewProperty(false),
		Error:     observer.NewProperty(err),
		HasError:  observer.NewProperty(false),
	}
}

func (sc *SettingsController) Page(ctx context.Context) *gtk.Grid {
	grid := gtk.NewGrid()
	hostLabel := gtk.NewLabel("Host")
	grid.Attach(hostLabel, 0, 0, 1, 1)

	hostEntry := gtk.NewEntry()
	hostEntry.SetCSSClasses([]string{"settingEntry"})
	hostEntry.SetText(sc.Host.Value())
	hostEntry.ConnectChanged(func() {
		sc.Host.Update(hostEntry.Text())
	})
	grid.Attach(hostEntry, 1, 0, 1, 1)

	usernameLabel := gtk.NewLabel("Username")
	grid.Attach(usernameLabel, 0, 1, 1, 1)

	usernameEntry := gtk.NewEntry()
	usernameEntry.SetCSSClasses([]string{"settingEntry"})
	usernameEntry.SetText(sc.Username.Value())
	usernameEntry.ConnectChanged(func() {
		sc.Username.Update(usernameEntry.Text())
	})
	grid.Attach(usernameEntry, 1, 1, 1, 1)

	connectBtn := gtk.NewButton()
	connectBtn.SetLabel("connect")
	connectBtn.ConnectClicked(func() {
		sc.Submitted.Update(true)
	})
	grid.Attach(connectBtn, 0, 2, 1, 1)
	errLabel := gtk.NewLabel("")
	errLabel.SetCSSClasses([]string{"error"})
	grid.Attach(errLabel, 1, 3, 1, 1)
	sc.HasError.Observe().OnChange(ctx, func(value bool) {
		if value {
			glib.IdleAdd(func() {
				str := sc.Error.Value().Error()
				multiline := utils.StringToMultiline(str, 20)
				errLabel.SetText(multiline)
			})
		}
	})

	return grid
}

func (sc *SettingsController) ValidateSettings() error {
	if sc.Username.Value() == "" {
		return errors.New("username is empty")
	}
	_, err := url.ParseRequestURI(sc.Host.Value())
	if err != nil {
		return fmt.Errorf("failed to parse host: %w", err)
	}
	return nil
}
