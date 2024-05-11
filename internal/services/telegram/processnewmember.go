package telegram

import (
	"fmt"
	"kedr/internal/config"
	"kedr/internal/models/telegram"
	"log"
	"strconv"
	"strings"
)

// ProcessNewMember действия с новым участником
func ProcessNewMember(message *telegram.ReceiveMessage) bool {
	// Ограничивание пользователя в правах
	result := RestrictChatMember(GetChatId(message), message.Message.NewChatMember.Id)
	if !result.Ok {
		log.Printf(
			"Ограничения не назначены, memberId: %d chatId: %d",
			message.Message.NewChatMember.Id,
			GetChatId(message),
		)
		return false
	}
	// Отправка сообщения с предложением ответить на него решением
	newMessage := telegram.NewMessage{
		ChatId:              GetChatId(message),
		Text:                fmt.Sprintf("%s %s %s", GetNameToMention(message), config.GetQuestionText()),
		DisableNotification: true,
	}

	if !SendMessage(newMessage) {
		log.Println("Сообщение не отправлено")
	}

	return true
}

// getCase строка с примером которая требует решения
func getCase(message *telegram.ReceiveMessage) string {
	number1, number2 := getNumbersAnExample(message)

	return fmt.Sprintf("%s + %s", number1, number2)
}

// getDecision  решение примера
func getDecision(message *telegram.ReceiveMessage) string {
	//number1, number2 := getNumbersAnExample(message)

	//number1, _ = strconv.ParseInt(number1, 10)
	//number2, _ = strconv.ParseInt(number2, 10)
	return ""
}

// getTheNumbersAnExample Строки с числами для использования в примере и в ответе решения
func getNumbersAnExample(message *telegram.ReceiveMessage) (string, string) {
	if message.Message.NewChatMember.Id == 0 {
		log.Fatalf("Новый пользователь содержит не верный id: %d", message.Message.NewChatMember.Id)
	}
	memberId := strconv.Itoa(int(message.Message.NewChatMember.Id))
	memberIdSlice := strings.Split(memberId, "")

	return memberIdSlice[len(memberIdSlice)-2], memberIdSlice[len(memberIdSlice)-1]
}
