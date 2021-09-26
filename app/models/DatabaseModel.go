package models

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"

	"github.com/lavender-snow/neosvr-musicscore/config"
	_ "github.com/mattn/go-sqlite3"
)

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", config.Config.DBfile)
	if err != nil {
		log.Printf("DB接続時にエラーが発生しました。\r\n%s", err.Error())
	}
	return db
}

func InsertScoreData(scoreData string) (string, error) {

	byteData := []byte(scoreData)
	id := fmt.Sprintf("%x", md5.Sum(byteData))

	db := connectDB()
	defer db.Close()

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
