package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	var db *gorm.DB
	var err error
	var retryInterval time.Duration = 5
	maxRetries := 5

	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dsn := fmt.Sprintf(`%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True`,
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(time.Second * retryInterval)
	}

	// エラーが発生した場合、エラー内容を表示
	if err != nil {
		log.Fatal(err)
	}
	// 接続に成功した場合、「db connected!!」と表示する
	fmt.Println("db connected!!")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
