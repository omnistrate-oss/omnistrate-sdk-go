/*
Omnistrate Registration API

Testing PromoteServiceEnvironmentRequest2

*/

package v1

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func TestPromoteServiceEnvironmentRequest2ProductTierVersionMarshalJSON(t *testing.T) {
	request := openapiclient.NewPromoteServiceEnvironmentRequest2()
	request.SetProductTierId("pt-123")
	request.SetProductTierVersion("1.2.3")

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	var body map[string]string
	if err := json.Unmarshal(data, &body); err != nil {
		t.Fatalf("unmarshal marshaled request: %v", err)
	}

	if got := body["productTierVersion"]; got != "1.2.3" {
		t.Fatalf("productTierVersion = %q, want %q", got, "1.2.3")
	}
	if got := body["productTierId"]; got != "pt-123" {
		t.Fatalf("productTierId = %q, want %q", got, "pt-123")
	}
}

func TestPromoteServiceEnvironmentRequest2ProductTierVersionOmittedWhenNil(t *testing.T) {
	request := openapiclient.NewPromoteServiceEnvironmentRequest2()
	request.SetProductTierId("pt-123")

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		t.Fatalf("unmarshal marshaled request: %v", err)
	}

	if _, ok := body["productTierVersion"]; ok {
		t.Fatalf("productTierVersion was serialized when nil: %s", string(data))
	}
}

func TestPromoteServiceEnvironmentRequest2ProductTierVersionUnmarshalJSON(t *testing.T) {
	var request openapiclient.PromoteServiceEnvironmentRequest2
	if err := json.Unmarshal([]byte(`{"productTierId":"pt-123","productTierVersion":"1.2.3","extra":"value"}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}

	if got := request.GetProductTierVersion(); got != "1.2.3" {
		t.Fatalf("ProductTierVersion = %q, want %q", got, "1.2.3")
	}
	if _, ok := request.AdditionalProperties["productTierVersion"]; ok {
		t.Fatal("productTierVersion should not be stored in AdditionalProperties")
	}
	if got := request.AdditionalProperties["extra"]; got != "value" {
		t.Fatalf("extra additional property = %#v, want %q", got, "value")
	}
}
