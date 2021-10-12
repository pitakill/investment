package main

import (
	"errors"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/pitakill/investment/internal"
)

const (
	// root directory name for loading the .env file
	root = "investment"
	// exitFail common exit fail code
	exitFail = 1
)

var (
	// Errors known
	errEnv = errors.New("error loading .env file")
)

func main() {
	if err := loadEnv(root); err != nil {
		// log.Fatalln(err)
	}

	options := &internal.AppOptions{
		AppName:  root,
		HttpPort: os.Getenv("HTTP_PORT"),
		MongoURI: os.Getenv("MONGO_URI"),
	}

	app := internal.NewApp(options)

	if err := app.Start(); err != nil {
		log.Fatalln(err)
	}
}

// loadEnv loads the .env file from the root directory
func loadEnv(projectName string) error {
	re := regexp.MustCompile(`^(.*` + projectName + `)`)
	directory, _ := os.Getwd()
	path := re.Find([]byte(directory))

	if err := godotenv.Load(string(path) + `/.env`); err != nil {
		return errEnv
	}

	return nil
}
