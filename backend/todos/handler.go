package todos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetTodos(c *gin.Context) {
	todos, status, err := h.Service.GetTodos()
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    todos,
	})
}

func (h *Handler) CreateTodo(c *gin.Context) {
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, status, err := h.Service.CreateTodos(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    res,
	})
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	var req DataRequest

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, status, err := h.Service.UpdateTodos(id)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    res,
	})
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	status, err := h.Service.DeleteTodos(id)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "delete success",
	})
}
