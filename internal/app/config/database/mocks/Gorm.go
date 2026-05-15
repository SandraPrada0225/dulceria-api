package mocks

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (sqlmock.Sqlmock, *gorm.DB) {
	//crea una base de datos falsa, completamente simulada
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)

	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}
	//usa la base de datos false
	mysqlConfig := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(mysqlConfig, &gorm.Config{})

	if err != nil {
		log.Fatalf("[gorm open] %s", err)
	}

	return mock, db
}
