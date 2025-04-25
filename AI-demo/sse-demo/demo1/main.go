package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// SSEHandler handles the SSE connection
func SSEHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a channel for client disconnect detection
	clientClosed := r.Context().Done()

	// Create a flusher to ensure messages are sent immediately
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Counter for demonstration
	counter := 0

	// Sample messages to stream
	messages := []string{
		"这是第一条消息",
		"This is the second message",
		"这是第三条消息",
		"Processing complete!",
	}

	// Simulate streaming text messages
	for _, msg := range messages {
		select {
		case <-clientClosed:
			// Client disconnected
			log.Println("Client closed connection")
			return
		default:
			// Format the SSE message
			counter++
			eventData := fmt.Sprintf("id: %d\nevent: message\ndata: %s\n\n", counter, msg)

			// Write to response
			fmt.Fprint(w, eventData)

			// Flush the data to the client
			flusher.Flush()

			// Simulate processing time
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Define the SSE endpoint
	http.HandleFunc("/stream", SSEHandler)

	// Home page with a simple demo
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
<!DOCTYPE html>
<html>
<head>
    <title>SSE Demo</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        #output { border: 1px solid #ccc; padding: 10px; min-height: 200px; margin-top: 20px; }
        .demo-links { margin-bottom: 20px; }
        .demo-links a { margin-right: 15px; text-decoration: none; color: #0066cc; }
        .demo-links a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>Go SSE 演示</h1>
    
    <div class="demo-links">
        <a href="/">基础演示</a>
        <a href="/advanced">高级演示</a>
    </div>
    
    <h2>基础 SSE 演示</h2>
    <button onclick="startStream()">开始流式输出</button>
    <div id="output"></div>

    <script>
        function startStream() {
            const outputDiv = document.getElementById('output');
            outputDiv.innerHTML = '';
            
            // Create SSE connection
            const eventSource = new EventSource('/stream');
            
            eventSource.addEventListener('message', function(e) {
                const newElement = document.createElement('div');
                newElement.textContent = e.data;
                outputDiv.appendChild(newElement);
            });
            
            eventSource.onerror = function(e) {
                console.log('SSE error', e);
                eventSource.close();
            };
            
            // Close connection when all messages are received
            eventSource.addEventListener('done', function(e) {
                eventSource.close();
            });
        }
    </script>
</body>
</html>
`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	// Start the advanced server components
	StartAdvancedServer()

	// Start the server
	port := ":8080"
	fmt.Printf("SSE server started on http://localhost%s\n", port)
	fmt.Printf("Basic demo: http://localhost%s\n", port)
	fmt.Printf("Advanced demo: http://localhost%s/advanced\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
