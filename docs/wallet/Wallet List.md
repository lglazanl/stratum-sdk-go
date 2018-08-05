
Create a instance stratumsdk pass three params 
1. **apiUser** (sent by the team Stratum)
2. **apiSecret** (sent by the team Stratum)
3. **sandbox** ( if true in developer mode)


```go
var sandbox = true
client := stratumsdk.Initial("apiUser","apiSecret", sandbox)
```

Create a object the walletListPayload with params:
1. **WalletEid(int)**  your key for sync with integration
2. **WalletBalanceMin (float)** min balance in wallet
3. **WalletBalanceMax (float)** max balance in wallet
4. **WalletGroupEid (int)**  your key for sync with integration
5. **WalletGroupId(int)**  is a unique id the Wallet group
6. **WalletType(string)**  
7. **Currency(string)** your Currency create in wallet
8. **CurrencyType(string)**  currency type

```go
walletListPayload := &stratumsdk.WalletsListPayload{
	WalletEid        : 100
	WalletBalanceMin : 10
	WalletBalanceMax : 100
	WalletGroupEid   : 10
	WalletGroupId    : 12
	WalletType       : "checking"
	Currency         : "BTC"
	CurrencyType     : "crypto"
}
wallets, apiErr, err := client.Wallets().List(walletListPayload)
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

