package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler[T any] struct {
	service interface {
		GetAll() ([]T, error)
		GetByID(int) (T, error)
		Create(T) (T, error)
		Update(int, T) (T, error)
		Delete(int) error
	}
}

func NewHandler[T any](s interface {
	GetAll() ([]T, error)
	GetByID(int) (T, error)
	Create(T) (T, error)
	Update(int, T) (T, error)
	Delete(int) error
}) *Handler[T] {
	return &Handler[T]{service: s}
}

//acutal logic

func (h *Handler[T]) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}
func (h *Handler[T]) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}
func (h *Handler[T]) Create(c *gin.Context) {
	var input T
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := h.service.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, data)
}
func (h *Handler[T]) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input T
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := h.service.Update(id, input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}

func (h *Handler[T]) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "deleted"})
}

// middlwares
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		println(c.Request.Method, c.Request.URL.Path, time.Since(start).String())
	}
}

//routes

func (h *Handler[T]) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.GetByID)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
