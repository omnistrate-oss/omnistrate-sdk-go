# DisableServiceModelFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | **string** | Name of the service model feature | 
**Id** | **string** | ID of a Service Model | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDisableServiceModelFeatureRequest

`func NewDisableServiceModelFeatureRequest(feature string, id string, serviceId string, token string, ) *DisableServiceModelFeatureRequest`

NewDisableServiceModelFeatureRequest instantiates a new DisableServiceModelFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableServiceModelFeatureRequestWithDefaults

`func NewDisableServiceModelFeatureRequestWithDefaults() *DisableServiceModelFeatureRequest`

NewDisableServiceModelFeatureRequestWithDefaults instantiates a new DisableServiceModelFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *DisableServiceModelFeatureRequest) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DisableServiceModelFeatureRequest) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DisableServiceModelFeatureRequest) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetId

`func (o *DisableServiceModelFeatureRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisableServiceModelFeatureRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisableServiceModelFeatureRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *DisableServiceModelFeatureRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DisableServiceModelFeatureRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DisableServiceModelFeatureRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DisableServiceModelFeatureRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisableServiceModelFeatureRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisableServiceModelFeatureRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


