package controller

import (
	"net/http"
	service "rest-api/service"

	"github.com/gin-gonic/gin"
)

// Employee Controller Interface

type EmployeeController interface {
	GetAll(ctx *gin.Context) error
	GetById(ctx *gin.Context) error
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.EmployeeInterface
}

// Constructor
func New(service service.EmployeeInterface) EmployeeController {
	return &controller{
		service: service,
	}
}

// GetAll Controller
func (c *controller) GetAll(ctx *gin.Context) error {
	employees := c.service.GetAll(ctx)

	data := gin.H{
		"employees": employees,
	}
	ctx.JSON(http.StatusOK, data)
	return nil
}

// GetbyId Controller
func (c *controller) GetById(ctx *gin.Context) error {
	employee := c.service.GetById(ctx) //
	data := gin.H{
		"Employee Details": employee,
	}
	ctx.JSON(http.StatusOK, data)
	return nil
}

// Save Employee Controller
func (c *controller) Save(ctx *gin.Context) error {
	err := c.service.Save(ctx)
	if err != nil {
		ctx.JSON(500, "Error in Inserting the Data")
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "New Employee Added",
	})
	return nil
}

// Update Controller
func (c *controller) Update(ctx *gin.Context) error {
	err := c.service.Update(ctx)
	if err != nil {
		ctx.JSON(500, "Error in Updating the Data")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Updated Employee Details",
	})
	return nil
}

// Delete Controller
func (c *controller) Delete(ctx *gin.Context) error {
	err := c.service.Delete(ctx)
	if err != nil {
		ctx.JSON(500, "Error in Deleting the data")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Deleted Employee",
	})
	return nil
}
