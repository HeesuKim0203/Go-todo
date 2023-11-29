package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(MakeHandler())

	defer ts.Close()

	const TESTNAME1 string = "Test todo1"
	const TESTNAME2 string = "Test todo2"

	id1 := todoCreatedTest(assert, ts, TESTNAME1)
	id2 := todoCreatedTest(assert, ts, TESTNAME2)

	todos := todoGetTodoListTest(assert, ts, 2)
	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal(t.Name, TESTNAME1)
		} else if t.ID == id2 {
			assert.Equal(t.Name, TESTNAME2)
		} else {
			assert.Error(fmt.Errorf("testID should be id1 or id2"))
		}
	}

	todoCompeleteUpdateTest(assert, ts, id1, "true")
	todos = todoGetTodoListTest(assert, ts, 2)
	for _, t := range todos {
		if t.ID == id1 {
			assert.True(t.Completed)
		}
	}

	todoDeleteTest(assert, ts, id1)
	todos = todoGetTodoListTest(assert, ts, 1)
	for _, t := range todos {
		assert.Equal(t.ID, id2)
	}
}

func todoCreatedTest(assert *assert.Assertions, ts *httptest.Server, testName string) int {

	var todo Todo

	// Todo Create Test
	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {testName}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	// Todo name Test
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, testName)
	id1 := todo.ID

	return id1
}

func todoGetTodoListTest(assert *assert.Assertions, ts *httptest.Server, testListLen int) []*Todo {

	// Todo Get Test
	resp, err := http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	// TodoList Len Test
	todos := []*Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), testListLen)

	return todos
}

func todoCompeleteUpdateTest(assert *assert.Assertions, ts *httptest.Server, testId int, testComplete string) {

	// Todo compelete update Test
	resp, err := http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(testId) + "?complete=" + testComplete)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	sucess := Success{}
	err = json.NewDecoder(resp.Body).Decode(&sucess)
	assert.NoError(err)
	assert.Equal(sucess.Success, true)

}

func todoDeleteTest(assert *assert.Assertions, ts *httptest.Server, testId int) {

	// Todo Delete Test
	req, _ := http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(testId), nil)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}
