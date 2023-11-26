package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	m := MakeHandler()
	m.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}
