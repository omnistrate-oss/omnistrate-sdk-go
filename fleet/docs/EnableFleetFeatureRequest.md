# EnableFleetFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | Pointer to **string** | FleetFeatureType is to enable / disable features per service provider | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewEnableFleetFeatureRequest

`func NewEnableFleetFeatureRequest(token string, ) *EnableFleetFeatureRequest`

NewEnableFleetFeatureRequest instantiates a new EnableFleetFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableFleetFeatureRequestWithDefaults

`func NewEnableFleetFeatureRequestWithDefaults() *EnableFleetFeatureRequest`

NewEnableFleetFeatureRequestWithDefaults instantiates a new EnableFleetFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *EnableFleetFeatureRequest) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *EnableFleetFeatureRequest) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *EnableFleetFeatureRequest) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *EnableFleetFeatureRequest) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetToken

`func (o *EnableFleetFeatureRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EnableFleetFeatureRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EnableFleetFeatureRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


