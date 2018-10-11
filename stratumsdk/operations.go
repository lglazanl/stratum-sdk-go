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
	WalletId               int     `json:"wallet_id"`
	OperationId            int     `json:"operation_id"`
	OperationAmount        float64 `json:"operation_amount"`
	OperationTotalAmount   float64 `json:"operation_tamount"`
	OperationFee           float64 `json:"operation_fee"`
	OperationDescription   string  `json:"operation_desc"`
	OperationExternalId    int     `json:"operation_eid"`
	OperationExternalTXId  string  `json:"operation_etxid"`
	OperationTs            int     `json:"operation_ts"`
	OperationUpdatedTs     int     `json:"operation_upd_ts"`
	OperationConfirmations int     `json:"operation_conf"`
	OperationConfRequired  int     `json:"operation_confreq"`
	DestinationTypeData    string  `json:"dest_type_data"`
	OperationInfo          string  `json:"operation_info"`
	OperationStatus        string  `json:"operation_status"` //"in:new,processing,done,failed"
	OperationType          string  `json:"operation_type"`   //"in:deposit,withdraw,transfer"
	DirectionType          string  `json:"direction_type"`   //"in:in,out,intra"
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
	OperationId        int    `json:"operation_id"`
}

type OperationResult struct {
	Data OperationData `json:"data"`
}

type OperationListResult struct {
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

// List - list operations in account
func (o *Operations) List(payload *OperationPayload) (*[]OperationData, *ApiError, error) {
	result := new(OperationListResult)
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

// Get - get operation in account
func (o *Operations) Get(payload *OperationPayload) (*OperationData, *ApiError, error) {
	result := new(OperationResult)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("%s", payloadJSON)
	apiErr, err := o.client.call("operations", "get", payloadJSON, &result)
	if err != nil {
		return nil, nil, err
	}
	return &result.Data, apiErr, nil

}
