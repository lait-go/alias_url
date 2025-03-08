package storage

import (
	"database/sql"
	"log"
	"os"

	erration "retsAPI/serv/error"

	_ "github.com/mattn/go-sqlite3"
)

func StorageWork(url, alias string) error{
	db, err := sql.Open("sqlite3", "./storage/storage.db")
	if err != nil {
		erration.LogError(err, "ERROR_STORAGE_FILE_OPEN")
		return err
	}

	tableExists(db)


	return nil
}

func tableExists(db *sql.DB){
	stmt, _:= db.Prepare(`SELECT * FROM url`)

	file , _ := stmt.Exec()
	if file == nil{
		createTables(db)
	}
}

func createTables(db *sql.DB){
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		erration.LogError(err, "ERROR_STORAGE_PREPARE")
	}

	_, err = stmt.Exec()
	if err != nil {
		erration.LogError(err, "ERROR_STORAGE_EXEC")
	}
}

func FileExists(file_name string) error{
	if _, err := os.Stat(file_name); err != nil {
		if os.IsNotExist(err) {
			log.Printf("FILE_IS_NOT_EXIST_ERROR: %s\n", file_name)
		}else{
			log.Println("FILE_ERROR")
		}
		return err
	}

	return nil
}