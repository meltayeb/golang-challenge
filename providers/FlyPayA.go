package providers

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

type DataProcessingFlyA struct{
    Transactions         []*FlyPayA    `json:"transactions"`
}

type FlyPayA struct {
    Amount         float32    `json:"amount"`
    Currency       string     `json:"currency"`
    StatusCode     int         `json:"statusCode"`
    OrderReference string      `json:"orderReference"`
    TransactionId  string      `json:"transactionId"`
}

func (FlyPayA) GetStatusCode(code int) string {
    var status string
    if code == 1 {status = "authorised"}
    if code == 2 {status = "decline"}
    if code == 3 {status = "refunded"}
    return status
}

func (FlyPayA) GetAllTransactions() []*FlyPay {
    var dataProcessing *DataProcessingFlyA
    pwd, _ := os.Getwd()
    file, _ := ioutil.ReadFile(pwd+"/data/flypayA.json")
    if err := json.Unmarshal([]byte(file), &dataProcessing);err != nil {
        panic(err)
    }

    var resultFly []* FlyPay
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

func NewFlyA() * FlyPayA  {
    return &FlyPayA{}
}