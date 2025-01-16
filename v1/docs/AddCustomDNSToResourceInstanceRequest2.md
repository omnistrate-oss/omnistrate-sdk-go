# AddCustomDNSToResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDNS** | **string** | The custom DNS to add | 
**TargetPort** | Pointer to **int64** | The target port | [optional] 

## Methods

### NewAddCustomDNSToResourceInstanceRequest2

`func NewAddCustomDNSToResourceInstanceRequest2(customDNS string, ) *AddCustomDNSToResourceInstanceRequest2`

NewAddCustomDNSToResourceInstanceRequest2 instantiates a new AddCustomDNSToResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCustomDNSToResourceInstanceRequest2WithDefaults

`func NewAddCustomDNSToResourceInstanceRequest2WithDefaults() *AddCustomDNSToResourceInstanceRequest2`

NewAddCustomDNSToResourceInstanceRequest2WithDefaults instantiates a new AddCustomDNSToResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDNS

`func (o *AddCustomDNSToResourceInstanceRequest2) GetCustomDNS() string`

GetCustomDNS returns the CustomDNS field if non-nil, zero value otherwise.

### GetCustomDNSOk

`func (o *AddCustomDNSToResourceInstanceRequest2) GetCustomDNSOk() (*string, bool)`

GetCustomDNSOk returns a tuple with the CustomDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDNS

`func (o *AddCustomDNSToResourceInstanceRequest2) SetCustomDNS(v string)`

SetCustomDNS sets CustomDNS field to given value.


### GetTargetPort

`func (o *AddCustomDNSToResourceInstanceRequest2) GetTargetPort() int64`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *AddCustomDNSToResourceInstanceRequest2) GetTargetPortOk() (*int64, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *AddCustomDNSToResourceInstanceRequest2) SetTargetPort(v int64)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *AddCustomDNSToResourceInstanceRequest2) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


