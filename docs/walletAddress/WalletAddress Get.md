
Create a instance stratumsdk pass three param 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```

Create a object WalletPayload with params:
1. **WalletAddressEid (int)**  your key for sync with integration
1. **WalletId (int)**  your walletId 
1. **WalletAddressLabel (int)**  your labael created
1. **WalletAddress (int)** wallet address created in assign
1. **Currency (int)** your currency used in assign

```go
walletAddressesPayload := &stratumsdk.WalletAddressesPayload{
		WalletAddress: WalletAddressTest,
		Currency:      "BTC",
	}
walletAddress, apiErr, err := client.WalletsAddresses().Get(walletAddressesPayload)
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
type WalletAddressData struct {
	WalletAddressEid   int    `json:"wallet_address_eid"`
	WalletId           int64  `json:"wallet_id"`
	WalletAddressLabel string `json:"wallet_address_label"`
	WalletAddress      string `json:"wallet_address"`
}


```

