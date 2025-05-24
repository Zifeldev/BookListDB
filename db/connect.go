package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var Conn *pgx.Conn

func ConnectDB(){
	var err  error
	_ = godotenv.Load()
	url := os.Getenv("DB_URL")	
	Conn, err = pgx.Connect(context.Background(), url)
	if err != nil{
		fmt.Println("❌ Ошибка подключения к базе:", err)
		os.Exit(1)
	}
	fmt.Println("Connected")
}

