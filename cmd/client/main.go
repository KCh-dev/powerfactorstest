package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

func main() {
	var count int
	flag.IntVar(&count, "n", 1, "Number of parallel connections")
	flag.Parse()

	uri := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/goapp/ws"}
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(id int) {
			defer wg.Done()

			// Setup custom Dialer with Origin header
			dialer := *websocket.DefaultDialer
			headers := make(http.Header)
			headers.Add("Origin", "http://localhost:8080")

			c, _, err := dialer.Dial(uri.String(), headers)
			if err != nil {
				log.Printf("Client %d error: %v", id, err)
				return
			}
			defer func(c *websocket.Conn) {
				err := c.Close()
				if err != nil {

				}
			}(c)
			log.Printf("Client %d connected", id)

			err = c.WriteMessage(websocket.TextMessage, []byte("Hello from client "+strconv.Itoa(id)))
			if err != nil {
				log.Printf("write error: %v", err)
				return
			}

			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Printf("read error: %v", err)
					break
				}
				log.Printf("Client %d received: %s", id, message)
			}
		}(i)
	}

	wg.Wait()
	log.Println("All clients have finished their operations")
}
