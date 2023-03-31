package DBconnection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-sql-driver/mysql"
)

var (
	DB              *sql.DB
	connectionError error
)

func Connect() {
	content, errcon := ioutil.ReadFile("D:\\Projects\\IVANLAVR\\kursach\\DBconnection\\config.json")
	if errcon != nil {
		log.Fatal(errcon.Error())
	}

	var config map[string]string

	if err := json.Unmarshal(content, &config); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print("config: ")
	fmt.Println(config)

	var cfg = mysql.Config{
		User:                 config["User"],
		Passwd:               config["Password"],
		Net:                  config["NetworkConnection"],
		Addr:                 config["Address"],
		DBName:               config["Name"],
		AllowNativePasswords: true,
	}

	DB, connectionError = sql.Open("mysql", cfg.FormatDSN())
	if connectionError != nil {
		log.Fatal(connectionError.Error())
	}

	log.Println("connected successfully!")
}

func Test() []string {
	var Names []string

	Rows, err := DB.Query("SELECT NameOfProduct FROM products")
	if err != nil {
		log.Printf("Test (cannot execute query): %v", err)
	}
	defer Rows.Close()

	for Rows.Next() {
		var Name string
		if scanErr := Rows.Scan(&Name); scanErr != nil {
			log.Printf("Test (cannot scan row to string) %v", scanErr)
			return nil
		}

		Names = append(Names, Name)
	}

	if rowErr := Rows.Err(); rowErr != nil {
		log.Printf("Test (unknown error) %v", rowErr)
		return nil
	}

	return Names
}
