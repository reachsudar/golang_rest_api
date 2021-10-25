package main

import (
	repository "rest-api/Repository"
	"rest-api/configdb"
	"rest-api/controller"
	"rest-api/service"

	"github.com/gin-gonic/gin"
)

var (
	empRepository repository.EmployeeRepository = &repository.Database{}
	empInterface  service.EmployeeInterface     = service.New(empRepository)
	EmpController controller.EmployeeController = controller.New(empInterface)
)

var Repo repository.EmployeeRepository

func init() {
	repo, err := repository.NewRepository("mysql", configdb.Connect(), 3, 3)
	if err != nil {
		panic(nil)
	}
	Repo = repo
}

func main() {
	//connect to db
	configdb.Connect()

	// set router
	r := gin.Default()

	//Register the Handlers

	r.GET("/employees", func(c *gin.Context) {
		EmpController.GetAll(c)
	})
	r.GET("/employees/:id", func(c *gin.Context) {
		EmpController.GetById(c)
	})
	r.POST("/employees", func(c *gin.Context) {
		EmpController.Save(c)
	})
	r.PUT("/employees/:id", func(c *gin.Context) {
		EmpController.Update(c)
	})
	r.DELETE("/employees/:id", func(c *gin.Context) {
		EmpController.Delete(c)
	})

	r.Run("localhost:8888")
}
