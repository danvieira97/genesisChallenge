package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/danvieira97/genesisChallenge/model"
	"gorm.io/gorm"
)

type repoGorm struct {
	db *gorm.DB
}

func (repo *repoGorm) GetAll(ctx context.Context) ([]model.CurrencyConverter, error) {
	var converters []model.CurrencyConverter
	repo.db.Table("currency-converters").Find(&converters)

	if repo.db.Error != nil {
		fmt.Println(repo.db.Error)
		return nil, errors.New("currency converter not found")
	}

	return converters, nil
}

func (repo *repoGorm) Create(ctx context.Context, newConverter model.CurrencyConverter) error {
	repo.db.Table("currency-converters").Create(&newConverter)

	if repo.db.Error != nil {
		fmt.Println(repo.db.Error)
		return errors.New("currency converter didn't create")
	}

	return nil
}
