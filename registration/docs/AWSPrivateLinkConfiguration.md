# AWSPrivateLinkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int64** | The port of the target group | 
**TargetGroupName** | **string** | The target group name | 

## Methods

### NewAWSPrivateLinkConfiguration

`func NewAWSPrivateLinkConfiguration(port int64, targetGroupName string, ) *AWSPrivateLinkConfiguration`

NewAWSPrivateLinkConfiguration instantiates a new AWSPrivateLinkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSPrivateLinkConfigurationWithDefaults

`func NewAWSPrivateLinkConfigurationWithDefaults() *AWSPrivateLinkConfiguration`

NewAWSPrivateLinkConfigurationWithDefaults instantiates a new AWSPrivateLinkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *AWSPrivateLinkConfiguration) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AWSPrivateLinkConfiguration) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AWSPrivateLinkConfiguration) SetPort(v int64)`

SetPort sets Port field to given value.


### GetTargetGroupName

`func (o *AWSPrivateLinkConfiguration) GetTargetGroupName() string`

GetTargetGroupName returns the TargetGroupName field if non-nil, zero value otherwise.

### GetTargetGroupNameOk

`func (o *AWSPrivateLinkConfiguration) GetTargetGroupNameOk() (*string, bool)`

GetTargetGroupNameOk returns a tuple with the TargetGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupName

`func (o *AWSPrivateLinkConfiguration) SetTargetGroupName(v string)`

SetTargetGroupName sets TargetGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


