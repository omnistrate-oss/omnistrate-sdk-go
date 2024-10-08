# UpdateServiceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the service | [optional] 
**Name** | Pointer to **string** | Name of the Service | [optional] 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 

## Methods

### NewUpdateServiceRequestBody

`func NewUpdateServiceRequestBody() *UpdateServiceRequestBody`

NewUpdateServiceRequestBody instantiates a new UpdateServiceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceRequestBodyWithDefaults

`func NewUpdateServiceRequestBodyWithDefaults() *UpdateServiceRequestBody`

NewUpdateServiceRequestBodyWithDefaults instantiates a new UpdateServiceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateServiceRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServiceRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServiceRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServiceRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateServiceRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServiceRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServiceRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServiceRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceLogoURL

`func (o *UpdateServiceRequestBody) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *UpdateServiceRequestBody) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *UpdateServiceRequestBody) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *UpdateServiceRequestBody) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


