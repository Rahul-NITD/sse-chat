package sse_chat

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

type ChatHandler struct {
	http.Handler
}

func NewChatHandler() *ChatHandler {
	r := http.NewServeMux()

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {})

	sse_server := sse.New()
	sse_server.OnSubscribe = func(streamID string, sub *sse.Subscriber) {
		sse_server.Publish(streamID, &sse.Event{
			Data: []byte("Hello User"),
		})
	}

	sse_server.AutoStream = true

	r.HandleFunc("/room", sse_server.ServeHTTP)

	hdlr := &ChatHandler{}
	hdlr.Handler = r
	return hdlr
}
