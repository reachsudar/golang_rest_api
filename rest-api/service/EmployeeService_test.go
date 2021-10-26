package service_test

import (
	"database/sql"
	"log"
	repository "rest-api/Repository"
	"rest-api/configdb"
	entity "rest-api/entity"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

//go test -tags=unit -v
var emp = &entity.Employee{
	ID:           113,
	FirstName:    "sudar",
	MiddleName:   "sasi",
	LastName:     "Kumar",
	Gender:       "female",
	Salary:       12000,
	DOB:          time.Now(),
	Email:        "sasisudar@gmail.com",
	Phone:        9866654,
	AddressLine1: "22  yendon road",
	AddressLine2: "carnegie",
	State:        "vic",
	PostCode:     3030,
	TFN:          9876543,
	SuperBalance: 230000.00,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err)
	}
	return db, mock
}

func TestGetAll(t *testing.T) {
	repo, err := repository.NewRepository("mysql", configdb.Connect(), 3, 3)

	if err != nil {
		log.Fatalln(err)
	}

	_, mock := NewMock()

	query := "Select (.+) from employees"

	rows := sqlmock.NewRows([]string{"id", "firstname", "middlename", "lastname", "gender", "salary",
		"dob", "email", "phone", "Address1", "address2", "state", "postcode", "tfn", "super"}).AddRow(12, "sudar",
		"sasi", "kumar", "female", 12000, "2018-12-10T13:49:51.141Z", "sasisudar@gmail.com", 9866654,
		"22  yendon road", "carnegie", "vic", 3030, 987654, 230000.00)

	mock.ExpectQuery(query).WillReturnRows(rows)

	employees := repo.GetAll()

	assert.NotEmpty(t, employees)

}
func TestGetByID(t *testing.T) {
	repo, err := repository.NewRepository("mysql", configdb.Connect(), 3, 3)

	if err != nil {
		log.Fatalln(err)
	}

	_, mock := NewMock()
	query := "Select (.+) from employees where id = ?"
	rows := sqlmock.NewRows([]string{"id", "firstname", "middlename", "lastname", "gender", "salary",
		"dob", "email", "phone", "Address1", "address2", "state", "postcode", "tfn", "super"}).AddRow(&emp.ID, &emp.FirstName, &emp.MiddleName, &emp.LastName,
		emp.Gender, emp.Salary, emp.DOB, emp.Email, emp.Phone, emp.AddressLine1,
		emp.AddressLine2, emp.State, emp.PostCode, emp.TFN, emp.SuperBalance)

	mock.ExpectQuery(query).WillReturnRows(rows)
	ID := strconv.Itoa(emp.ID)
	employee := repo.GetById(ID)
	assert.NotNil(t, employee)

}
func TestSave(t *testing.T) {
	repo, err := repository.NewRepository("mysql", configdb.Connect(), 3, 3)

	if err != nil {
		log.Fatalln(err)
	}

	_, mock := NewMock()
	query := "INSERT INTO employees"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(emp.ID, emp.FirstName, emp.MiddleName, emp.LastName,
		emp.Gender, emp.Salary, emp.DOB, emp.Email, emp.Phone, emp.AddressLine1,
		emp.AddressLine2, emp.State, emp.PostCode, emp.TFN, emp.SuperBalance).WillReturnResult(sqlmock.NewResult(0, 1))
	repo.Save(*emp)
	assert.NoError(t, err)

}
func TestDelete(t *testing.T) {
	repo, err := repository.NewRepository("mysql", configdb.Connect(), 3, 3)

	if err != nil {
		log.Fatalln(err)
	}

	_, mock := NewMock()
	query := "DELETE FROM employees WHERE id = ?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))
	repo.Delete(*emp)
	assert.NoError(t, err)
	// mock.MatchExpectationsInOrder(false)
	// mock.ExpectExec(query).WithArgs(emp.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	// repo.Delete(entity.Employee{})

}
func TestUpdate(t *testing.T) {
	repo, err := repository.NewRepository("mysql", configdb.Connect(), 3, 3)

	if err != nil {
		log.Fatalln(err)
	}

	_, mock := NewMock()
	query := "Update employees set first_name=?,middle_name=?,last_name=?,gender=?,salary=?,dob=?,email=?,phone=?,address_line1=?,address_line2=?,state=?,post_code=?,tfn=?,super_balance=? WHERE id =?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(emp.ID, emp.FirstName, emp.MiddleName, emp.LastName,
		emp.Gender, emp.Salary, emp.DOB, emp.Email, emp.Phone, emp.AddressLine1,
		emp.AddressLine2, emp.State, emp.PostCode, emp.TFN, emp.SuperBalance).WillReturnResult(sqlmock.NewResult(0, 0))

	repo.Update(*emp)
	assert.NoError(t, err)
}
