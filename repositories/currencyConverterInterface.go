package repositories

import (
	"context"

	"github.com/danvieira97/genesisChallenge/model"
)

type UserRepository interface {
	Create(ctx context.Context, newConverter model.CurrencyConverter) error
	GetAll(ctx context.Context) ([]model.CurrencyConverter, error)
}
