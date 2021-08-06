package build

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type MessageSSE struct {
	Id   int64
	Name string
	Data string
}

func NewMessageSSE(name, data string) *MessageSSE {
	return &MessageSSE{
		Id:   time.Now().Unix(),
		Name: name,
		Data: data,
	}
}

var sseChan chan *MessageSSE

func sseHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Get __sse from client")

	// prepare the header
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// instantiate the channel
	sseChan = make(chan *MessageSSE)

	// close the channel after exit the function
	defer func() {
		close(sseChan)
		sseChan = nil
		log.Printf("Client connection is closed")
	}()

	// prepare the flusher
	flusher, _ := w.(http.Flusher)

	// trap the request under loop forever
	for {

		select {

		// message will received here and printed
		case message := <-sseChan:
			fmt.Fprintf(w, "id: %d\n", message.Id)
			fmt.Fprintf(w, "event: %s\n", message.Name)
			fmt.Fprintf(w, "data: %s\n\n", message.Data)
			flusher.Flush()

		// connection is closed then defer will be executed
		case <-r.Context().Done():
			return

		}
	}
}

func doReload() {
	if sseChan != nil {
		log.Printf("Do client reload")

		// send the message through the available channel
		sseChan <- NewMessageSSE("reload", "Do Reload")
	}
}
