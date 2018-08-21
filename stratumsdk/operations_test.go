package stratumsdk_test

import (
	"testing"

	"github.com/rafaeltokyo/stratum-sdk-go/stratumsdk"
)

func TestOperationFees(t *testing.T) {
	client := stratumsdk.Initial(apiUser, apiSecret, true)
	feePayload := new(stratumsdk.FeePayload)
	_, apiErr, err := client.Operations().Fees(feePayload)
	if err != nil {
		t.Errorf("sdk error: TestOperationFees() -> %s ", err.Error())
	}
	if apiErr != nil {
		t.Errorf("apiError: TestOperationFees() -> %s ", apiErr.Data)
	}
}

func TestOperationList(t *testing.T) {
	client := stratumsdk.Initial(apiUser, apiSecret, true)
	feePayload := new(stratumsdk.OperationPayload)
	_, apiErr, err := client.Operations().List(feePayload)
	if err != nil {
		t.Errorf("sdk error: TestOperationList() -> %s ", err.Error())
	}
	if apiErr != nil {
		t.Errorf("apiError: TestOperationList() -> %s ", apiErr.Data)
	}
}
