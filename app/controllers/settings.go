package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/rm4n0s/go-observer"
	"github.com/rm4n0s/wt-video-audio-transmission-example/app/utils"
)

type Settings struct {
	Host     string
	Username string
}

type SettingsController struct {
	Submittion observer.Property[Settings]
	Error      observer.Property[string]
}

func NewSettingsController() *SettingsController {
	return &SettingsController{
		Submittion: observer.NewProperty(Settings{}),
		Error:      observer.NewProperty(""),
	}
}

func (sc *SettingsController) Container(ctx context.Context) fyne.CanvasObject {
	hostEntry := widget.NewEntry()
	hostEntry.Text = "https://localhost:4433"
	usernameEntry := widget.NewEntry()
	submitBtn := widget.NewButton("Connect", func() {
		sc.Submittion.Update(Settings{
			Host:     hostEntry.Text,
			Username: usernameEntry.Text,
		})
	})
	errLabel := canvas.NewText("", color.RGBA{255, 0, 0, 255})
	sc.Error.Observe().OnChange(ctx, func(value string) {
		str := utils.StringToMultiline(value, 70)
		errLabel.Text = str
	})
	formBox := container.New(layout.NewFormLayout(),
		widget.NewLabel("Host"), hostEntry,
		widget.NewLabel("Username"), usernameEntry,
		submitBtn, layout.NewSpacer(),
		layout.NewSpacer(), errLabel,
	)
	return formBox
}

func (sc *SettingsController) ValidateSettings(sets Settings) error {
	if sets.Username == "" {
		return errors.New("username is empty")
	}
	_, err := url.ParseRequestURI(sets.Host)
	if err != nil {
		return fmt.Errorf("failed to parse host: %w", err)
	}
	return nil
}
