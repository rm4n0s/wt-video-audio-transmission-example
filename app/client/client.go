package client

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/quic-go/webtransport-go"
)

type Client struct {
	d         webtransport.Dialer
	connected bool
	session   *webtransport.Session
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Connect(ctx context.Context, link string) error {
	rsp, session, err := c.d.Dial(ctx, link, nil)
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

func (c *Client) Connected() bool {
	return c.connected
}
