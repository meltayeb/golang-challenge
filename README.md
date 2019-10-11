<p align="center">

![alt text](../fly365.png)

</p>

## Challenge Idea
The main idea of the task is to build an API that read the payment transaction data from files (Json format) and return them in the API response as json format.
There are two payment providers `flypayA` and `flypayB`.

Use these files as a source of data for the this task
- `flypayA` data is stored in [flypayA.json](./flypayA.json)
- `flypayB` data is stored in [flypayB.json](./flypayB.json)


flypay A transaction schema is 
```
{
  amount:1000,
  currency:'AUD',
  statusCode:1,
  orderReference:'2e58bd43-7abb'
  transactionId: '4fc2-a8d1'
}
```

we have three status for flypay A
- `authorised` which will have statusCode `1`
- `decline` which will have statusCode `2`
- `refunded` which will have statusCode `3`


flypay B transaction schema is 
```
{
  value:200,
  transactionCurrency:'AUD',
  statusCode:100,
  orderInfo:'2e58bd43-7abb'
  paymentId: '4dc2-a8a1'
}
```

we have three status for flypay B
- `authorised` which will have statusCode `100`
- `decline` which will have statusCode `200`
- `refunded` which will have statusCode `300`


## Acceptance Criteria

Implement this API `/api/payment/transaction `
- it should list all transaction which combine transactaions from all the available provider(`flypayA` and `flypay`B)
- it should be able to filter transaction by payment providers for example `/api/payment/transaction?provider=flypayA` it should return transaction for flypay only
- it should be able to filter transaction three statusCode (`authorised`, `decline`, `refunded`) for example /api/payment/transaction?statusCode=authorised it should return all the transactions from all payment providers that have status code authorised
- it should be able to filer by amount range for example `/api/payment/transaction?amountMin=10&amountMax=100` it should return all transaction between 10 and 100 including 10 and 100
- it should be able to filer by `currency` 
- it should be able to combine all this filter together 

## Expectations

Task will be evaluated based on
1. Code quality
2. Application performance in reading large files and it will be measured by time and memory usage (hint Concurreny may improve the performance but `be careful!` it may double the memory usage which will be affect the performance try to balance between time and memory)
3. Code scalabilty we should be able to add a new payment provider like flypayC with a new schema with the minimum possible changes 
5. Unit tests


## Bonus

1. Use Docker.
2. Benchmarck the task.