package internal

import (
	"net/http"
	"testing"
)

func TestNewApp(t *testing.T) {
	options := &AppOptions{
		Port: "80",
	}

	want := &App{
		srv: &http.Server{
			Addr: ":80",
		},
	}

	output := NewApp(options)

	if output.srv.Addr != want.srv.Addr {
		t.Errorf("want %q, got %q", want.srv.Addr, output.srv.Addr)
	}
}
