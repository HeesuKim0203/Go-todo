package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render
var id int = 0

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreateAt  time.Time `json:"creaeted_at"`
}

var todoMap map[int]*Todo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}

	for _, v := range todoMap {
		list = append(list, v)
	}

	rd.JSON(w, http.StatusOK, list)
}

func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id += 1
	todoMap[id] = &Todo{id, name, false, time.Now()}

	rd.JSON(w, http.StatusCreated, todoMap[id])
}

type Success struct {
	Success bool `json:"success"`
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func changeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"

	log.Println(complete)

	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)

	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", postTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", deleteTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", changeTodoHandler).Methods("GET")

	return r
}
