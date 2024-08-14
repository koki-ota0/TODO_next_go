package main

import (
	"log"
	"net/http"
	"todo-server/internal/api"
	"todo-server/internal/db"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// データベース初期化
	db.Init()

	// ルータの設定
	router := mux.NewRouter()

	// エンドポイントの設定
	router.HandleFunc("/todos", api.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todos", api.GetTodosHandler).Methods("GET")
	router.HandleFunc("/todos/{id}", api.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todos/{id}", api.DeleteTodoHandler).Methods("DELETE")

	// CORS 設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // ダブルクオートを追加
		AllowedHeaders: []string{"Content-Type"},
	})

	// CORS ミドルウェアを追加
	handler := c.Handler(router)

	// サーバーの起動
	log.Println("Server is running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
