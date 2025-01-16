# DeleteUpgradePathRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of an Upgrade Path | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDeleteUpgradePathRequest

`func NewDeleteUpgradePathRequest(token string, ) *DeleteUpgradePathRequest`

NewDeleteUpgradePathRequest instantiates a new DeleteUpgradePathRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUpgradePathRequestWithDefaults

`func NewDeleteUpgradePathRequestWithDefaults() *DeleteUpgradePathRequest`

NewDeleteUpgradePathRequestWithDefaults instantiates a new DeleteUpgradePathRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteUpgradePathRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteUpgradePathRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteUpgradePathRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeleteUpgradePathRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceId

`func (o *DeleteUpgradePathRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteUpgradePathRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteUpgradePathRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeleteUpgradePathRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetToken

`func (o *DeleteUpgradePathRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteUpgradePathRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteUpgradePathRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


