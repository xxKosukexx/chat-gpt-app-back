package main

import (
	"app/internal/db"
	"app/internal/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.ChatRoom{}, &model.ChatMessage{})
}
