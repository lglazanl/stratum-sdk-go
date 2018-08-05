
Create a instance stratumsdk pass three params 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```

Create a object a new WalletGroupData with params:
1. **WalletGroupEid (int)**  your key for sync with integration
2. **WalletGroupId(int)**  is a unique id the Wallet group
3. **WalletGroupLabel(int)**  label name for walletgroup

```go
wallgroupCreatePayload := &stratumsdk.WalletGroupPayload{
		WalletGroupLabel: "99xtest",
		WalletGroupEid:   200,
}
walletGroup, apiErr, err := client.WalletGroup().Create(wallgroupCreatePayload)
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

