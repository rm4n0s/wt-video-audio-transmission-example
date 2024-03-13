package connections

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/quic-go/webtransport-go"
)

type InputConnection struct {
	session *webtransport.Session
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewInputConnection() *InputConnection {
	return &InputConnection{}
}

func (ic *InputConnection) StartConnection(host, username string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ic.ctx = ctx
	ic.cancel = cancel
	var dialer webtransport.Dialer

	url := fmt.Sprintf("%s/input?username=%s", host, username)
	rsp, session, err := dialer.Dial(ic.ctx, url, nil)
	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(rsp.Body)
		return fmt.Errorf("failure response: '%s' with status code %d", string(b), rsp.StatusCode)
	}

	ic.session = session
	return nil
}

func (ic *InputConnection) CloseConnection() {
	ic.cancel()
	log.Println("closedinput connection")
}

func (ic *InputConnection) SendData() {
	log.Println("SendData")
	str, err := ic.session.OpenUniStream()
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		select {
		case <-ic.ctx.Done():
			str.Close()
		default:
			s := uuid.NewString()
			str.Write([]byte("hello world " + s + " " + t.String()))
		}
	}
}
