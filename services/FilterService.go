package services

import (
	"github.com/go-task/models"
	"github.com/go-task/providers"
	"github.com/go-task/repositories"
)


func Filter(filter *models.Filter) [] *providers.FlyPay {

	result:= repositories.ApplyFilters(filter)
	return result
}

