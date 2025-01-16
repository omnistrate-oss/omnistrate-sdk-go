# UpdateLimitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the limit | [optional] 
**Family** | **string** | The limit family | 
**Key** | **string** | Unique key to identify the limit | 
**Name** | Pointer to **string** | Name of the limit | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Value** | **int64** | Value of the limit being enforced | 

## Methods

### NewUpdateLimitRequest

`func NewUpdateLimitRequest(family string, key string, token string, value int64, ) *UpdateLimitRequest`

NewUpdateLimitRequest instantiates a new UpdateLimitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLimitRequestWithDefaults

`func NewUpdateLimitRequestWithDefaults() *UpdateLimitRequest`

NewUpdateLimitRequestWithDefaults instantiates a new UpdateLimitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateLimitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLimitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLimitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLimitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFamily

`func (o *UpdateLimitRequest) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *UpdateLimitRequest) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *UpdateLimitRequest) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetKey

`func (o *UpdateLimitRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateLimitRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateLimitRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *UpdateLimitRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLimitRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLimitRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLimitRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateLimitRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateLimitRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateLimitRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetValue

`func (o *UpdateLimitRequest) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateLimitRequest) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateLimitRequest) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


