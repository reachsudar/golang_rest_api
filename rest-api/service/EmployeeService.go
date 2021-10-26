package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	repository "rest-api/Repository"
	"strconv"
	"strings"
	"sync"
	"time"

	entity "rest-api/entity"

	"github.com/gin-gonic/gin"
)

// Employee Services
type EmployeeInterface interface {
	GetAll(c *gin.Context) []entity.Employee
	GetById(c *gin.Context) entity.Employee
	Save(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type employeeStruct struct {
	repository repository.EmployeeRepository
}
type Database struct {
	Db *sql.DB
}
type Emp struct {
	ID           string  `json:"employeeId"`
	SuperBalance float64 `json:"superBalance"`
}

//Constructor
func New(repo repository.EmployeeRepository) EmployeeInterface {
	return &employeeStruct{
		repository: repo,
	}

}

// with channels
// GetAll Employees
// func (service *employeeStruct) GetAll(c *gin.Context) []entity.Employee {
// 	employees := service.repository.GetAll()
// 	done := make(chan bool)
// 	for i := 0; i < len(employees); i++ {
// 		go Super(i, employees, done)
// 	}
// 	<-done
// 	time.Sleep(3 * time.Second)
// 	return employees
// }

//with wait group
func (service *employeeStruct) GetAll(c *gin.Context) []entity.Employee {
	employees := service.repository.GetAll()
	var wg sync.WaitGroup
	for i := 0; i < len(employees); i++ {
		wg.Add(1)
		go Super(i, employees, &wg)
	}

	time.Sleep(100 * time.Millisecond)
	return employees
}

// go routine
func Super(Num int, employees []entity.Employee, wg *sync.WaitGroup) {

	super, err := getSuper(strconv.Itoa(employees[Num].ID))
	if err != nil {
		employees[Num].SuperBalance = 0
	}
	employees[Num].SuperBalance = super
	wg.Done()
}

// // With Anonymous Function
// func (service *employeeStruct) GetAll(c *gin.Context) []entity.Employee {
// 	employees := service.repository.GetAll()
// 	var wg sync.WaitGroup
// 	for i := 0; i < len(employees); i++ {
// 		num := i
// 		wg.Add(1)
// 		go func(num int) {
// 			super, err := getSuper(strconv.Itoa(employees[num].ID))
// 			if err != nil {
// 				employees[num].SuperBalance = 0
// 			}
// 			employees[num].SuperBalance = super
// 		}(num)
// 	}
// 	wg.Wait()
// 	return employees
// }

// Get Employee Details with ID
func (service *employeeStruct) GetById(c *gin.Context) entity.Employee {
	id := c.Params.ByName("id")
	employee := service.repository.GetById(id) //
	empId := strconv.Itoa(employee.ID)
	super, err := getSuper(empId)
	if err != nil {
		c.JSON(500, "Error in fetching the Superbalance ")
	}

	employee.SuperBalance = super

	return employee
}

//get super from mb
func getSuper(id string) (float64, error) {
	var emp Emp
	var url = "http://localhost:4545/ato/employees/?/super"
	url = strings.Replace(url, "?", id, 1)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return 0.0, nil
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := Read(resp)
		if err != nil {
			log.Fatalln("Error in Reading Response")
		}
		emp = UnmarshalData(body)

		fmt.Println("super from mounte", emp.SuperBalance)
	}
	if resp.StatusCode == 404 {
		log.Println("Error in Response 404")
	}
	return emp.SuperBalance, nil
}

//ReadAll function
func Read(resp *http.Response) ([]byte, error) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error in Readall")
	}
	return body, err

}

// Unmarshalling function
func UnmarshalData(body []byte) Emp {
	var emp Emp
	json.Unmarshal(body, &emp)
	return emp
}

// Add Emplyee Employee
func (service *employeeStruct) Save(c *gin.Context) error {
	var emp entity.Employee
	err := c.BindJSON(&emp)
	service.repository.Save(emp)
	return err
}

//Update Employee Detail BY ID
func (service *employeeStruct) Update(c *gin.Context) error {
	var emp entity.Employee
	err := c.BindJSON(&emp)
	if err != nil {
		return err
	}
	empid := c.Params.ByName("id")
	i, err := strconv.Atoi(empid)
	emp.ID = i
	service.repository.Update(emp)
	return err
}

// Delete Employee By ID
func (service *employeeStruct) Delete(c *gin.Context) error {
	var emp entity.Employee
	id := c.Params.ByName("id")
	i, err := strconv.Atoi(id)
	emp.ID = i
	service.repository.Delete(emp)
	return err
}
