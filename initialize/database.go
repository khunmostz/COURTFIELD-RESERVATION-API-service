package initialize

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var Db *sql.DB

func Conn() {
	var err error
	Db, err = sql.Open("mysql", "root:gumostza168@tcp(localhost:3306)/courtfield_reservation?parseTime=true")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to database.")
}

func LoadSQLFile(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	rawQuery := strings.Split(string(file), ";")
	fmt.Println(rawQuery)
	for _, q := range rawQuery {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := Db.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
