# ListVUnitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the Infra Provider | 
**Region** | **string** | Region code specific to the cloud-provider | 
**ServiceId** | **string** | ID of a Service | 
**ServiceModelId** | **string** | ID of a Service Model | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListVUnitsRequest

`func NewListVUnitsRequest(cloudProvider string, region string, serviceId string, serviceModelId string, token string, ) *ListVUnitsRequest`

NewListVUnitsRequest instantiates a new ListVUnitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVUnitsRequestWithDefaults

`func NewListVUnitsRequestWithDefaults() *ListVUnitsRequest`

NewListVUnitsRequestWithDefaults instantiates a new ListVUnitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ListVUnitsRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ListVUnitsRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ListVUnitsRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *ListVUnitsRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ListVUnitsRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ListVUnitsRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetServiceId

`func (o *ListVUnitsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListVUnitsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListVUnitsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *ListVUnitsRequest) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *ListVUnitsRequest) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *ListVUnitsRequest) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetToken

`func (o *ListVUnitsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListVUnitsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListVUnitsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


