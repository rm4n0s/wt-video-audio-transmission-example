package components

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Settings struct {
	Host     string
	Username string
}

func ConnectionSettingsComponent(submit func(settings Settings) error, next func()) *gtk.Grid {
	grid := gtk.NewGrid()
	hostLabel := gtk.NewLabel("Host")
	grid.Attach(hostLabel, 0, 0, 1, 1)

	hostText := gtk.NewText()
	hostText.SetCSSClasses([]string{"settingText"})
	hostText.SetText("http://localhost:4433")
	grid.Attach(hostText, 1, 0, 1, 1)

	usernameLabel := gtk.NewLabel("Username")
	grid.Attach(usernameLabel, 0, 1, 1, 1)

	usernameText := gtk.NewText()
	usernameText.SetCSSClasses([]string{"settingText"})
	grid.Attach(usernameText, 1, 1, 1, 1)

	connectBtn := gtk.NewButton()
	connectBtn.SetLabel("connect")
	connectBtn.ConnectClicked(func() {
		err := submit(Settings{
			Host:     hostText.Text(),
			Username: usernameText.Text(),
		})
		if err == nil {
			next()
		} else {
			errLabel := gtk.NewLabel("Error:" + err.Error())
			errLabel.SetCSSClasses([]string{"error"})
			errLabel.SetXAlign(0)
			errLabel.SetYAlign(0)
			grid.Attach(errLabel, 0, 3, 1, 1)
		}
	})
	grid.Attach(connectBtn, 0, 2, 1, 1)

	return grid
}
