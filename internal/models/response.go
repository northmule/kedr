package models

type ResponseStatus struct {
	Status string `json:"status"`
}

// GetResponseOk Положительный ответ на запрос
func GetResponseOk() ResponseStatus {
	return ResponseStatus{Status: "ok"}
}
