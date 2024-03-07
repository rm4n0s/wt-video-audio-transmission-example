package components

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

type Picture struct {
	X      int
	Y      int
	Pix    []byte
	Stride int
}

func UserScreenComponent(username string, frames <-chan Picture) *gtk.Box {
	box := gtk.NewBox(gtk.OrientationVertical, 4)
	box.Append(gtk.NewLabel(username))
	pic := gtk.NewImage()
	pic.SetSizeRequest(640, 480)
	box.Append(pic)
	return box
}
