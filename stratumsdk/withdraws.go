package stratumsdk

import (
	"encoding/json"
)

// Withdraws strcture for consumer widthdraws
type Withdraws struct {
	client *apiClient
}

type WithdrawsData struct {
	DestAddress     string  `json:"dest_address"`
	OperationAmount float64 `json:"operation_amount"`
	OperationDesc   string  `json:"operation_amount"`
	OperationEid    int     `json:"operation_eid"`
	OperationOtp    int     `json:"operation_otp"`
	WalletId        int     `json:"walletId"`
}
type WithdrawsPayload struct {
	DestAddress     string  `json:"dest_address"`
	OperationAmount float64 `json:"operation_amount"`
	OperationDesc   string  `json:"operation_amount"`
	OperationEid    int     `json:"operation_eid"`
	OperationOtp    int     `json:"operation_otp"`
	WalletId        int     `json:"walletId"`
}

type WithdrawsResult struct {
	Data WithdrawsData `json:"data"`
}

func (c *apiClient) Withdraws() *Withdraws {
	return &Withdraws{client: c}
}

func (w *Withdraws) Crypto(payload *WithdrawsPayload) (*WithdrawsData, *ApiError, error) {
	result := new(WithdrawsResult)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	apiErr, err := w.client.call("withdraws", "crypto", payloadJSON, &result)
	if err != nil {
		return nil, nil, err
	}
	return &result.Data, apiErr, nil
}
