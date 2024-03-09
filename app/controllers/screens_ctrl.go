package controllers

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/rm4n0s/go-observer"
)

type Picture struct {
	X      int
	Y      int
	Pix    []byte
	Stride int
}

type ScreenController struct {
	Cancelled observer.Property[bool]
}

func NewScreenController() *ScreenController {
	return &ScreenController{
		Cancelled: observer.NewProperty(false),
	}
}

func (sc *ScreenController) Page(username string) *gtk.Box {
	box := gtk.NewBox(gtk.OrientationVertical, 4)
	box.Append(gtk.NewLabel(username))
	pic := gtk.NewImage()
	pic.SetSizeRequest(640, 480)
	box.Append(pic)
	cancelBtn := gtk.NewButton()
	cancelBtn.SetLabel("Cancel")
	cancelBtn.ConnectClicked(func() {
		sc.Cancelled.Update(true)
	})
	box.Append(cancelBtn)
	return box
}
