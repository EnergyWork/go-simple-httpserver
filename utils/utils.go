package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	models "go-test-program/models"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	dbport   = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "postgres"
)

func GetSqlConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, dbport, user, password, dbname)
	var db, err = sql.Open("postgres", psqlconn)
	return  db, err
}

func Remove(arr []models.Message, i int) []models.Message {
	if i == 0 {
		return append(arr[1:])
	} else if i == len(arr) {
		return append(arr[:len(arr)-1])
	} else {
		return append(arr[:i], arr[i+1:]...)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetAnswer(status uint, message interface{}) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
