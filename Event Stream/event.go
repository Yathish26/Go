package main

import (
	"fmt"
	"net/http"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	messages := []string{"Hello, world!", "Streaming data...", "This is an event stream", "Goodbye!"}

	for _, message := range messages {
		fmt.Fprintf(w, "data: %s\n\n", message)
		w.(http.Flusher).Flush()
		time.Sleep(2 * time.Second)
	}
}

func main() {
	http.HandleFunc("/stream", streamHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
<!DOCTYPE html>
<html>
<head>
    <title>Event Stream</title>
</head>
<body>
    <h1>Event Stream</h1>
    <div id="output"></div>
    <script>
        const output = document.getElementById('output');
        const eventSource = new EventSource('/stream');

        eventSource.onmessage = function(event) {
            const p = document.createElement('p');
            p.textContent = event.data;
            output.appendChild(p);
        };
    </script>
</body>
</html>`
		w.Write([]byte(html))
	})

	http.ListenAndServe(":8080", nil)
}
