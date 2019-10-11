package repositories

import (
	"github.com/go-task/models"
	"github.com/go-task/providers"
)

func ApplyFilters(filter *models.Filter) [] *providers.FlyPay {

	result := FilterByProvider(filter)
	var filteredResult [] *providers.FlyPay
	for _, item := range result {
		var validator = true
		if !FilterByStatusCode(item.StatusCode, filter) {
			validator = false
			continue
		}
		if !FilterByCurrency(item.Currency, filter) {
			validator = false
			continue
		}
		if !FilterByAmount(item.Amount, filter) {
			validator = false
			continue
		}
		if validator {
			filteredResult = append(filteredResult, item)
		}

	}
	return filteredResult
}

func FilterByProvider(filter *models.Filter) [] *providers.FlyPay {
	FlyA := providers.NewFlyA()
	FlyB := providers.NewFlyB()
	var DataOfA [] *providers.FlyPay
	var DataOfB [] *providers.FlyPay

	if filter.Provider == "flypayA" {
		DataOfA = FlyA.GetAllTransactions()
	} else if filter.Provider == "flypayB" {
		DataOfB = FlyB.GetAllTransactions()
	} else {
		DataOfA = FlyA.GetAllTransactions()
		DataOfB = FlyB.GetAllTransactions()
	}
	return append(DataOfA, DataOfB...)
}

func FilterByStatusCode(item string, filter *models.Filter) bool {
	if filter.StatusCode == "" {
		return true
	}

	if filter.StatusCode == item {
		return true
	}
	return false
}

func FilterByCurrency(item string, filter *models.Filter) bool {
	if filter.Currency == "" {
		return true
	}
	if filter.Currency == item {
		return true
	}
	return false
}

func FilterByAmount(item float32, filter *models.Filter) bool {
	if item >= filter.AmountMin && filter.AmountMax == 0{
		return true
	}

	if item <= filter.AmountMax && filter.AmountMin == 0{
		return true
	}

	if item >= filter.AmountMin && item <= filter.AmountMax {
		return true
	}
	return false
}
