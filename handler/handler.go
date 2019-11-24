package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shuStudent/model"
	"shuStudent/service/fetchName"
	"shuStudent/service/token"
)

func getStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	tokenInHeader := r.Header.Get("Authorization")
	if len(tokenInHeader) <= 7 {
		w.WriteHeader(401)
		return
	}
	idFromToken := token.StudentIdForToken(tokenInHeader[7:])
	if id != idFromToken {
		w.WriteHeader(403)
		return
	}
	student, err := model.Get(id)
	if err != nil || student.Name == "" {
		student, err = fetchName.FetchName(student, tokenInHeader[7:])
		model.Put(student)
		if err != nil {
			if err.Error() == "not logined" {
				w.WriteHeader(401)
			} else {
				w.WriteHeader(404)
			}
			return
		}
	}
	data, _ := json.Marshal(student)
	_, _ = w.Write(data)
}

func setStudentHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	student := model.Student{}
	err := json.Unmarshal(data, &student)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	tokenInHeader := r.Header.Get("Authorization")
	if len(tokenInHeader) <= 7 {
		w.WriteHeader(401)
		return
	}
	id := token.StudentIdForToken(tokenInHeader[7:])
	if id != student.Id {
		w.WriteHeader(403)
		return
	}
	model.Put(student)
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getStudentHandler(w, r)
	case "POST":
		setStudentHandler(w, r)
	case "PUT":
		setStudentHandler(w, r)
	}
}

func PingPongHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
