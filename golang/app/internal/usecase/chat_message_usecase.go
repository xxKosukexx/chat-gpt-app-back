package usecase

import (
	"app/internal/model"
	"app/internal/repository"
	"app/internal/validator"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type IChatMessageUsecase interface {
	Create(chatMessage model.ChatMessage) (model.ChatMessageResponse, error)
	RequestChatGPTAnswer(chatRoomId uint, question string) (string, error)
}

type chatMessageUsecase struct {
	crr repository.IChatMessageRepository
	crv validator.IChatMessageValidator
}

func NewChatMessageUsecase(crr repository.IChatMessageRepository, crv validator.IChatMessageValidator) IChatMessageUsecase {
	return &chatMessageUsecase{crr, crv}
}

func (cru *chatMessageUsecase) Create(chatMessage model.ChatMessage) (model.ChatMessageResponse, error) {
	if err := cru.crv.ChatMessageValidate(chatMessage); err != nil {
		return model.ChatMessageResponse{}, err
	}
	newChatMessage := model.ChatMessage{Question: chatMessage.Question, Answer: chatMessage.Answer, ChatRoomId: chatMessage.ChatRoomId}
	if err := cru.crr.CreateChatMessage(&newChatMessage); err != nil {
		return model.ChatMessageResponse{}, err
	}
	resChatMessage := model.ChatMessageResponse{
		ID:       newChatMessage.ID,
		Question: newChatMessage.Question,
		Answer:   newChatMessage.Answer,
	}
	return resChatMessage, nil
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGptRequestData struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func (cru *chatMessageUsecase) RequestChatGPTAnswer(chatRoomId uint, question string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	chatMessages, err := cru.crr.GetAllChatMessageByChatRoomId(chatRoomId)
	if err != nil {
		return "", err
	}
	var messages []Message
	for _, chatMessage := range chatMessages {
		messages = append(messages, Message{Role: "user", Content: chatMessage.Question})
		messages = append(messages, Message{Role: "system", Content: chatMessage.Answer})
	}
	messages = append(messages, Message{Role: "user", Content: question})

	data := ChatGptRequestData{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("CHAT_GPT_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result map[string]interface{}

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return "", err
	}

	return result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string), nil
}
