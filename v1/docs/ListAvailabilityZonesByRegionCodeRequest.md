# ListAvailabilityZonesByRegionCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**RegionCode** | **string** |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAvailabilityZonesByRegionCodeRequest

`func NewListAvailabilityZonesByRegionCodeRequest(cloudProviderName string, regionCode string, token string, ) *ListAvailabilityZonesByRegionCodeRequest`

NewListAvailabilityZonesByRegionCodeRequest instantiates a new ListAvailabilityZonesByRegionCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAvailabilityZonesByRegionCodeRequestWithDefaults

`func NewListAvailabilityZonesByRegionCodeRequestWithDefaults() *ListAvailabilityZonesByRegionCodeRequest`

NewListAvailabilityZonesByRegionCodeRequestWithDefaults instantiates a new ListAvailabilityZonesByRegionCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *ListAvailabilityZonesByRegionCodeRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *ListAvailabilityZonesByRegionCodeRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *ListAvailabilityZonesByRegionCodeRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetRegionCode

`func (o *ListAvailabilityZonesByRegionCodeRequest) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ListAvailabilityZonesByRegionCodeRequest) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ListAvailabilityZonesByRegionCodeRequest) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetToken

`func (o *ListAvailabilityZonesByRegionCodeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAvailabilityZonesByRegionCodeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAvailabilityZonesByRegionCodeRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


