package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
	
	"github.com/gorilla/mux"
	. "github.com/adubenion/todo_api/models"
	. "github.com/adubenion/todo_api/dao"
	. "github.com/adubenion/todo_api/config"
)

var config = Config{}
var dao = TodoDAO{}




func GetTodoEndpoint(w http.ResponseWriter, r *http.Request) {
	todos, err := dao.FindAllTodos()
	if err != nil {
		respondWithError( w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, todos)
}

func CreateTodoEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	todo.ID = bson.NewObjectId()
	if err := dao.InsertTodos(todo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, todo)
}

func UpdateTodoEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateTodoEndpoint TODO")
}

func DeleteTodoEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteTodoEndpoint TODO")
}

func LoginDoEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "LoginDoEndpoint TODO")
}

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateUserEndpoint TODO")
}

func AuthEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AuthEndpoint TODO")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	//TODOS
	r.HandleFunc("/api/todo/", GetTodoEndpoint).Methods("GET")
	r.HandleFunc("/api/create_todo/", CreateTodoEndpoint).Methods("POST")
	r.HandleFunc("/api/update_todo/", UpdateTodoEndpoint).Methods("PUT")
	r.HandleFunc("/api/delete_todo/", DeleteTodoEndpoint).Methods("DELETE")
	//USER
	r.HandleFunc("/login/do", LoginDoEndpoint).Methods("POST")
	r.HandleFunc("/api/create_user", CreateUserEndpoint).Methods("POST")
	r.HandleFunc("/auth", AuthEndpoint).Methods("GET")

	return r	
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := newRouter()
	log.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is online.")
}