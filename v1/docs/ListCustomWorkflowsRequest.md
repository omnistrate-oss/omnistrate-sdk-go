# ListCustomWorkflowsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierVersion** | Pointer to **string** | The product tier version to list. If not specified, the latest version is described. | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListCustomWorkflowsRequest

`func NewListCustomWorkflowsRequest(productTierId string, serviceId string, token string, ) *ListCustomWorkflowsRequest`

NewListCustomWorkflowsRequest instantiates a new ListCustomWorkflowsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomWorkflowsRequestWithDefaults

`func NewListCustomWorkflowsRequestWithDefaults() *ListCustomWorkflowsRequest`

NewListCustomWorkflowsRequestWithDefaults instantiates a new ListCustomWorkflowsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *ListCustomWorkflowsRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListCustomWorkflowsRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListCustomWorkflowsRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierVersion

`func (o *ListCustomWorkflowsRequest) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *ListCustomWorkflowsRequest) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *ListCustomWorkflowsRequest) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *ListCustomWorkflowsRequest) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetServiceId

`func (o *ListCustomWorkflowsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListCustomWorkflowsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListCustomWorkflowsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ListCustomWorkflowsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListCustomWorkflowsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListCustomWorkflowsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


