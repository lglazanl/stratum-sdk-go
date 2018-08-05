
Create a instance stratumsdk pass three param 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```

Create a object WalletPayload with params:
1. **WalletGroupId (int)**  field walletGroup by Id

```go
walletGroupPayload := &stratumsdk.WalletGroupPayload{
		WalletGroupId: 21,
}
walletGroup, apiErr, err := client.WalletGroup().Get(walletGroupPayload)
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
The call returns the structure below

```go 
type WalletGroupData struct {
	WalletGroupEid   int    `json:"wallet_group_eid"`
	WalletGroupId    int    `json:"wallet_group_id"`
	WalletGroupLabel string `json:"wallet_group_label"`
}



```

