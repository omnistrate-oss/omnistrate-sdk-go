# DescribeResourceInstanceInstallerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallerURL** | Pointer to **string** | URL from where the resource instance installer can be downloaded | [optional] 
**InstanceId** | **string** | ID of a Resource Instance | 

## Methods

### NewDescribeResourceInstanceInstallerResult

`func NewDescribeResourceInstanceInstallerResult(instanceId string, ) *DescribeResourceInstanceInstallerResult`

NewDescribeResourceInstanceInstallerResult instantiates a new DescribeResourceInstanceInstallerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceInstanceInstallerResultWithDefaults

`func NewDescribeResourceInstanceInstallerResultWithDefaults() *DescribeResourceInstanceInstallerResult`

NewDescribeResourceInstanceInstallerResultWithDefaults instantiates a new DescribeResourceInstanceInstallerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallerURL

`func (o *DescribeResourceInstanceInstallerResult) GetInstallerURL() string`

GetInstallerURL returns the InstallerURL field if non-nil, zero value otherwise.

### GetInstallerURLOk

`func (o *DescribeResourceInstanceInstallerResult) GetInstallerURLOk() (*string, bool)`

GetInstallerURLOk returns a tuple with the InstallerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerURL

`func (o *DescribeResourceInstanceInstallerResult) SetInstallerURL(v string)`

SetInstallerURL sets InstallerURL field to given value.

### HasInstallerURL

`func (o *DescribeResourceInstanceInstallerResult) HasInstallerURL() bool`

HasInstallerURL returns a boolean if a field has been set.

### GetInstanceId

`func (o *DescribeResourceInstanceInstallerResult) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DescribeResourceInstanceInstallerResult) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DescribeResourceInstanceInstallerResult) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


