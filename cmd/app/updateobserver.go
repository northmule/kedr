package main

import (
	"encoding/json"
	"fmt"
	"kedr/internal/config"
	"kedr/internal/models/telegram"
	"kedr/internal/runner"
	telegramService "kedr/internal/services/telegram"
	"kedr/internal/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Response Ответ от сервиса телеграм на запрос getUpdates
type Response struct {
	Ok     bool                      `json:"ok"`
	Result []telegram.ReceiveMessage `json:"result"`
}

type responseLocal struct {
	Ok bool `json:"ok"`
}

func main_() {
	for {
		if runObserve() {
			return
		}
		time.Sleep(1 * time.Second)
	}

}

// ObserverStart Запуск получения и отправки данных для локального сервера
func ObserverStart() {
	for {
		runObserve()
		time.Sleep(1 * time.Second)
	}
}

func runObserve() bool {

	runner.InitApp()
	url := telegramService.GetUrlAction("getUpdates")
	response := requestGetTelegram(url)

	if !response.Ok {
		log.Printf("Запрос не выполнен: %#v", response)
		return false
	}
	fileList, err := os.ReadDir(config.DirObserverData)
	utils.CheckError(err)

	for _, receiveMessage := range response.Result {
		fileName := strconv.Itoa(int(receiveMessage.UpdateID))
		file, err := os.OpenFile(filepath.Join(config.DirObserverData, fileName), os.O_CREATE|os.O_WRONLY, 0666)
		utils.CheckError(err)
		defer file.Close()
		isExist := false
		for _, file := range fileList {
			if file.Name() == fileName {
				isExist = true
				break
			}
		}
		if !isExist {
			jsonString, _ := json.Marshal(receiveMessage)
			_, err := file.Write(jsonString)
			utils.CheckError(err)
			requestPostLocalApp(receiveMessage)
		}
	}
	return true
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

func requestPostLocalApp(message telegram.ReceiveMessage) {
	jsonString, _ := json.Marshal(message)
	post, err := http.Post(
		fmt.Sprintf("http://%s/api/v1/telegram/update", config.GetHostAddress()),
		"application/json",
		strings.NewReader(string(jsonString)),
	)
	utils.CheckError(err)
	response := &responseLocal{}
	err = json.NewDecoder(post.Body).Decode(&response)
	utils.CheckError(err)
	log.Printf("Отправил данные: %#v. Ответ сервера: %#v", string(jsonString), response)
}
