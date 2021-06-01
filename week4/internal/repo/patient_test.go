package repo

import (
	"phudt/week4/internal/model"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

// PatientTestSuite test suite
type PatientTestSuite struct {
	suite.Suite
	repo *patient
	mock sqlmock.Sqlmock
}

func (suite *PatientTestSuite) SetupTest() {
	// Mock DB
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	suite.mock = mock
	suite.mock.MatchExpectationsInOrder(true)

	// Gorm from Mock
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	suite.repo = &patient{db: gormDB}
}

func (suite *PatientTestSuite) TearDownTest() {
	// TODO
}

func TestPatientTestSuite(t *testing.T) {
	suite.Run(t, new(PatientTestSuite))
}

func (suite *PatientTestSuite) TestCreate() {
	// input
	m := model.Patient{
		Fullname: "Tran Phong Phu",
		Address:  "HCM",
		Birthday: "2021-01-01",
		Gender:   1,
		Age:      10,
	}
	// Mock SQL Query
	var expectedId int64 = 3
	sql := "INSERT INTO `patient` (`fullname`,`address`,`birthday`,`gender`,`age`) VALUES (?,?,?,?,?)"
	suite.mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(expectedId, 1))
	// Excute
	actual, err := suite.repo.Create(m)
	if err != nil {
		suite.Error(err, "Should not return error here")
	}
	if actual.Id != expectedId {
		suite.Fail("should return expected id", "expectedId", expectedId, "actualId", actual.Id)
	}
	if err := suite.mock.ExpectationsWereMet(); err != nil {
		suite.Error(err, "there were unfulfilled expectations")
	}
}
