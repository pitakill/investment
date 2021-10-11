package mongo

import (
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewClient(t *testing.T) {
	opts := &Options{URI: ""}

	output, _ := NewClient(opts)

	want, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(""))

	if !reflect.DeepEqual(output, want) {
		t.Errorf("want %v, got %v", want, output)
	}
}
