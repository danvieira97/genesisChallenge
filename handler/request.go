package handler

import (
	"fmt"
	"strconv"
)

type CurrencyConverterRequest struct {
	amount string
	from   string
	to     string
	rate   string
}

/*
* De Dólar para Real;
* De Euro para Real;
* De BTC para Dolar;
* De BTC para Real;
 */

func currencyConverterPossibilites(from, to string) error {
	if (from == "BRL") && to != "USD" || to != "EUR" {
		return fmt.Errorf("currency converter with %s it's just possible with USD and EUR", from)
	}
	if from == "USD" && to != "BRL" {
		return fmt.Errorf("currency converter with %s it's just possible with USD and EUR", from)
	}
	return nil
}

func errParamIsRequiredOrIncorrect(name string, typ string) error {
	if typ == "" {
		return fmt.Errorf("param: %s is required, url example: exchange/{amount}/{from}/{to}/{rate}", name)
	}
	return fmt.Errorf("param: %s need to be a %s, url example: exchange/10/BRL/USD/4.50", name, typ)
}

func (r *CurrencyConverterRequest) Validate() error {
	if r.amount == "" {
		return errParamIsRequiredOrIncorrect("amount", "")
	}
	if _, err := strconv.ParseFloat(r.amount, 64); err != nil {
		return errParamIsRequiredOrIncorrect("amount", "number")
	}
	if r.from == "" {
		return errParamIsRequiredOrIncorrect("from", "")
	}
	if r.from != "USD" && r.from != "EUR" && r.from != "BRL" && r.from != "BTC" {
		return fmt.Errorf("param: %s need be USD, EUR, BRL or BTC", "from")
	}
	if r.to == "" {
		return errParamIsRequiredOrIncorrect("to", "")
	}
	if r.to != "USD" && r.to != "EUR" && r.to != "BRL" && r.to != "BTC" {
		return fmt.Errorf("param: %s need be USD, EUR, BRL or BTC", "to")
	}
	if r.rate == "" {
		return errParamIsRequiredOrIncorrect("rate", "")
	}
	if _, err := strconv.ParseFloat(r.rate, 64); err != nil {
		return errParamIsRequiredOrIncorrect("rate", "number")
	}
	if r.to == r.from {
		return fmt.Errorf("params: %s and %s can't be equals", "from", "to")
	}

	return nil
}
