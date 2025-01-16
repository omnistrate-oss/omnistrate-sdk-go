# RemoveComputeInstanceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Id** | **string** | ID of a Compute Config | 
**InstanceType** | **string** | The instance type for this compute instance type config | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRemoveComputeInstanceTypeRequest

`func NewRemoveComputeInstanceTypeRequest(cloudProviderName string, id string, instanceType string, serviceId string, token string, ) *RemoveComputeInstanceTypeRequest`

NewRemoveComputeInstanceTypeRequest instantiates a new RemoveComputeInstanceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveComputeInstanceTypeRequestWithDefaults

`func NewRemoveComputeInstanceTypeRequestWithDefaults() *RemoveComputeInstanceTypeRequest`

NewRemoveComputeInstanceTypeRequestWithDefaults instantiates a new RemoveComputeInstanceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *RemoveComputeInstanceTypeRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *RemoveComputeInstanceTypeRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *RemoveComputeInstanceTypeRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetId

`func (o *RemoveComputeInstanceTypeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoveComputeInstanceTypeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoveComputeInstanceTypeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceType

`func (o *RemoveComputeInstanceTypeRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *RemoveComputeInstanceTypeRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *RemoveComputeInstanceTypeRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetServiceId

`func (o *RemoveComputeInstanceTypeRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RemoveComputeInstanceTypeRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RemoveComputeInstanceTypeRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *RemoveComputeInstanceTypeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RemoveComputeInstanceTypeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RemoveComputeInstanceTypeRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


