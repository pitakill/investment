package main

import (
	"errors"
	"testing"
)

func Test_loadEnv(t *testing.T) {
	tests := []struct {
		path string
		err  error
	}{
		{root, nil},
		{"notADirectory", errEnv},
	}

	for _, test := range tests {
		err := loadEnv(test.path)
		if !errors.Is(err, test.err) {
			t.Error(err)
		}
	}

}
