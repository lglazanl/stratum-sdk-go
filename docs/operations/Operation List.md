
Create a instance stratumsdk pass three params 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```

Create a object the walletListPayload with params:
1. **DestType(string)**
2. **DirectionType(string)**
3. **OperationEid(int)**
4. **OperationStatus(string)**
5. **OperationTsFrom(int)**
6. **OperationTsTo(int)**
6. **OperationType(string)**
7. **OperationUpdTsFrom(int)**
8. **OperationUpdTsTo(int)**
9. **WalletEid(int)**
10. **WalletGroupEid(int)**
11. **WalletId(int)**

```go
operationPayload := &stratumsdk.OperationPayload{
	WalletEid        : 100
	WalletBalanceMin : 10
	WalletBalanceMax : 100
	WalletGroupEid   : 10
	WalletGroupId    : 12
	WalletType       : "checking"
	Currency         : "BTC"
	CurrencyType     : "crypto"
}

operationList, apiErr, err := client.Operations().List(operationPayload)
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

[]OperationData

type OperationData struct {
	DestType           string `json:"dest_type"`     
	DirectionType      string `json:"direction_type"` 
	OperationEid       int    `json:"operation_eid"`
	OperationStatus    string `json:"operation_status"`      
	OperationTsFrom    int    `json:"operation_ts_from"`     
	OperationTsTo      int    `json:"operation_ts_to"`       
	OperationType      string `json:"operation_type"`        
	OperationUpdTsFrom int    `json:"operation_upd_ts_from"` 
	OperationUpdTsTo   int    `json:"operation_upd_ts_to"`
	WalletEid          int    `json:"wallet_eid"`
	WalletGroupEid     int    `json:"wallet_group_eid"`
	WalletId           int    `json:"wallet_id"`
}

```

