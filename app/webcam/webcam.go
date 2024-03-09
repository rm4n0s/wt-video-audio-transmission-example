package webcam

import (
	"context"
	"image"
	"image/draw"

	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec/x264"
	_ "github.com/pion/mediadevices/pkg/driver/camera"
	"github.com/pion/mediadevices/pkg/prop"
)

type Webcam struct {
	codecSelector *mediadevices.CodecSelector
	rgba          *image.RGBA
	width         int
	height        int
}

type Picture struct {
	Pix    []byte
	Width  int
	Height int
	Stride int
}

func NewWebcam(width, height int) (*Webcam, error) {
	x264Params, _ := x264.NewParams()
	x264Params.Preset = x264.PresetMedium
	x264Params.BitRate = 1_000_000 // 1mbps

	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&x264Params),
	)

	return &Webcam{
		codecSelector: codecSelector,
		rgba:          image.NewRGBA(image.Rect(0, 0, width, height)),
		width:         width,
		height:        height,
	}, nil
}

func (wc *Webcam) Read(ctx context.Context, deviceID string, picChan chan Picture) error {
	mediaStream, err := mediadevices.GetUserMedia(mediadevices.MediaStreamConstraints{
		Video: func(c *mediadevices.MediaTrackConstraints) {
			c.Width = prop.Int(wc.width)
			c.Height = prop.Int(wc.height)
			c.DeviceID = prop.String(deviceID)
		},
		Codec: wc.codecSelector,
	})
	if err != nil {
		return err
	}
	videoTrack := mediaStream.GetVideoTracks()[0].(*mediadevices.VideoTrack)
	videoReader := videoTrack.NewReader(false)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			frame, release, err := videoReader.Read()
			if err != nil {
				return err
			}
			defer release()
			draw.Draw(wc.rgba, wc.rgba.Bounds(), frame, image.Point{}, draw.Src)
			pic := Picture{
				Pix:    []byte{},
				Width:  wc.width,
				Height: wc.height,
				Stride: wc.rgba.Stride,
			}
			copy(pic.Pix, wc.rgba.Pix)
			picChan <- pic
		}
	}
}
