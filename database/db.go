package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type disconnectedError struct {
}

func (e disconnectedError) Error() string {
	return "database is closed"
}

type Database struct {
	// connection
	conn *sql.DB
}

func SanitizeString(input string) string {
	return strings.Replace(input, "'", "''", -1)
}

func (database *Database) Exec(query string) {
	if !database.IsConnectionOpened() {
		log.Print(disconnectedError{}.Error())
		return
	}

	_, err := database.conn.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func (database *Database) Query(query string) (*sql.Rows, error) {
	if !database.IsConnectionOpened() {
		log.Print()
		return nil, disconnectedError{}
	}

	return database.conn.Query(query)
}

func (database *Database) Connect(fileName string) error {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	database.conn = db

	return nil
}

func (database *Database) Disconnect() {
	database.conn.Close()
	database.conn = nil
}

func (database *Database) IsConnectionOpened() bool {
	return database.conn != nil
}

func (database *Database) CreateUniqueRecord(table string, values string) int64 {
	var err error
	if len(values) == 0 {
		_, err = database.conn.Exec(fmt.Sprintf("INSERT INTO %s DEFAULT VALUES ", table))
	} else {
		_, err = database.conn.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s)", table, values))
	}

	if err != nil {
		log.Fatal(err.Error())
		return -1
	}

	rows, err := database.conn.Query(fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", table))

	if err != nil {
		log.Fatal(err.Error())
		return -1
	}
	defer rows.Close()

	if rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err.Error())
			return -1
		}

		return id
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal("No record created")
	return -1
}
