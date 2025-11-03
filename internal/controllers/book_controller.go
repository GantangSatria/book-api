package controllers

import (
	"book-api/internal/models"
	"book-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct{ svc *services.BookService }

func NewBookController(s *services.BookService) *BookController { return &BookController{s} }

func (bc *BookController) Create(c *gin.Context) {
	var req models.Book
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	id, err := bc.svc.Create(req)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (bc *BookController) GetAll(c *gin.Context) {
	list, err := bc.svc.GetAll()
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, list)
}

func (bc *BookController) GetByID(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	item, err := bc.svc.GetByID(id)
	if err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, item)
}

func (bc *BookController) Delete(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	if err := bc.svc.Delete(id); err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"message":"deleted"})
}

func (bc *BookController) GetByCategory(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	list, err := bc.svc.GetByCategory(id)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, list)
}
