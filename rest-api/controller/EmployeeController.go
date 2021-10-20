package controller

import (
	"net/http"
	"rest-api/service"

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
	employees, err := c.service.GetAll(ctx)
	if err != nil {
		return err
	}
	data := gin.H{
		"employees": employees,
	}
	ctx.JSON(http.StatusOK, data)
	return nil
}

// GetbyId Controller
func (c *controller) GetById(ctx *gin.Context) error {
	employee, err := c.service.GetById(ctx)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, employee)
	return nil
}

// Save Employee Controller
func (c *controller) Save(ctx *gin.Context) error {
	err := c.service.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Update Controller
func (c *controller) Update(ctx *gin.Context) error {
	err := c.service.Update(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete Controller
func (c *controller) Delete(ctx *gin.Context) error {
	err := c.service.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
