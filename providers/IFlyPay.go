package providers

type IFlyPay interface {
	getAllTransactions() [] *FlyPay
}

type FlyPay struct {
	Amount         float32
	Currency       string
	StatusCode     string
	OrderReference string
	TransactionId  string
}

type PaymentProvider struct {
	Provider IFlyPay
}

