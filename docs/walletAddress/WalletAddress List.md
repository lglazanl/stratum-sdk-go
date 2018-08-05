
Create a instance stratumsdk pass three params 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```

Create a object the WalletsAddressesListPayload with params:
1. **WalletAddressEid (int)**  your key for sync with integration
2. **WalletId(int)**  your WalletId in create new wallet 
3. **WalletEid(int)** your key for sync with integration wallet
4. **Currency(string)** your Currency

```go
walletsAddressesListPayload := &stratumsdk.WalletsAddressesListPayload{
		WalletId: 10,
}
walletAddressList, apiErr, err := client.WalletsAddresses().List(walletsAddressesListPayload)
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
 []WalletAddressData 

type WalletAddressData struct {
	WalletAddressEid   int    `json:"wallet_address_eid"`
	WalletId           int64  `json:"wallet_id"`
	WalletAddressLabel string `json:"wallet_address_label"`
	WalletAddress      string `json:"wallet_address"`
}


```

