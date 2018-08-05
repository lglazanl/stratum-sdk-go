
Create a instance stratumsdk pass three params 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```

Create a object a new WalletData with params: 
1. **WalletGroupId (int)**  is a unique id the Wallet group
2. **WalletLabel (string)** is a your custom label the wallet
3. **Currency (string)** is a currency for create a new wallet 
4. **WalletType(string)**  
5. **WalletEid(int)**  your key for sync with integration

```go
walletCreatePayload := &stratumsdk.WalletData{
		WalletGroupId: 10,
		WalletLabel:   "99xtest",
		Currency:      "BTC",
		WalletType:    "checking",
                WalletEid      100,
}
wallet, apiErr, err := client.Wallets().Create(walletCreatePayload)
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
type WalletData struct {
	WalletId           int64   `json:"wallet_id"`
	WalletEid          int64   `json:"wallet_eid"`
	WalletLabel        string  `json:"wallet_label"`
	WalletBalance      float64 `json:"wallet_balance"`
	WalletGroupId      int64   `json:"wallet_group_id"`
	WalletGroupLabel   string  `json:"wallet_group_label"`
	WalletGroupEid     int64   `json:"wallet_group_eid"`
	WalletType         string  `json:"wallet_type"`
	CurrencyName       string  `json:"currency_name"`
	Currency           string  `json:"currency"`
	CurrencyType       string  `json:"currency_type,omitempty"`
	CurrencyUnit       string  `json:"currency_unit"`
	CurrencyUnitSymbol string  `json:"currency_unit_symbol"`
	CurrencyUnitDigits int     `json:"currency_unit_digits"`
}



```

