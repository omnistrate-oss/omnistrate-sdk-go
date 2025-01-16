# DescribeFleetFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | **string** | FleetFeatureType is to enable / disable features per service provider | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeFleetFeatureRequest

`func NewDescribeFleetFeatureRequest(feature string, token string, ) *DescribeFleetFeatureRequest`

NewDescribeFleetFeatureRequest instantiates a new DescribeFleetFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeFleetFeatureRequestWithDefaults

`func NewDescribeFleetFeatureRequestWithDefaults() *DescribeFleetFeatureRequest`

NewDescribeFleetFeatureRequestWithDefaults instantiates a new DescribeFleetFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *DescribeFleetFeatureRequest) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DescribeFleetFeatureRequest) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DescribeFleetFeatureRequest) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetToken

`func (o *DescribeFleetFeatureRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeFleetFeatureRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeFleetFeatureRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


