package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-test-program/models"
	u "go-test-program/utils"
	"log"
	"net/http"
)

var messages []models.Message

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("GET request")
	var msgs []models.Message
	db, err := u.GetSqlConnection()
	if err != nil {
		u.Respond(w, u.GetAnswer(111, err.Error()))
		return
	}
	defer db.Close()
	query := fmt.Sprintf("select * from \"Message\"")
	rows, err := db.Query(query)
	if err != nil {
		u.Respond(w, u.GetAnswer(222, err.Error()))
		return
	}
	defer rows.Close()
	for rows.Next() {
		m := models.Message{}
		err := rows.Scan(&m.Id, &m.Text)
		if err != nil {
			u.Respond(w, u.GetAnswer(333, err.Error()))
			return
		}
		msgs = append(msgs, m)
	}
	//json.NewEncoder(w).Encode(msgs)
	u.Respond(w, u.GetAnswer(200, msgs))
}

func AddMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("POST request")
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		u.Respond(w, u.GetAnswer(444, err.Error()))
		return
	}
	query := fmt.Sprintf("insert into \"Message\" (text) values ('%s')", message.Text)
	db, err := u.GetSqlConnection()
	if err != nil {
		u.Respond(w, u.GetAnswer(555, err.Error()))
		return
	}
	defer db.Close()
	res, err := db.Exec(query)
	if err != nil {
		u.Respond(w, u.GetAnswer(666, err.Error()))
		return
	}
	mes, err := res.RowsAffected()
	u.Respond(w, u.GetAnswer(200, mes))
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("GET request")
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	query := fmt.Sprintf("select * from \"Message\" where id=%s", params["id"])
	db, err := u.GetSqlConnection()
	if err != nil {
		u.Respond(w, u.GetAnswer(0, err.Error()))
		return
	}
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		u.Respond(w, u.GetAnswer(0, err.Error()))
		return
	}
	defer rows.Close()
	for rows.Next() {
		m := models.Message{}
		err := rows.Scan(&m.Id, &m.Text)
		if err != nil {
			u.Respond(w, u.GetAnswer(0, err.Error()))
			return
		}
		json.NewEncoder(w).Encode(m)
		return
	}
	u.Respond(w, u.GetAnswer(0, "unknown id:" + params["id"]))
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE request")
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	//id, err := strconv.ParseUint(params["id"], 10, 64)
	/*if err != nil {
		u.Respond(w, u.GetAnswer(0, err.Error()))
		return
	}*/
	query := fmt.Sprintf("delete from \"Message\" where id=%s", params["id"])
	db, err := u.GetSqlConnection()
	if err != nil {
		u.Respond(w, u.GetAnswer(0, err.Error()))
		return
	}
	defer db.Close()
	res, err := db.Exec(query)
	if err != nil {
		u.Respond(w, u.GetAnswer(0, err.Error()))
		return
	}
	mes, err := res.RowsAffected()
	u.Respond(w, u.GetAnswer(200, mes))
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT request")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		u.Respond(w, u.GetAnswer(444, err.Error()))
		return
	}
	query := fmt.Sprintf("update \"Message\" set text='%s' where id='%s'", message.Text, params["id"])
	db, err := u.GetSqlConnection()
	if err != nil {
		u.Respond(w, u.GetAnswer(0, err.Error()))
		return
	}
	defer db.Close()
	res, err := db.Exec(query)
	if err != nil {
		u.Respond(w, u.GetAnswer(0, err.Error()))
		return
	}
	mes, err := res.RowsAffected()
	u.Respond(w, u.GetAnswer(200, mes))
}