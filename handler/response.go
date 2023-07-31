package handler

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func responseError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func responseSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfull %s handler", op),
		"data":    data,
	})
}

type CurrencyConverterResponse struct {
	ValorConvertido float64 `json:"valorConvertido"`
	SimboloMoeda    string  `json:"simboloMoeda"`
}

func prepareResponse(req *CurrencyConverterRequest) (res CurrencyConverterResponse) {
	amountFloat, _ := strconv.ParseFloat(req.amount, 64)
	rateFloat, _ := strconv.ParseFloat(req.rate, 64)

	amountConverted := amountFloat * rateFloat
	amountConverted = math.Round(amountConverted*100) / 100

	res = CurrencyConverterResponse{
		ValorConvertido: amountConverted,
		SimboloMoeda:    "$",
	}

	return res
}
