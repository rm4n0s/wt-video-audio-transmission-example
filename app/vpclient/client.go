package vpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/quic-go/webtransport-go"
)

type VoipClient struct {
	d         webtransport.Dialer
	connected bool
	session   *webtransport.Session
}

func NewVoipClient() *VoipClient {
	return &VoipClient{}
}

func (c *VoipClient) Connect(ctx context.Context, url string) error {
	rsp, session, err := c.d.Dial(ctx, url, nil)
	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(rsp.Body)
		return fmt.Errorf("failure response: '%s' with status code %d", string(b), rsp.StatusCode)
	}
	c.connected = true
	c.session = session
	return nil
}

func (c *VoipClient) Connected() bool {
	return c.connected
}
