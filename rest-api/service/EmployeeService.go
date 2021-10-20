package service

import (
	"log"
	"net/http"
	entity "rest-api/Entity"
	repository "rest-api/Repository"
	"rest-api/configdb"
	_ "rest-api/configdb"

	"github.com/gin-gonic/gin"
)

// Employee Services
type EmployeeInterface interface {
	GetAll(c *gin.Context) ([]*entity.Employee, error)
	GetById(c *gin.Context) (entity.Employee, error)
	Save(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type employeeStruct struct {
	repository repository.EmployeeRepository
}

//Constructor
func New(repo repository.EmployeeRepository) EmployeeInterface {
	return &employeeStruct{
		repository: repo,
	}

}

// GetAll Employees
func (service *employeeStruct) GetAll(c *gin.Context) ([]*entity.Employee, error) {
	emp1 := make([]*entity.Employee, 0)

	rows, err := configdb.DB.Query("Select * from employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		emp := new(entity.Employee)
		rows.Scan(&emp.ID, &emp.FirstName, &emp.MiddleName,
			&emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB,
			&emp.Email, &emp.Phone, &emp.AddressLine1, &emp.AddressLine2, &emp.State,
			&emp.PostCode, &emp.TFN, &emp.SuperBalance)
		if err != nil {
			log.Fatalln(err)
		}
		emp1 = append(emp1, emp)
		c.JSON(200, emp1)
	}
	return emp1, nil

}

// Get Employee Details with ID
func (service *employeeStruct) GetById(c *gin.Context) (entity.Employee, error) {
	id := c.Params.ByName("id")
	var emp entity.Employee
	rows, err := configdb.DB.Query("SELECT * FROM employees where id=?", id)
	if err != nil {
		log.Fatal("Error in Query", err)

	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&emp.ID, &emp.FirstName, &emp.MiddleName,
			&emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB,
			&emp.Email, &emp.Phone, &emp.AddressLine1, &emp.AddressLine2, &emp.State,
			&emp.PostCode, &emp.TFN, &emp.SuperBalance); err != nil {
			log.Fatal(err)

		}

		c.IndentedJSON(http.StatusOK, emp)
	}
	if rows.Err(); err != nil {
		c.JSON(200, gin.H{"error": "data not found"})
		panic(err)
	}
	return emp, nil
}

// Add Emplyee Employee
func (service *employeeStruct) Save(c *gin.Context) error {
	var emp entity.Employee
	c.BindJSON(&emp)

	_, err := configdb.DB.Exec("INSERT INTO employees(id,first_name,middle_name,last_name,gender,salary,dob,email,phone,address_line1,address_line2,state,post_code,tfn,super_balance) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", emp.ID, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB, emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN, emp.SuperBalance)
	if err != nil {
		c.String(http.StatusInternalServerError, "error")
	}

	c.JSON(http.StatusCreated, emp)
	return err

}

//Update Employee Detail BY ID
func (service *employeeStruct) Update(c *gin.Context) error {
	var emp entity.Employee
	c.BindJSON(&emp)
	id := c.Params.ByName("id")
	_, err := configdb.DB.Exec("Update employees set first_name=?,middle_name=?,last_name=?,gender=?,salary=?,dob=?,email=?,phone=?,address_line1=?,address_line2=?,state=?,post_code=?,tfn=?,super_balance=? WHERE id =?", emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB, &emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN, emp.SuperBalance, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "error")
	}
	c.JSON(http.StatusOK, emp)
	return err
}

// Delete Employee By ID
func (service *employeeStruct) Delete(c *gin.Context) error {
	id := c.Params.ByName("id")
	_, err := configdb.DB.Exec("delete from Employees where id=?", id)
	if err != nil {
		c.JSON(404, gin.H{"Sorry": "ID not found"})
	}
	c.JSON(200, gin.H{
		"Deleted From ID": id,
	})
	return err
}
