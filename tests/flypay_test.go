package tests

import (
    "github.com/go-task/providers"
    "github.com/stretchr/testify/assert"
    "testing"
)


func TestGetStatusCodeValidInput(t *testing.T) {
    FlyPayAObject := providers.NewFlyA()
    actual := FlyPayAObject.GetStatusCode(2)
    expected := "decline"
    assert.Equal(t,expected,actual)
}