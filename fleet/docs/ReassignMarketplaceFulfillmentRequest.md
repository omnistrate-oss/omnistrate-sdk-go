# ReassignMarketplaceFulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The confirmed Omnistrate contract identifier | 
**OrgId** | **string** | The existing customer organization to receive the confirmed contract and its own subscription. The subscription id, deployments and metering activation state are preserved. A repeat request for the same organization is idempotent | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewReassignMarketplaceFulfillmentRequest

`func NewReassignMarketplaceFulfillmentRequest(id string, orgId string, token string, ) *ReassignMarketplaceFulfillmentRequest`

NewReassignMarketplaceFulfillmentRequest instantiates a new ReassignMarketplaceFulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReassignMarketplaceFulfillmentRequestWithDefaults

`func NewReassignMarketplaceFulfillmentRequestWithDefaults() *ReassignMarketplaceFulfillmentRequest`

NewReassignMarketplaceFulfillmentRequestWithDefaults instantiates a new ReassignMarketplaceFulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReassignMarketplaceFulfillmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReassignMarketplaceFulfillmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReassignMarketplaceFulfillmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *ReassignMarketplaceFulfillmentRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ReassignMarketplaceFulfillmentRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ReassignMarketplaceFulfillmentRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetToken

`func (o *ReassignMarketplaceFulfillmentRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ReassignMarketplaceFulfillmentRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ReassignMarketplaceFulfillmentRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


