package handler

import (
	"net/http"

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

	allCurrency = append(allCurrency, res)

	responseSuccess(ctx, "currency-converter", res)
}

func AllCurrencyConverter(ctx *gin.Context) {
	repositories.ConverterRepository.GetAll()
}
