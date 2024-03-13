package main

import (
	"context"
	"image"
	"image/draw"

	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec/x264"
	_ "github.com/pion/mediadevices/pkg/driver/camera"
	"github.com/pion/mediadevices/pkg/prop"
)

var camFrame = image.NewRGBA(image.Rect(0, 0, 640, 480))

func startWebcam(ctx context.Context, readyFrame chan struct{}) error {
	x264Params, _ := x264.NewParams()
	x264Params.Preset = x264.PresetUltrafast
	x264Params.BitRate = 500_000
	x264Params.KeyFrameInterval = 30

	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&x264Params),
	)

	mediaStream, err := mediadevices.GetUserMedia(mediadevices.MediaStreamConstraints{
		Video: func(c *mediadevices.MediaTrackConstraints) {
			c.Width = prop.Int(640)
			c.Height = prop.Int(480)
			c.DeviceID = prop.String("0")
		},
		Codec: codecSelector,
	})
	if err != nil {
		return err
	}
	videoTrack := mediaStream.GetVideoTracks()[0].(*mediadevices.VideoTrack)
	defer videoTrack.Close()

	videoReader := videoTrack.NewReader(true)
	for {
		select {
		case <-ctx.Done():
		default:
			frame, release, err := videoReader.Read()
			if err != nil {
				return err
			}
			draw.Draw(camFrame, camFrame.Bounds(), frame, image.Point{0, 0}, draw.Over)
			release()
			readyFrame <- struct{}{}
		}
	}
}
