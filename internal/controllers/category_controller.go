package controllers

import (
	"book-api/internal/models"
	"book-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{ svc *services.CategoryService }

func NewCategoryController(s *services.CategoryService) *CategoryController { return &CategoryController{s} }

func (cc *CategoryController) Create(c *gin.Context) {
	var req models.Category
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	id, err := cc.svc.Create(req)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (cc *CategoryController) GetAll(c *gin.Context) {
	list, err := cc.svc.GetAll()
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, list)
}

func (cc *CategoryController) GetByID(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	item, err := cc.svc.GetByID(id)
	if err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, item)
}

func (cc *CategoryController) Delete(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	if err := cc.svc.Delete(id); err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"message":"deleted"})
}