package tests

import (
    "github.com/go-task/models"
    "github.com/go-task/repositories"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestFilterByStatusCodeValidInput(t *testing.T) {
    filterObject :=  models.Filter{StatusCode:"authorised"}

    actual := repositories.FilterByStatusCode("authorised",&filterObject)
    expected := true
    assert.Equal(t,expected,actual)
}

func TestFilterByStatusCodeInvalidInput(t *testing.T) {
    filterObject :=  models.Filter{StatusCode:"authorised"}
    actual := repositories.FilterByStatusCode("auth",&filterObject)
    expected := false
    assert.Equal(t,expected,actual)
}


func TestFilterByCurrencyValidInput(t *testing.T) {
    filterObject :=  models.Filter{Currency:"USD"}

    actual := repositories.FilterByStatusCode("USD",&filterObject)
    expected := true
    assert.Equal(t,expected,actual)
}

func TestFilterByCurrencyEmpty(t *testing.T) {
    filterObject :=  models.Filter{StatusCode:"authorised"}
    actual := repositories.FilterByStatusCode("",&filterObject)
    expected := false
    assert.Equal(t,expected,actual)
}

func TestFilterByValidMinMaxAmount(t *testing.T) {
    filterObject :=  models.Filter{AmountMin:200,AmountMax:400}

    actual := repositories.FilterByAmount(250,&filterObject)
    expected := true
    assert.Equal(t,expected,actual)
}

func TestFilterByAmountNotBetweenMinAndMax(t *testing.T) {
    filterObject :=  models.Filter{AmountMin:200,AmountMax:400}

    actual := repositories.FilterByAmount(100,&filterObject)
    expected := false
    assert.Equal(t,expected,actual)
}

func TestFilterByAmountWithMinOnly(t *testing.T) {
    filterObject :=  models.Filter{AmountMin:200}

    actual := repositories.FilterByAmount(220,&filterObject)
    expected := true
    assert.Equal(t,expected,actual)
}

func TestFilterByAmountWithMaxOnly(t *testing.T) {
    filterObject :=  models.Filter{AmountMax:200}

    actual := repositories.FilterByAmount(180,&filterObject)
    expected := true
    assert.Equal(t,expected,actual)
}

func TestFilterByAmountEqualMin(t *testing.T) {
    filterObject :=  models.Filter{AmountMin:200,AmountMax:500}
    actual := repositories.FilterByAmount(200,&filterObject)
    expected := true
    assert.Equal(t,expected,actual)
}

func TestFilterByAmountEqualMax(t *testing.T) {
    filterObject :=  models.Filter{AmountMin:200,AmountMax:500}
    actual := repositories.FilterByAmount(500,&filterObject)
    expected := true
    assert.Equal(t,expected,actual)
}