package sse

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ServerSendEvents(c *gin.Context) {
	c.SSEvent("test", fmt.Sprintf("data: %v\n", time.Now().Format(time.RFC3339)))
}

func sseHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}
	for {
		fmt.Fprintf(w, "data: %v\n", time.Now().Format(time.RFC3339))
		flusher.Flush()
		time.Sleep(time.Second)
	}
}
