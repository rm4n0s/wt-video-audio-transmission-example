package connections

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/quic-go/webtransport-go"
)

type OutputConnection struct {
	session *webtransport.Session
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewOutputConnection() *OutputConnection {
	return &OutputConnection{}
}

func (oc *OutputConnection) StartConnection(host, username string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	oc.ctx = ctx
	oc.cancel = cancel
	var dialer webtransport.Dialer

	url := fmt.Sprintf("%s/output?username=%s", host, username)
	rsp, session, err := dialer.Dial(oc.ctx, url, nil)
	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(rsp.Body)
		return fmt.Errorf("failure response: '%s' with status code %d", string(b), rsp.StatusCode)
	}

	oc.session = session
	return nil
}

func (oc *OutputConnection) Close() {
	oc.cancel()
	log.Println("closed output connection")
}

func (oc *OutputConnection) ReceiveData() {
	log.Println("ReceiveData")
	str, err := oc.session.AcceptUniStream(oc.ctx)
	if err != nil {
		log.Fatal(err)
	}
	readBytes := make([]byte, 31800)
	for {
		str.Read(readBytes)
		log.Println(string(readBytes))
		//clear(readBytes)
	}
}
