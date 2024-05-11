package main

import (
	"encoding/json"
	"kedr/internal/config"
	telegram2 "kedr/internal/models/telegram"
	"kedr/internal/runner"
	"kedr/internal/services/telegram"
	"kedr/internal/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Response Ответ от сервиса телеграм на запрос getUpdates
type Response struct {
	Ok     bool                       `json:"ok"`
	Result []telegram2.ReceiveMessage `json:"result"`
}

func main() {

	runner.InitApp()
	url := telegram.GetUrlAction("getUpdates")
	response := requestGetTelegram(url)

	if !response.Ok {
		log.Printf("Запрос не выполнен: %#v", response)
	}
	fileList, err := os.ReadDir(config.DirObserverData)
	utils.CheckError(err)

	for _, message := range response.Result {
		fileName := strconv.Itoa(int(message.UpdateID))
		file, err := os.OpenFile(filepath.Join(config.DirObserverData, fileName), os.O_CREATE|os.O_WRONLY, 0666)
		utils.CheckError(err)
		defer file.Close()
	}
}

func requestGetTelegram(url string) *Response {
	get, err := http.Get(url)
	defer get.Body.Close()
	utils.CheckError(err)
	response := &Response{}
	err = json.NewDecoder(get.Body).Decode(&response)
	utils.CheckError(err)
	return response
}

func init() {
	file, err := os.OpenFile(filepath.Join(config.DirLog, "update-observer.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}
