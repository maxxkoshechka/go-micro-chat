package chat

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"golang.org/x/net/websocket"
)

func main()  {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewHandler(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err = ws.Handshake(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} )
}

func NewHandler(w http.ResponseWriter, r *http.Request) (*WS, error) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (ws *WS) Handshake() error {
	hash := func(key string) string {
		h := sha1.New()
		h.Write([]byte(key))
		h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
		return base64.StdEncoding.EncodeToString(h.Sum(nil))
	}(ws.header.Get("Sec-WebSocket-Key"))
}
