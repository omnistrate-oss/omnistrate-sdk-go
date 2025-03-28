# ReleaseImageConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of an Image Config | 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewReleaseImageConfigRequest

`func NewReleaseImageConfigRequest(id string, serviceId string, token string, ) *ReleaseImageConfigRequest`

NewReleaseImageConfigRequest instantiates a new ReleaseImageConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseImageConfigRequestWithDefaults

`func NewReleaseImageConfigRequestWithDefaults() *ReleaseImageConfigRequest`

NewReleaseImageConfigRequestWithDefaults instantiates a new ReleaseImageConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseImageConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseImageConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseImageConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProductTierId

`func (o *ReleaseImageConfigRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ReleaseImageConfigRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ReleaseImageConfigRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ReleaseImageConfigRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetServiceId

`func (o *ReleaseImageConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ReleaseImageConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ReleaseImageConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ReleaseImageConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ReleaseImageConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ReleaseImageConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


