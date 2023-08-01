package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Converters struct {
	gorm.Model
	ID              int     `gorm:"primaryKey" json:"id"`
	ValorConvertido float64 `db:"valor_convertido" json:"valor_convertido"`
	SimboloMoeda    string  `db:"simbolo_moeda" json:"simbolo_moeda"`
}

func ConnectDB() {
	dsn := "host=localhost user=gorm password=genesis dbname=gorm port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Converters{})
}
