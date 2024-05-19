package httpsrv

import (
	"goapp/internal/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

// Ensure your server setup initializes the configuration appropriately
func setupTestServer() *Server {
	// Mock configuration with allowed origins
	cfg := &config.Config{
		AllowedOrigins: []string{"http://allowedorigin.com"},
	}

	// Initialize your server with the test configuration
	return New(nil, cfg) // Assuming `New` takes a config. Adjust if your actual signature differs.
}

func TestWebSocketOriginValidation(t *testing.T) {
	server := setupTestServer()
	srv := httptest.NewServer(http.HandlerFunc(server.handlerWebSocket))
	defer srv.Close()

	// Convert http server URL to websocket URL
	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http")

	// Define a custom dialer with WebSocket options if needed
	dialer := websocket.Dialer{}

	// Test Allowed Origin
	header := http.Header{}
	header.Add("Origin", "http://allowedorigin.com")
	_, resp, err := dialer.Dial(wsURL, header)
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
	err = resp.Body.Close()
	if err != nil {
		return
	}

	// Test Disallowed Origin
	header.Set("Origin", "http://disallowedorigin.com")
	_, resp, err = dialer.Dial(wsURL, header)
	if err == nil {
		t.Errorf("Connection was supposed to be denied for disallowed origin")
	}
	if resp != nil {
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}
}
