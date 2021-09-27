package models

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"

	"github.com/lavender-snow/neosvr-musicscore/config"
	_ "github.com/mattn/go-sqlite3"
)

// connectDB DB接続処理
func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", config.Config.DBfile)
	if err != nil {
		log.Printf("DB接続時にエラーが発生しました。\r\n%s", err.Error())
	}
	return db
}

// InsertScoreData DBへ譜面情報を登録する
func InsertScoreData(scoreData string) (string, error) {

	byteData := []byte(scoreData)
	id := fmt.Sprintf("%x", md5.Sum(byteData))

	db := connectDB()
	defer db.Close()

	// 既に同一IDのデータが存在する場合はINSERT文を実行しない
	if data, _ := SelectScoreData(id); data == "" {
		query := "INSERT INTO score_data('id','data') VALUES($1, $2);"
		stmt, _ := db.Prepare(query)

		defer stmt.Close()

		if _, err := stmt.Exec(id, scoreData); err != nil {
			log.Printf(err.Error())
			return "", err
		}
	}

	return id, nil
}

// SelectScoreData DBからIDに対応する譜面情報を取得する
func SelectScoreData(id string) (string, error) {
	db := connectDB()
	defer db.Close()

	query := "SELECT data FROM SCORE_DATA WHERE ID = $1;"
	row := db.QueryRow(query, id)

	var scoreData string
	err := row.Scan(&scoreData)

	if err != nil {
		return "", err
	}

	return scoreData, nil
}
