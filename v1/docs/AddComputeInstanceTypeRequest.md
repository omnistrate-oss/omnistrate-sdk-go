# AddComputeInstanceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**ConfigOverride** | Pointer to [**ComputeInstanceTypeConfigOverride**](ComputeInstanceTypeConfigOverride.md) |  | [optional] 
**Id** | **string** | ID of a Compute Config | 
**InstanceType** | **string** | The instance type for this compute instance type config | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewAddComputeInstanceTypeRequest

`func NewAddComputeInstanceTypeRequest(cloudProviderName string, id string, instanceType string, serviceId string, token string, ) *AddComputeInstanceTypeRequest`

NewAddComputeInstanceTypeRequest instantiates a new AddComputeInstanceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddComputeInstanceTypeRequestWithDefaults

`func NewAddComputeInstanceTypeRequestWithDefaults() *AddComputeInstanceTypeRequest`

NewAddComputeInstanceTypeRequestWithDefaults instantiates a new AddComputeInstanceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *AddComputeInstanceTypeRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *AddComputeInstanceTypeRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *AddComputeInstanceTypeRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetConfigOverride

`func (o *AddComputeInstanceTypeRequest) GetConfigOverride() ComputeInstanceTypeConfigOverride`

GetConfigOverride returns the ConfigOverride field if non-nil, zero value otherwise.

### GetConfigOverrideOk

`func (o *AddComputeInstanceTypeRequest) GetConfigOverrideOk() (*ComputeInstanceTypeConfigOverride, bool)`

GetConfigOverrideOk returns a tuple with the ConfigOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOverride

`func (o *AddComputeInstanceTypeRequest) SetConfigOverride(v ComputeInstanceTypeConfigOverride)`

SetConfigOverride sets ConfigOverride field to given value.

### HasConfigOverride

`func (o *AddComputeInstanceTypeRequest) HasConfigOverride() bool`

HasConfigOverride returns a boolean if a field has been set.

### GetId

`func (o *AddComputeInstanceTypeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddComputeInstanceTypeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddComputeInstanceTypeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceType

`func (o *AddComputeInstanceTypeRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AddComputeInstanceTypeRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AddComputeInstanceTypeRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetServiceId

`func (o *AddComputeInstanceTypeRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AddComputeInstanceTypeRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AddComputeInstanceTypeRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *AddComputeInstanceTypeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddComputeInstanceTypeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddComputeInstanceTypeRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


