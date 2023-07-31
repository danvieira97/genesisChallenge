package repositories

import (
	"gorm.io/gorm"
)

type Container struct {
	Converter ConverterRepository
}

type Options struct {
	db *gorm.DB
}

func New(opts Options) *Container {
	return &Container{
		Converter: ConverterRepository(),
	}
}
