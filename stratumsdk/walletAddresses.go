package stratumsdk

import (
	"encoding/json"
)

type WalletAddressData struct {
	WalletAddressEid   int    `json:"wallet_address_eid"`
	WalletId           int64  `json:"wallet_id"`
	WalletAddressLabel string `json:"wallet_address_label"`
	WalletAddress      string `json:"wallet_address"`
}

type WalletsAddressesListPayload struct {
	WalletAddressEid string `json:"wallet_address_eid"`
	WalletId         int64  `json:"wallet_id"`
	WalletEid        int64  `json:"wallet_eid"`
	Currency         string `json:"currency"`
}

type WalletAddressesPayload struct {
	WalletAddressEid   int    `json:"wallet_address_eid"`
	WalletId           int64  `json:"wallet_id"`
	WalletAddressLabel string `json:"wallet_address_label"`
	WalletAddress      string `json:"wallet_address"`
	Currency           string `json:"currency"`
}

// Success reply from a list action
type WalletsAddressesListResult struct {
	Data []WalletAddressData `json:"data"`
}
type WalletAddressesResult struct {
	Data WalletAddressData `json:"data"`
}

// WalletsAddresses is a structure create a structure
type WalletsAddresses struct {
	client *apiClient
}

//WalletsAddresses - Create a new instance struct
func (c *apiClient) WalletsAddresses() *WalletsAddresses {
	return &WalletsAddresses{client: c}
}

//List - list wallet addresses assign
func (wa *WalletsAddresses) List(payload *WalletsAddressesListPayload) (*[]WalletAddressData, *ApiError, error) {
	result := new(WalletsAddressesListResult)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	apiErr, err := wa.client.call("wallets", "list", payloadJSON, result)
	if err != nil {
		return nil, nil, err
	}
	return &result.Data, apiErr, nil
}

//Assign - assign the wallet by WalletId
func (wa *WalletsAddresses) Assign(payload *WalletAddressesPayload) (*WalletAddressData, *ApiError, error) {
	result := new(WalletAddressesResult)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	apiErr, err := wa.client.call("walletAddresses", "assign", payloadJSON, result)
	if err != nil {
		return nil, nil, err
	}
	return &result.Data, apiErr, nil

}

func (wa *WalletsAddresses) Get(payload *WalletAddressesPayload) (*WalletAddressData, *ApiError, error) {
	result := new(WalletAddressesResult)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	apiErr, err := wa.client.call("walletAddresses", "get", payloadJSON, result)
	if err != nil {
		return nil, nil, err
	}
	return &result.Data, apiErr, nil

}
