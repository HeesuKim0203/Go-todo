package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/goWeb/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetTodoList()

	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) postTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	todo := a.db.AddTodo(name)

	rd.JSON(w, http.StatusCreated, todo)
}

func (a *AppHandler) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) changeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"

	ok := a.db.CompleteTodo(id, complete)

	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()
	a := &AppHandler{Handler: r, db: model.NewDBHandler()}

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.postTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.deleteTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.changeTodoHandler).Methods("GET")

	return a
}
