package http

import (
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Global variables for this package
var (
	appName = "app name"
	mng     = &mongo.Client{}
)

// ServerOption are the options to configure the http.Server from the standard
// library
type ServerOption func(*http.Server)

// NewServer is a wrapper over http.Server to configure it with a Function
// Pattern as in:
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
func NewServer(options ...ServerOption) *http.Server {
	const (
		defaultReadTimeout       = 1 * time.Second
		defaultWriteTimeout      = 1 * time.Second
		defaultIdleTimeout       = 30 * time.Second
		defaultReadHeaderTimeout = 2 * time.Second
	)

	srv := &http.Server{
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
		IdleTimeout:       defaultIdleTimeout,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		Handler:           Handlers(),
	}

	for _, option := range options {
		option(srv)
	}

	return srv
}

// WithReadTimeout sets the ReadTimeout
func WithReadTimeout(seconds time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.ReadTimeout = seconds * time.Second
	}
}

// WithWriteTimeout sets the WriteTimeout
func WithWriteTimeout(seconds time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.WriteTimeout = seconds * time.Second
	}
}

// WithIdleTimeout sets the IdleTimeout
func WithIdleTimeout(seconds time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.IdleTimeout = seconds * time.Second
	}
}

// WithReadHeaderTimeout sets the ReadHeaderTimeout
func WithReadHeaderTimeout(seconds time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.ReadHeaderTimeout = seconds * time.Second
	}
}

// WithPort sets the port of the http server
func WithPort(port string) ServerOption {
	return func(srv *http.Server) {
		srv.Addr = fmt.Sprintf(":%s", port)
	}
}

// WithAppName sets the app name
func WithAppName(name string) ServerOption {
	return func(srv *http.Server) {
		appName = name
	}
}

// WithMongo sets the mongo client
func WithMongo(m *mongo.Client) ServerOption {
	return func(srv *http.Server) {
		mng = m
	}
}
