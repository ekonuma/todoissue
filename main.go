/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ekonuma/todoissue/cmd"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbSession *gorm.DB

func main() {
	initDB()
	cmd.Execute()
	http.HandleFunc("/callback", handler)
	http.ListenAndServe(":8080", nil)
	cmd.GetToken()
}

type Authentication struct {
	name        string `gorm:"primaryKey"`
	AccessToken string
}

func initDB() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	dbSession = db
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Authentication{})
}

func handler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	w.Write([]byte("Authorization successful! You can close this window."))
	dbSession.Where(Authentication{name: "todoist"}).Assign(Authentication{AccessToken: code}).FirstOrCreate(&Authentication{name: "todoist", AccessToken: code})
	fmt.Println("Authorization code received:", code)
	os.Exit(0)
}
