package initialize

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Conn() {
	var err error
	// Db, err = sql.Open("mysql", "root:gumostza168@tcp(localhost:3306)/courtfield_reservation?parseTime=true")
	dsn := "root:gumostza168@tcp(127.0.0.1:3306)/courtfield_reservation?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
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
		Db.Exec(q)
	}

	return nil
}
