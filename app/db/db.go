package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB is method to initialize DB
func InitDB() {

	// DB接続
	log.Println("== DB 接続 ==")
	db, err := sql.Open("mysql", "root@/gosample")
	if err != nil {
		log.Println("DB 接続 エラー")
		log.Fatalln(err)
	}
	log.Println("== DB 接続 成功 ==")

	// 既存のテーブル削除
	log.Println("== DB 削除 ==")
	db.Exec("DROP TABLE users")
	log.Println("== DB 削除 成功 ==")

	// テーブル生成
	log.Println("== DB テーブル 作成 ==")
	_, err = db.Exec(`CREATE TABLE users ( 
		id INTEGER AUTO_INCREMENT PRIMARY KEY, 
		name VARCHAR(32) NOT NULL
		)`)
	if err != nil {
		outputErrorLog("usersテーブル作成失敗", err)
	}

	// データ投入
	log.Println("== DB データ投入 ==")
	_, err = db.Exec(`INSERT INTO users 
		(name)
		VALUES
		('John F Kenned')
	`)
	if err != nil {
		outputErrorLog("データ投入 失敗", err)
	}
	log.Println("== DB データ投入 完了 ==")
}

// CloseDB is method to close DB connection
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// ConnDB is method to get DB connection
func ConnDB() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root@/gosample")
		if err != nil {
			log.Println("DB 接続 エラー")
			log.Fatalln(err)
		}
	}
	return db
}

// output error log and stop app
func outputErrorLog(message string, err error) {
	log.Println(message)
	log.Fatalln(err)
}