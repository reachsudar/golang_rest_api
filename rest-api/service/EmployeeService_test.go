package service_test

import (
	"database/sql"
	"log"
	repository "rest-api/Repository"
	"rest-api/configdb"
	entity "rest-api/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var emp = &entity.Employee{
	ID:           1,
	FirstName:    "sudar",
	MiddleName:   "sasi",
	LastName:     "Kumar",
	Gender:       "female",
	Salary:       12000,
	DOB:          time.Now(),
	Email:        "sasisudar@gmail.com",
	Phone:        9866654323,
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
		"sasi", "kumar", "female", 12000, "2018-12-10T13:49:51.141Z", "sasisudar@gmail.com", 9866654323,
		"22  yendon road", "carnegie", "vic", 3030, 9876543, 230000.00)

	mock.ExpectQuery(query).WillReturnRows(rows)
	// w := httptest.NewRecorder()
	// gin.SetMode(gin.TestMode)
	// c, _ := gin.CreateTestContext(w)

	employees := repo.GetAll()
	//assert.NoError(t, err)
	assert.NotEmpty(t, employees)

}
