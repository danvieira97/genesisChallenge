package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danvieira97/genesisChallenge/model"
	"github.com/danvieira97/genesisChallenge/repositories"
	"github.com/gin-gonic/gin"
)

var allCurrency []CurrencyConverterResponse

func CurrencyConverter(ctx *gin.Context) {
	request :=
		CurrencyConverterRequest{
			amount: ctx.Param("amount"),
			from:   ctx.Param("from"),
			to:     ctx.Param("to"),
			rate:   ctx.Param("rate"),
		}

	if err := request.Validate(); err != nil {
		responseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res := prepareResponse(&request)

	repo := repositories.RepoGorm{}

	newConverter := model.CurrencyConverter{
		ValorConvertido: res.ValorConvertido,
		SimboloMoeda:    res.SimboloMoeda,
	}

	err := repo.Create(ctx.Request.Context(), newConverter)
	if err != nil {
		fmt.Println(err.Error())
	}

	allCurrency = append(allCurrency, res)

	responseSuccess(ctx, "currency-converter", res)
}

func AllCurrencyConverter(ctx *gin.Context) {
	repo := repositories.RepoGorm{}
	allConverter, err := repo.GetAll(context.Background())

	if err != nil {
		fmt.Println(err.Error())
	}

	responseSuccess(ctx, "all-currency-converter", allConverter)
}
