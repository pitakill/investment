package http

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	// Default Server
	want := &http.Server{
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	output := NewServer()

	// Testing the Handlers is another Test
	output.Handler = nil

	if !reflect.DeepEqual(output, want) {
		t.Errorf("want %v, got %v", want, output)
	}

	// Set Port
	want = &http.Server{
		Addr: ":8080",
	}

	output = NewServer(WithPort("8080"))
	if output.Addr != want.Addr {
		t.Errorf("want %v, got %v", want, output)
	}

	// Set appName
	wanted := "test"

	output = NewServer(WithAppName("test"))
	if appName != wanted {
		t.Errorf("want %v, got %v", wanted, output)
	}

	// Set ReadTimeout
	want = &http.Server{
		ReadTimeout: 20 * time.Second,
	}

	output = NewServer(WithReadTimeout(20))
	if !reflect.DeepEqual(output.ReadTimeout, want.ReadTimeout) {
		t.Errorf("want %v, got %v", wanted, output)
	}

	// Set WriteTimeout
	want = &http.Server{
		WriteTimeout: 20 * time.Second,
	}

	output = NewServer(WithWriteTimeout(20))
	if !reflect.DeepEqual(output.WriteTimeout, want.WriteTimeout) {
		t.Errorf("want %v, got %v", wanted, output)
	}

	// Set IdleTimeout
	want = &http.Server{
		IdleTimeout: 20 * time.Second,
	}

	output = NewServer(WithIdleTimeout(20))
	if !reflect.DeepEqual(output.IdleTimeout, want.IdleTimeout) {
		t.Errorf("want %v, got %v", wanted, output)
	}

	// Set ReadHeaderTimeout
	want = &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
	}

	output = NewServer(WithReadHeaderTimeout(20))
	if !reflect.DeepEqual(output.ReadHeaderTimeout, want.ReadHeaderTimeout) {
		t.Errorf("want %v, got %v", wanted, output)
	}
}
