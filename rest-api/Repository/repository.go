package Repository

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/configdb"
	entity "rest-api/entity"

	//"rest-api/configdb"

	_ "github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

// Employee Repository
type EmployeeRepository interface {
	GetAll() []entity.Employee
	GetById(string) entity.Employee
	Save(emp entity.Employee) error
	Update(emp entity.Employee) error
	Delete(emp entity.Employee) error
}

//Database Struct
type Database struct {
	DB *sql.DB
}

//constructor
func NewRepository(dialect string, configdb mysql.Config, idleconn, maxconn int) (EmployeeRepository, error) {
	db, err := sql.Open(dialect, configdb.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("connection Failed")
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping Error")
	}
	db.SetMaxIdleConns(idleconn)
	db.SetMaxOpenConns(maxconn)

	return &Database{DB: db}, err

}

func (db *Database) GetAll() []entity.Employee {
	emp1 := make([]entity.Employee, 0)
	rows, err := configdb.DB.Query("Select * from employees")
	if err != nil {
		log.Println("Error in Query")
	}
	defer rows.Close()

	for rows.Next() {
		var emp entity.Employee
		rows.Scan(&emp.ID, &emp.FirstName, &emp.MiddleName,
			&emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB,
			&emp.Email, &emp.Phone, &emp.AddressLine1, &emp.AddressLine2, &emp.State,
			&emp.PostCode, &emp.TFN, &emp.SuperBalance)
		if err != nil {
			log.Fatalln(err)
		}
		emp1 = append(emp1, emp)

	}
	return emp1
}
func (db *Database) GetById(id string) entity.Employee {

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

	}

	return emp
}

func (db *Database) Save(emp entity.Employee) error {
	_, err := configdb.DB.Query("INSERT INTO employees(id,first_name,middle_name,last_name,gender,salary,dob,email,phone,address_line1,address_line2,state,post_code,tfn,super_balance) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		emp.ID, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB,
		emp.Email, emp.Phone, emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN, emp.SuperBalance)
	if err != nil {
		log.Println("Error in Query")
	}
	return err
}

func (db *Database) Update(emp entity.Employee) error {
	_, err := configdb.DB.Query("Update employees set first_name=?,middle_name=?,last_name=?,gender=?,salary=?,dob=?,email=?,phone=?,address_line1=?,address_line2=?,state=?,post_code=?,tfn=?,super_balance=? WHERE id =?",
		emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB, &emp.Email, emp.Phone,
		emp.AddressLine1, emp.AddressLine2, emp.State, emp.PostCode, emp.TFN, emp.SuperBalance, emp.ID)
	if err != nil {
		log.Println("Error in Query")
	}
	return err

}
func (db *Database) Delete(emp entity.Employee) error {
	_, err := configdb.DB.Exec("delete from Employees where id=?", emp.ID)
	if err != nil {
		log.Println("Error in Query")
	}
	return err
}
