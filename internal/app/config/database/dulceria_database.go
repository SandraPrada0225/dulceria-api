package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Este es el cliente de la conexion, el puento central de la conexion
type Client struct{}

// aqui abrimos la conexion, inicializamos GORM
// devolvemos una instancia DB
func (c Client) Connect() (db *gorm.DB, err error) {
	data := GetConnectionLocal()

	connectionString := data.GetUrl()
	//se conecta el ORM
	db, err = gorm.Open(
		//usamos el driver MySQL
		mysql.Open(connectionString),
		&gorm.Config{},
	)

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}
