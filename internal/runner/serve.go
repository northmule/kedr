package runner

import (
	"github.com/gofor-little/env"
	"html/template"
	"kedr/internal/config"
	"kedr/internal/services/telegram"
	"kedr/internal/utils"
	"net/http"
)

// StartApp Запуск приложения
func StartApp() {
	loadRouteHttp()
	err := http.ListenAndServe(config.GetHostAddress(), nil)
	utils.CheckError(err)
}

func InitApp() {
	loadEnvironmentVariables()
}

// Все доступные роуты
func loadRouteHttp() {
	http.HandleFunc("/", indexHandler)
	// Api
	http.HandleFunc("/api/v1/telegram/update", telegram.Update)
}

// Стартовая страница
func indexHandler(w http.ResponseWriter, r *http.Request) {
	path := utils.GetTemplatePath("index.html")
	html, err := template.ParseFiles(path)
	utils.CheckError(err)
	err = html.Execute(w, nil)
}

func loadEnvironmentVariables() {
	if err := env.Load(config.FileEnv); err != nil {
		panic(err)
	}
}
