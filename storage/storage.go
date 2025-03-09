package storage

import (
	"database/sql"
	"log"
	"os"

	erration "retsAPI/serv/error"

	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	db *sql.DB
}

func StorageCheck(storagePath string) error{
	d := Db{}
	var err error

	d.db, err = sql.Open("sqlite3", storagePath)
	if err != nil {
		erration.LogError(err, "ERROR_STORAGE_FILE_OPEN")
		return err
	}

	d.tableExists()
	
	return nil
}

func (d *Db)tableExists(){
	var exists bool

	query := `
	SELECT COUNT(*) > 0 
	FROM sqlite_master 
	WHERE type='table' AND name=?;
	`

	d.db.QueryRow(query, "url").Scan(&exists)
	if !exists {
		d.createTables()
	}
}

func (d *Db)createTables(){
	stmt, err := d.db.Prepare(`
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
