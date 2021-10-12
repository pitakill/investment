package internal

import (
	"context"
	originalHttp "net/http"

	originalMongo "go.mongodb.org/mongo-driver/mongo"

	"github.com/pitakill/investment/internal/http"
	"github.com/pitakill/investment/internal/mongo"
)

// AppOptions configures the App
type AppOptions struct {
	// AppName is the name of the binary and also the root path
	AppName string
	// Port is the listening port of the http server
	Port string
	// MongoURI is the listening URI of the mongo server
	MongoURI string
}

// App struct wraps all the dependencies of the Application
type App struct {
	// srv reference to the http server from the original standard library
	srv *originalHttp.Server
	// srv reference to the http server from the original standard library
	mongo *originalMongo.Client
}

// NewApp returns an App wint the AppOptions
func NewApp(options *AppOptions) *App {
	mngOptions := &mongo.Options{
		URI: options.MongoURI,
	}
	mng, _ := mongo.NewClient(mngOptions)

	srv := http.NewServer(
		http.WithAppName(options.AppName),
		http.WithIdleTimeout(20),
		http.WithPort(options.Port),
		http.WithMongo(mng),
	)

	return &App{
		srv: srv,
		// mongo: mng,
	}
}

// Start initialize the http server and the mongo client
func (app *App) Start() error {
	defer func() {
		app.mongo.Disconnect(context.TODO())
	}()

	return app.srv.ListenAndServe()
}
