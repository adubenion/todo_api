package dao

import (
	"log"
	. "github.com/adubenion/todo_api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoDAO struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	TODO_COLLECTION = "todos"
	USERS_COLLECTION = "users"
)

func(t *TodoDAO) Connect() {
	session, err := mgo.Dial(t.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(t.Database)
}

func (t *TodoDAO) FindAllTodos() ([]Todo, error) {
	var todos []Todo
	err := db.C(TODO_COLLECTION).Find(bson.M{}).All(&todos)
	return todos, err
}
 
func (t *TodoDAO) FindTodosById(_id string) (Todo, error) {
	var todo Todo
	err := db.C(TODO_COLLECTION).FindId(bson.ObjectIdHex(_id)).One(&todo)
	return todo, err
}
 
func (t *TodoDAO) InsertTodos(todo Todo) error {
	err := db.C(TODO_COLLECTION).Insert(&todo)
	return err
}
 
func (t *TodoDAO) DeleteTodos(todo Todo) error {
	err := db.C(TODO_COLLECTION).Remove(&todo)
	return err
}
 
func (t *TodoDAO) UpdateTodos(todo Todo) error {
	err := db.C(TODO_COLLECTION).UpdateId(todo.ID, &todo)
	return err
}

func (m *TodoDAO) FindAllUsers() ([]User, error) {
	var users []User
	err := db.C(USERS_COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}
 
func (m *TodoDAO) FindUsersByEmail(_id string) (User, error) {
	var user User
	err := db.C(USERS_COLLECTION).FindId(bson.ObjectIdHex(_id)).One(&user)
	return user, err
}
 
func (m *TodoDAO) InsertUsers(user User) error {
	err := db.C(USERS_COLLECTION).Insert(&user)
	return err
}
 
func (m *TodoDAO) DeleteUsers(user User) error {
	err := db.C(USERS_COLLECTION).Remove(&user)
	return err
}
 
func (m *TodoDAO) UpdateUsers(user User) error {
	err := db.C(USERS_COLLECTION).UpdateId(user.ID, &user)
	return err
}