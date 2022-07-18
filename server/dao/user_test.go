package dao

import (
	"chess/proto"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

var mock sqlmock.Sqlmock
var gormDB *gorm.DB

func init() {
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatal("mock mysql err:", err)
	}
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("mock mysql err:", err)
	}
	DB = gormDB
}

func TestSelectMobile(t *testing.T) {
	req := &proto.RegisterRequest{
		Mobile: "123",
	}
	rows := sqlmock.NewRows([]string{"mobile"}).
		AddRow(req.Mobile)
	mock.ExpectQuery("^select mobile FROM `user` WHERE mobile = \\?").
		WithArgs(req.Mobile).WillReturnRows(rows)
	_, e := SelectMobile(req)
	if e != nil {
		if e != gorm.ErrRecordNotFound {
			t.Fatalf("fail to get mobile: %s", e)
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
