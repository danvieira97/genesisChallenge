package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"gorm.io/gorm"
)

type Converters struct {
	gorm.Model
	ID              int `gorm:"primaryKey"`
	ValorConvertido float64
	SimboloMoeda    string
}

func ConnectDB() {
	db, err := gorm.Open("mysql", DB_CONNECTION)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Converters{})
}

func (c *Converters) GetAll()
