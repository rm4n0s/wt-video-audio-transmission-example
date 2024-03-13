package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/webtransport-go"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var server *webtransport.Server
	userDB := NewUserDB()
	go userDB.StartTransmission()
	http.HandleFunc("/input", func(w http.ResponseWriter, r *http.Request) {
		session, err := server.Upgrade(w, r)
		if err != nil {
			log.Printf("Error %s: %s", r.URL.Path, err)
			io.WriteString(w, "Error:"+err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		username := r.URL.Query().Get("username")
		if len(username) == 0 {
			io.WriteString(w, "Error: username is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		inputSessionID := fmt.Sprintf("input-%s-%s", username, uuid.New().String())
		stream, err := session.AcceptUniStream(session.Context())
		if err != nil {
			io.WriteString(w, "Error: failed to AcceptUniStream for "+inputSessionID+":"+err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		userDB.submitInputStream(username, inputSessionID, session, stream)

		log.Printf("New input session %s", inputSessionID)
		log.Println(session)
	})

	http.HandleFunc("/output", func(w http.ResponseWriter, r *http.Request) {
		session, err := server.Upgrade(w, r)
		if err != nil {
			log.Printf("Error %s: %s", r.URL.Path, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		username := r.URL.Query().Get("username")
		if len(username) == 0 {
			io.WriteString(w, "Error: username is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		outputSessionID := fmt.Sprintf("output-%s-%s", username, uuid.New().String())
		stream, err := session.OpenUniStream()
		if err != nil {
			io.WriteString(w, "Error: failed to OpenUniStream for "+outputSessionID+":"+err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		userDB.submitOutputStream(username, outputSessionID, session, stream)

		log.Printf("New output session %s", outputSessionID)
		log.Println(session)
	})

	// Note: "new-tab-page" in AllowedOrigins lets you access the server from a blank tab (via DevTools Console).
	// "" in AllowedOrigins lets you access the server from JavaScript loaded from disk (i.e. via a file:// URL)
	server = &webtransport.Server{
		H3: http3.Server{Addr: ":4433", EnableDatagrams: true},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	log.Println("Launching WebTransport server at:", server.H3.Addr)
	err := server.ListenAndServeTLS("certs/certificate.pem", "certs/certificate.key")
	if err != nil {
		log.Fatal("Server error:", err)
	}

}
