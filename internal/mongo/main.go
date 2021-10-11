package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Options are the options to configure the mongo.client
type Options struct {
	// URI to connect
	URI string
}

func NewClient(opts *Options) (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(opts.URI))
}
