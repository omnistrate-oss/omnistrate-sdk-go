package fleet

import (
	"encoding/json"
	"testing"
)

// minimalResourceInstanceJSON returns a JSON object containing only the required
// fields for a ResourceInstance — notably WITHOUT input_params.
func minimalResourceInstanceJSON() string {
	return `{
		"adopted": false,
		"cloudProvider": "aws",
		"consumptionResourceInstanceResult": {
			"id": "inst-1",
			"status": "RUNNING"
		},
		"defaultSubscription": true,
		"environmentId": "env-1",
		"instanceDebugCommands": [],
		"integrationsStatus": [],
		"organizationId": "org-1",
		"organizationName": "TestOrg",
		"productTierId": "pt-1",
		"productTierName": "Free",
		"productTierType": "OMNISTRATE_DEDICATED",
		"resourceVersionSummaries": [],
		"serviceEnvName": "prod",
		"serviceId": "svc-1",
		"serviceModelId": "sm-1",
		"serviceModelName": "Model",
		"serviceModelType": "CUSTOMER_HOSTED",
		"serviceName": "TestService",
		"subscriptionId": "sub-1",
		"subscriptionOwnerName": "owner",
		"tierVersion": "1.0",
		"tierVersionReleasedAt": "2025-01-01T00:00:00Z",
		"tierVersionReleasedByUserId": "user-1",
		"tierVersionReleasedByUserName": "admin",
		"tierVersionStatus": "Preferred"
	}`
}

func TestResourceInstance_UnmarshalJSON_WithoutInputParams(t *testing.T) {
	data := []byte(minimalResourceInstanceJSON())
	var ri ResourceInstance
	err := json.Unmarshal(data, &ri)
	if err != nil {
		t.Fatalf("UnmarshalJSON should succeed without input_params, got: %v", err)
	}
	if ri.InputParams != nil {
		t.Errorf("expected InputParams to be nil, got: %v", ri.InputParams)
	}
	if ri.ServiceId != "svc-1" {
		t.Errorf("expected ServiceId 'svc-1', got: %s", ri.ServiceId)
	}
}

func TestResourceInstance_UnmarshalJSON_WithInputParams(t *testing.T) {
	data := []byte(minimalResourceInstanceJSON())

	// Inject input_params into the JSON
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	raw["input_params"] = map[string]interface{}{"key1": "val1"}
	data, _ = json.Marshal(raw)

	var ri ResourceInstance
	err := json.Unmarshal(data, &ri)
	if err != nil {
		t.Fatalf("UnmarshalJSON should succeed with input_params, got: %v", err)
	}
	if ri.InputParams == nil {
		t.Fatal("expected InputParams to be non-nil")
	}
}

func TestResourceInstance_UnmarshalJSON_MissingRequiredField(t *testing.T) {
	// Remove a truly required field (serviceId)
	data := []byte(minimalResourceInstanceJSON())
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	delete(raw, "serviceId")
	data, _ = json.Marshal(raw)

	var ri ResourceInstance
	err := json.Unmarshal(data, &ri)
	if err == nil {
		t.Fatal("expected error when required field 'serviceId' is missing")
	}
}
