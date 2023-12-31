package model

import "time"

var id int = 0

type memoryHandler struct {
	todoMap map[int]*Todo
}

func (m *memoryHandler) GetTodoList() []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddTodo(name string) *Todo {
	id += 1
	m.todoMap[id] = &Todo{id, name, false, time.Now()}
	return m.todoMap[id]
}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler) CompleteTodo(id int, complete bool) bool {
	if todo, ok := m.todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}

func (m *memoryHandler) Close() {
	m.todoMap = make(map[int]*Todo)
}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)

	return m
}
