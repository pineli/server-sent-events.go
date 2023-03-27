package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/sse-server", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		for {
			currentTime := time.Now()
			fmt.Fprintf(w, "event: server\n")
			fmt.Fprintf(w, "data: %s\n\n", currentTime.Format("2006-01-02 15:04:05"))
			w.(http.Flusher).Flush()
			time.Sleep(1000 * time.Millisecond)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "page/index.html")
	})
	fmt.Println("Servidor executando na porta 8080")

	http.ListenAndServe(":8080", nil)
}
