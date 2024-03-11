package containers

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type RoomContainer struct{}

func NewRoomContainer() *RoomContainer {
	return &RoomContainer{}
}

func (rc *RoomContainer) Container(params map[string]string) fyne.CanvasObject {
	fmt.Println(params)
	ctn := container.NewVScroll(
		container.NewVBox(
			widget.NewLabel("Room"),
		),
	)

	return ctn
}
