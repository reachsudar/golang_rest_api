package main

import (
	repository "rest-api/Repository"
	"rest-api/configdb"
	"rest-api/controller"
	"rest-api/service"

	"github.com/gin-gonic/gin"
)

var (
	empRepository repository.EmployeeRepository = repository.NewRepository()

	empInterface  service.EmployeeInterface     = service.New(empRepository)
	empController controller.EmployeeController = controller.New(empInterface)
)

func main() {
	configdb.Connect()
	r := gin.Default()

	r.GET("/employees", func(c *gin.Context) {
		empController.GetAll(c)
	})
	r.GET("/employees/:id", func(c *gin.Context) {
		empController.GetById(c)
	})
	r.POST("/employees", func(c *gin.Context) {
		empController.Save(c)
	})
	r.PUT("/employees/:id", func(c *gin.Context) {
		empController.Update(c)
	})
	r.DELETE("/employees/:id", func(c *gin.Context) {
		empController.Delete(c)
	})

	r.Run()
}
