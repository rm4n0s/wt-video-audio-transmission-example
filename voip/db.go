package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/quic-go/webtransport-go"
)

type InputStream struct {
	ID       string
	Username string
	Session  *webtransport.Session
	Stream   webtransport.ReceiveStream
}

type OutputStream struct {
	ID       string
	Username string
	Session  *webtransport.Session
	Stream   webtransport.SendStream
}

type UserDB struct {
	inputs  sync.Map
	outputs sync.Map
}

func NewUserDB() *UserDB {
	return &UserDB{
		inputs:  sync.Map{},
		outputs: sync.Map{},
	}
}

func (udb *UserDB) StartTransmission() {
	readBytes := make([]byte, 800)
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		udb.inputs.Range(func(key, value any) bool {
			inputUsername := key.(string)
			inputStream := value.(*InputStream)
			inputStream.Stream.Read(readBytes)
			msg := string(readBytes)
			udb.outputs.Range(func(key, value any) bool {
				outputUsername := key.(string)
				outputStream := value.(*OutputStream)
				l := fmt.Sprintf("from: %s to: %s msg: %s", inputUsername, outputUsername, msg)
				log.Println(l)

				outputStream.Stream.Write([]byte(l))
				return true
			})

			return true
		})
	}
}

func (udb *UserDB) submitInputStream(username, id string, session *webtransport.Session, receiveStream webtransport.ReceiveStream) {
	udb.inputs.Store(username, &InputStream{
		Username: username,
		ID:       id,
		Session:  session,
		Stream:   receiveStream,
	})
}

func (udb *UserDB) submitOutputStream(username, id string, session *webtransport.Session, sendStream webtransport.SendStream) {
	udb.outputs.Store(username, &OutputStream{
		Username: username,
		ID:       id,
		Session:  session,
		Stream:   sendStream,
	})
}
