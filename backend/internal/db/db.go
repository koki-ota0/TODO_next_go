package db

import (
	"context"
	"log"
	"todo-server/internal/config"
	"todo-server/internal/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func Init() {
	// 設定の読み込み
	config.LoadConfig()

	// データベースの接続
	db = pg.Connect(&pg.Options{
		Addr:     config.AppConfig.DBAddr,
		User:     config.AppConfig.DBUser,
		Password: config.AppConfig.DBPassword,
		Database: config.AppConfig.DBName,
	})

	// 接続テスト
	err := db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	// テーブルを作成（または既に存在する場合はスキップ）
	err = createSchema()
	if err != nil {
		log.Fatalf("Error creating schema: %v\n", err)
	}
}

func createSchema() error {
	models := []interface{}{
		(*models.Todo)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDB() *pg.DB {
	return db
}
