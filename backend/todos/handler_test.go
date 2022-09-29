package todos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Fatur525/learninggo/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/send", handler.CreateTodo)
	payload := `{"task": "task 1"}`
	req, err := http.NewRequest("POST", "/send", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	type response struct {
		Message string      `json:"message"`
		Data    model.Todos `json:"data"`
		Error   string
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	fmt.Println(res)
	assert.Equal(t, "success", res.Message)
	assert.Equal(t, "task 1", res.Data.Task)
}

func TestGetTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	db.Create(&model.Todos{
		Task: "task 1",
	})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", handler.GetTodos)
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	type response struct {
		Message string        `json:"message"`
		Data    []model.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Equal(t, "task 1", res.Data[0].Task)
}
