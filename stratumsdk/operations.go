package stratumsdk

import (
	"encoding/json"
	"fmt"
)

type Operations struct {
	client *ApiClient
}

type FeeData struct {
	Currency      string  `json:"currency"`
	DestType      string  `json:"dest_type"`
	OperationType string  `json:"operation_type"`
	OperationFee  float64 `json:"operation_fee"`
}

type OperationData struct {
	DestType           string `json:"dest_type"`      // types: in,out,intra
	DirectionType      string `json:"direction_type"` //types in,out,intra
	OperationEid       int    `json:"operation_eid"`
	OperationStatus    string `json:"operation_status"`      // new,processing,done,failed
	OperationTsFrom    int    `json:"operation_ts_from"`     //  doubt ask for Sven
	OperationTsTo      int    `json:"operation_ts_to"`       // doubt ask for Sven
	OperationType      string `json:"operation_type"`        // types: deposit,withdraw,transfer"
	OperationUpdTsFrom int    `json:"operation_upd_ts_from"` // doubt ask for Sven
	OperationUpdTsTo   int    `json:"operation_upd_ts_to"`   // doubt ask for Sven
	WalletEid          int    `json:"wallet_eid"`
	WalletGroupEid     int    `json:"wallet_group_eid"`
	WalletId           int    `json:"wallet_id"`
}
type OperationPayload struct {
	DestType           string `json:"dest_type,omitempty"`      // types: in,out,intra
	DirectionType      string `json:"direction_type,omitempty"` //types in,out,intra
	OperationEid       int    `json:"operation_eid,omitempty"`
	OperationStatus    string `json:"operation_status,omitempty"`      // new,processing,done,failed
	OperationTsFrom    int    `json:"operation_ts_from,omitempty"`     //  doubt ask for Sven
	OperationTsTo      int    `json:"operation_ts_to,omitempty"`       // doubt ask for Sven
	OperationType      string `json:"operation_type,omitempty"`        // types: deposit,withdraw,transfer"
	OperationUpdTsFrom int    `json:"operation_upd_ts_from,omitempty"` // doubt ask for Sven
	OperationUpdTsTo   int    `json:"operation_upd_ts_to,omitempty"`   // doubt ask for Sven
	WalletEid          int    `json:"wallet_eid,omitempty"`
	WalletGroupEid     int    `json:"wallet_group_eid,omitempty"`
	WalletId           int    `json:"wallet_id,omitempty"`
}

type OperationResult struct {
	Data []OperationData `json:"data"`
}

type FeePayload struct {
	Currency      string `json:"currency"`
	DestType      string `json:"dest_type"`
	OperationType string `json:"operation_type"`
}

type FeeResult struct {
	Data []FeeData `json:"data"`
}

func (c *ApiClient) Operations() *Operations {
	return &Operations{client: c}
}

// Fees - get fees
func (o *Operations) Fees(payload *FeePayload) (*[]FeeData, *ApiError, error) {
	result := new(FeeResult)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	apiErr, err := o.client.call("operations", "fees", payloadJSON, &result)
	if err != nil {
		return nil, nil, err
	}
	return &result.Data, apiErr, nil

}

// List - list opearation in account
func (o *Operations) List(payload *OperationPayload) (*[]OperationData, *ApiError, error) {
	result := new(OperationResult)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("%s", payloadJSON)
	apiErr, err := o.client.call("operations", "list", payloadJSON, &result)
	if err != nil {
		return nil, nil, err
	}
	return &result.Data, apiErr, nil

}
