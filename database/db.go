package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Converters struct {
	gorm.Model
	ID              int     `gorm:"primaryKey" json:"id"`
	ValorConvertido float64 `db:"valor_convertido" json:"valor_convertido"`
	SimboloMoeda    string  `db:"simbolo_moeda" json:"simbolo_moeda"`
}

func ConnectDB() {
	db, err := gorm.Open("mysql", DB_CONNECTION)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Converters{})
}
