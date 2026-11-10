# ReassignMarketplaceFulfillmentRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | The existing customer organization to receive the confirmed contract and its own subscription. The subscription id, deployments and metering activation state are preserved. A repeat request for the same organization is idempotent | 

## Methods

### NewReassignMarketplaceFulfillmentRequest2

`func NewReassignMarketplaceFulfillmentRequest2(orgId string, ) *ReassignMarketplaceFulfillmentRequest2`

NewReassignMarketplaceFulfillmentRequest2 instantiates a new ReassignMarketplaceFulfillmentRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReassignMarketplaceFulfillmentRequest2WithDefaults

`func NewReassignMarketplaceFulfillmentRequest2WithDefaults() *ReassignMarketplaceFulfillmentRequest2`

NewReassignMarketplaceFulfillmentRequest2WithDefaults instantiates a new ReassignMarketplaceFulfillmentRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ReassignMarketplaceFulfillmentRequest2) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ReassignMarketplaceFulfillmentRequest2) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ReassignMarketplaceFulfillmentRequest2) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


