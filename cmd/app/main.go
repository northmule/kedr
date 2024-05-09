package main

import (
	"kedr/internal/config"
	"kedr/internal/runner"
	"log"
	"os"
	"path/filepath"
)

func main() {

	log.Println("Start runner")
	runner.StartApp()
}

func init() {
	file, err := os.OpenFile(filepath.Join(config.DirLog, "app.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}
