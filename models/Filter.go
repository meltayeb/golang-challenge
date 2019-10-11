package models

type Filter struct {
	Provider   string `schema:"provider"`
	StatusCode   string `schema:"statusCode"`
	AmountMin   float32 `schema:"amountMin"`
	AmountMax   float32 `schema:"amountMax"`
	Currency   string `schema:"currency"`
}
