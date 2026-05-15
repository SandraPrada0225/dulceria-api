package database

import "fmt"

// datos necesarios para conectarnos a MySQL
type connectionData struct {
	Host     string
	Schema   string
	UserName string
	Password string
	Dialect  string
}

// retorna una instancia de configuracion
func GetConnectionLocal() connectionData {
	return connectionData{
		Host:     "localhost:3306", //nuestra maquina:puerto MySQL
		Schema:   "dulceria",       //nombre de la base de datos
		UserName: "root",           //usuario MySQL
		Password: "",
		Dialect:  "mysql",
	}
}

// DNS: nombre d ela fuente de datos
// root:@tcp(localhost:3306)/dulceria?parseTime=true
func (c connectionData) GetUrl() (url string) {
	url = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true", //le dice a MySQL/GORM convierte DATETIME/TIMESTAMP a time.Time de Go
		c.UserName,
		c.Password,
		c.Host,
		c.Schema,
	)
	return
}
