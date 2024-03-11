package containers

import (
	"context"

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

type SettingsContainer struct {
	Submittion observer.Property[Settings]
	Error      observer.Property[string]
}

func NewSettingsContainer() *SettingsContainer {
	return &SettingsContainer{
		Submittion: observer.NewProperty(Settings{}),
		Error:      observer.NewProperty(""),
	}
}

func (sc *SettingsContainer) Container(ctx context.Context) fyne.CanvasObject {
	hostEntry := widget.NewEntry()
	hostEntry.Text = "https://localhost:4433"
	usernameEntry := widget.NewEntry()
	submitBtn := widget.NewButton("Submit", func() {
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
