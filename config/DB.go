package config

import (
	"gintut/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// modelの宣言

var DB *gorm.DB

// DB接続関数
func ConnectDB() {
	err := godotenv.Load("./.env")
	checkErr(err)

	//ここは自分の環境に合った設定で行う
	dsn := "host=" + os.Getenv("DATABASE_ADDRESS") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PASSWORD") + " dbname=" + os.Getenv("DATABASE_NAME") + " port=" + os.Getenv("DATABASE_PORT") + " sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	checkErr(err)

	//DBのマイグレーション
	db.AutoMigrate(&models.Memo{})

	DB = db
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
