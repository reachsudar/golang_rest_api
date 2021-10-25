package main

import (
	"net/http/httptest"
	"rest-api/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	w := httptest.NewRecorder()

	var service service.EmployeeInterface

	r := gin.Default()

	gin.SetMode(gin.TestMode)

	// Get endpoint Testing
	r.GET("/employees", func(c *gin.Context) {

	})
	t.Run("Get", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})

	//GetbyId endpoint Testing
	r.GET("/employees/:id", func(c *gin.Context) {
		service.GetAll(c)
	})
	t.Run("GetbyId", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})

	//Post api Testing
	r.POST("/employees", func(c *gin.Context) {
		service.Save(c)
	})
	t.Run("Post", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})

	//Put api Testing
	r.PUT("/employees/:id", func(c *gin.Context) {

		service.Update(c)
	})
	t.Run("Put", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})

	//Delete api Testing
	r.DELETE("/employees/:id", func(c *gin.Context) {

		service.Delete(c)
	})
	t.Run("Delete", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})

}
