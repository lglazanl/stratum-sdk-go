
Create a instance stratumsdk pass three param 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```
Create a object WalletPayload with params:
1. **Currency (string)**  currency is fee field
1. **DestType (int)**  wallet by Id created 
1. **OperationType (int)** your operation fee (withdraw)

```go
feePayload := &stratumsdk.FeePayload{
		Currency: "BTC",
}
feeList, apiErr, err := client.Operations().Fees(feePayload)
```

treat the errors
```go 
if err != nil {
	fmt.Printf("sdk error:  %s ", err.Error())
}
if apiErr != nil {
	fmt.Printf("apiError: %s ", apiErr.Data)
}
```
The call returns array the structure below

```go 
[]FeeData
type FeeData struct {
	Currency      string  `json:"currency"`
	DestType      string  `json:"dest_type"`
	OperationType string  `json:"operation_type"`
	OperationFee  float64 `json:"operation_fee"`
}



```

