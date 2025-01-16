# DisableFleetFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | Pointer to **string** | FleetFeatureType is to enable / disable features per service provider | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDisableFleetFeatureRequest

`func NewDisableFleetFeatureRequest(token string, ) *DisableFleetFeatureRequest`

NewDisableFleetFeatureRequest instantiates a new DisableFleetFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableFleetFeatureRequestWithDefaults

`func NewDisableFleetFeatureRequestWithDefaults() *DisableFleetFeatureRequest`

NewDisableFleetFeatureRequestWithDefaults instantiates a new DisableFleetFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *DisableFleetFeatureRequest) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DisableFleetFeatureRequest) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DisableFleetFeatureRequest) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *DisableFleetFeatureRequest) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetToken

`func (o *DisableFleetFeatureRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DisableFleetFeatureRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DisableFleetFeatureRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


