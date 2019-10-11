package providers

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

type DataProcessingFlyB struct{
    Transactions         []*FlyPayB    `json:"transactions"`

}
type FlyPayB struct {
    Amount         float32    `json:"value"`
    Currency       string     `json:"transactionCurrency"`
    StatusCode     int         `json:"statusCode"`
    OrderReference string      `json:"orderInfo"`
    TransactionId  string      `json:"paymentId"`
}

func (FlyPayB) GetStatusCode(code int) string {
    var status string
    if code == 100 {status = "authorised"}
    if code == 200 {status = "decline"}
    if code == 300 {status = "refunded"}
    return status
}

func (FlyPayB) GetAllTransactions() []*FlyPay {
    var dataProcessing *DataProcessingFlyB
    pwd, _ := os.Getwd()
    file, _ := ioutil.ReadFile(pwd+"/data/flypayB.json")
    if err := json.Unmarshal([]byte(file), &dataProcessing);err != nil {
        panic(err)
    }

    var resultFly []*FlyPay
    for _,item := range dataProcessing.Transactions {
        resultFly = append(resultFly,&FlyPay{
            Amount : item.Amount,
            Currency : item.Currency,
            StatusCode : item.GetStatusCode(item.StatusCode),
            OrderReference : item.OrderReference,
            TransactionId : item.TransactionId,
        })
    }

    return resultFly
}

func NewFlyB() * FlyPayB  {
    return &FlyPayB{}
}