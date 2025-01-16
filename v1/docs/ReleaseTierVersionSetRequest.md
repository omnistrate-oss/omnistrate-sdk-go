# ReleaseTierVersionSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPreferred** | Pointer to **bool** | Indicates whether this version set is preferred. | [optional] 
**Name** | Pointer to **string** | The name of the product-tier version set. | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Version** | **string** | The version number for the specific version set. | 

## Methods

### NewReleaseTierVersionSetRequest

`func NewReleaseTierVersionSetRequest(productTierId string, serviceId string, token string, version string, ) *ReleaseTierVersionSetRequest`

NewReleaseTierVersionSetRequest instantiates a new ReleaseTierVersionSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseTierVersionSetRequestWithDefaults

`func NewReleaseTierVersionSetRequestWithDefaults() *ReleaseTierVersionSetRequest`

NewReleaseTierVersionSetRequestWithDefaults instantiates a new ReleaseTierVersionSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPreferred

`func (o *ReleaseTierVersionSetRequest) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *ReleaseTierVersionSetRequest) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *ReleaseTierVersionSetRequest) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *ReleaseTierVersionSetRequest) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetName

`func (o *ReleaseTierVersionSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleaseTierVersionSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleaseTierVersionSetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReleaseTierVersionSetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductTierId

`func (o *ReleaseTierVersionSetRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ReleaseTierVersionSetRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ReleaseTierVersionSetRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *ReleaseTierVersionSetRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ReleaseTierVersionSetRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ReleaseTierVersionSetRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ReleaseTierVersionSetRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ReleaseTierVersionSetRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ReleaseTierVersionSetRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVersion

`func (o *ReleaseTierVersionSetRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReleaseTierVersionSetRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReleaseTierVersionSetRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


