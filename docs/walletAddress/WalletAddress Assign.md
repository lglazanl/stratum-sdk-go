
Create a instance stratumsdk pass three params 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```

Create a object a new WalletGroupData with params:
1. **WalletAddressEid (int)**  your key for sync with integration
2. **WalletAddressLabel(int)**  label name for walletAddress
3. **WalletId(int)**  is a unique id the Wallet
3. **Currency(int)**  currency assign your wallet

```go
walletAddressesPayload := &stratumsdk.WalletAddressesPayload{
		WalletAddressEid:   200,
		WalletAddressLabel: "99xtest",
		WalletId:           40,
}
walletAddress, apiErr, err := client.WalletsAddresses().Assign(walletAddressesPayload)
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

