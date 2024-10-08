# AddComputeInstanceTypeRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | The cloud provider for this compute instance type config | 
**InstanceType** | **string** | The instance type for this compute instance type config | 

## Methods

### NewAddComputeInstanceTypeRequestBody

`func NewAddComputeInstanceTypeRequestBody(cloudProviderName string, instanceType string, ) *AddComputeInstanceTypeRequestBody`

NewAddComputeInstanceTypeRequestBody instantiates a new AddComputeInstanceTypeRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddComputeInstanceTypeRequestBodyWithDefaults

`func NewAddComputeInstanceTypeRequestBodyWithDefaults() *AddComputeInstanceTypeRequestBody`

NewAddComputeInstanceTypeRequestBodyWithDefaults instantiates a new AddComputeInstanceTypeRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *AddComputeInstanceTypeRequestBody) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *AddComputeInstanceTypeRequestBody) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *AddComputeInstanceTypeRequestBody) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetInstanceType

`func (o *AddComputeInstanceTypeRequestBody) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AddComputeInstanceTypeRequestBody) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AddComputeInstanceTypeRequestBody) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


