package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// TextProcessor simulates a text processing function that generates results over time
type TextProcessor struct {
	Results chan string
	Steps   chan string
	Done    chan bool
}

// Process simulates processing text and sending results through the channel
func (tp *TextProcessor) Process(text string) {
	go func() {
		defer close(tp.Steps)
		defer close(tp.Results)
		defer close(tp.Done)

		// Simulate text analysis or generation steps
		steps := []struct {
			eventType string
			message   string
			delay     time.Duration
		}{
			{"step", "开始处理文本...", 500 * time.Millisecond},
			{"step", "分析中...", 800 * time.Millisecond},
			{"message", fmt.Sprintf("处理文本: %s", text), 1 * time.Second},
			{"step", "生成结果中...", 700 * time.Millisecond},
			{"message", "第一部分结果: 这是对文本的分析", 1 * time.Second},
			{"message", "第二部分结果: 这是一些附加信息", 900 * time.Millisecond},
			{"step", "处理完成!", 500 * time.Millisecond},
		}

		for _, step := range steps {

			if step.eventType == "step" {
				time.Sleep(step.delay)
				tp.Steps <- step.message
			}

			if step.eventType == "message" {
				time.Sleep(step.delay)
				tp.Results <- step.message
			}

		}

		tp.Done <- true
	}()
}

// NewTextProcessor creates a new text processor with channels
func NewTextProcessor() *TextProcessor {
	return &TextProcessor{
		Steps:   make(chan string),
		Results: make(chan string),
		Done:    make(chan bool),
	}
}

// AdvancedSSEHandler handles SSE with real-time text processing
func AdvancedSSEHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get flusher
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Get input text from query parameter (default if not provided)
	inputText := r.URL.Query().Get("text")
	if inputText == "" {
		inputText = "默认示例文本"
	}

	// Create text processor
	processor := NewTextProcessor()
	processor.Process(inputText)

	// Channel for client disconnect
	clientClosed := r.Context().Done()

	// Event counter
	counter := 0

	// Stream SSE messages until done or client disconnects
	for {
		select {
		case <-clientClosed:
			log.Println("Client closed connection")
			return

		case result, ok := <-processor.Results:
			if !ok {
				// Channel closed
				return
			}

			// Send SSE event
			counter++
			eventData := fmt.Sprintf("id: %d\nevent: message\ndata: %s\n\n", counter, result)
			fmt.Fprint(w, eventData)
			flusher.Flush()

		case result, ok := <-processor.Steps:
			if !ok {
				// Channel closed
				return
			}

			// Send SSE event
			counter++
			eventData := fmt.Sprintf("id: %d\nevent: step\ndata: %s\n\n", counter, result)
			fmt.Fprint(w, eventData)
			flusher.Flush()

		case <-processor.Done:
			// Final event
			counter++
			eventData := fmt.Sprintf("id: %d\nevent: done\ndata: Processing completed\n\n", counter)
			fmt.Fprint(w, eventData)
			flusher.Flush()
			return
		}
	}
}

func StartAdvancedServer() {
	// Register the advanced SSE handler
	http.HandleFunc("/advanced-stream", AdvancedSSEHandler)

	// Add a demo page
	http.HandleFunc("/advanced", func(w http.ResponseWriter, r *http.Request) {
		html := `
<!DOCTYPE html>
<html>
<head>
    <title>Advanced SSE Demo</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        #output { border: 1px solid #ccc; padding: 10px; min-height: 200px; margin-top: 20px; }
        input { padding: 5px; width: 300px; }
        button { padding: 5px 10px; }
    </style>
</head>
<body>
    <h1>Advanced SSE Text Processing Demo</h1>
    <div>
        <input type="text" id="inputText" placeholder="输入要处理的文本" value="测试文本">
        <button onclick="startStream()">开始处理</button>
    </div>
    <div id="output"></div>

    <script>
        function startStream() {
            const outputDiv = document.getElementById('output');
            const inputText = document.getElementById('inputText').value || '默认文本';
            
            outputDiv.innerHTML = '';
            
            // Create SSE connection with query parameter
            const eventSource = new EventSource('/advanced-stream?text=' + encodeURIComponent(inputText));
            
            eventSource.addEventListener('message', function(e) {
                const newElement = document.createElement('div');
                newElement.textContent = e.data;
                outputDiv.appendChild(newElement);
            });
            
            eventSource.addEventListener('done', function(e) {
                const newElement = document.createElement('div');
                newElement.textContent = '✅ ' + e.data;
                newElement.style.fontWeight = 'bold';
                outputDiv.appendChild(newElement);
                
                // Close the connection
                eventSource.close();
            });
            
            eventSource.onerror = function(e) {
                console.log('SSE error', e);
                eventSource.close();
            };
        }
    </script>
</body>
</html>
`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	fmt.Println("Advanced SSE demo available at http://localhost:8080/advanced")
}
