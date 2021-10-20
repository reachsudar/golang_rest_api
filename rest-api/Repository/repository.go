package repository

import (
	"database/sql"
	entity "rest-api/Entity"
	"rest-api/configdb"

	_ "github.com/gin-gonic/gin"
)

type EmployeeRepository interface {
	GetAll() []entity.Employee
	GetById() entity.Employee
	Save(emp entity.Employee)
	Update(emp entity.Employee)
	Delete(emp entity.Employee)
}
type database struct {
	conn *sql.DB
}

func NewRepository() EmployeeRepository {
	db := configdb.DB
	return &database{conn: db}

}

func (db *database) GetAll() []entity.Employee {
	var emp []entity.Employee
	db.GetAll()
	return emp
}
func (db *database) GetById() entity.Employee {
	var emp entity.Employee
	db.GetById()
	return emp
}

func (db *database) Save(emp entity.Employee) {

	db.Save(emp)

}
func (db *database) Update(emp entity.Employee) {

	db.Update(emp)

}
func (db *database) Delete(emp entity.Employee) {

	db.Delete(emp)

}
