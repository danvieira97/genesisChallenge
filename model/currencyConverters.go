package model

type CurrencyConverter struct {
	ID              int     `gorm:"primary key;autoIncrement" json:"id"`
	ValorConvertido float64 `db:"valor_convertido" json:"valor_convertido"`
	SimboloMoeda    string  `db:"simbolo_moeda" json:"simbolo_moeda"`
}
