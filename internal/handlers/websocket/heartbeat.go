package ws

import (
	"context"
	"fmt"
	"log"

	"github.com/coder/websocket"
)

// Sends a heartbeat to Discord's WSS
func SendHeartbeat(lastSequenceNumber any) {
	var heartbeat string

	if lastSequenceNumber != nil {
		heartbeat = fmt.Sprintf("{\"op\": 1, \"d\": %v}", lastSequenceNumber)
	} else {
		heartbeat = "{\"op\": 1, \"d\": null}"
	}

	err := Connection.Write(context.Background(), websocket.MessageText, []byte(heartbeat))
	if err != nil {
		log.Fatal(err)
	}
}