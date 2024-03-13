package controllers

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rm4n0s/go-observer"
	"github.com/rm4n0s/wt-video-audio-transmission-example/app/connections"
)

type RoomController struct {
	inputConn  *connections.InputConnection
	outputConn *connections.OutputConnection
	Closed     observer.Property[bool]
}

func NewRoomController() *RoomController {
	ic := connections.NewInputConnection()
	oc := connections.NewOutputConnection()
	return &RoomController{
		inputConn:  ic,
		outputConn: oc,
		Closed:     observer.NewProperty(false),
	}
}

func (rc *RoomController) InitConns(host, username string) error {
	log.Println("InitConns")
	err := rc.inputConn.StartConnection(host, username)
	if err != nil {
		return err
	}
	err = rc.outputConn.StartConnection(host, username)
	if err != nil {
		return err
	}

	go func(ic *connections.InputConnection) {
		ic.SendData()
	}(rc.inputConn)

	go func(oc *connections.OutputConnection) {
		oc.ReceiveData()
	}(rc.outputConn)
	return nil
}

func (rc *RoomController) Container(params map[string]string) fyne.CanvasObject {
	ctn := container.NewVScroll(
		container.NewVBox(
			widget.NewLabel("Room"),
			widget.NewButton("Close", func() {
				rc.inputConn.CloseConnection()
				rc.Closed.Update(true)
			}),
		),
	)

	return ctn
}
