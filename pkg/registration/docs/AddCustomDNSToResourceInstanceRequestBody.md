# AddCustomDNSToResourceInstanceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDNS** | **string** | The custom DNS to add | 
**TargetPort** | Pointer to **int64** | The target port | [optional] 

## Methods

### NewAddCustomDNSToResourceInstanceRequestBody

`func NewAddCustomDNSToResourceInstanceRequestBody(customDNS string, ) *AddCustomDNSToResourceInstanceRequestBody`

NewAddCustomDNSToResourceInstanceRequestBody instantiates a new AddCustomDNSToResourceInstanceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCustomDNSToResourceInstanceRequestBodyWithDefaults

`func NewAddCustomDNSToResourceInstanceRequestBodyWithDefaults() *AddCustomDNSToResourceInstanceRequestBody`

NewAddCustomDNSToResourceInstanceRequestBodyWithDefaults instantiates a new AddCustomDNSToResourceInstanceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDNS

`func (o *AddCustomDNSToResourceInstanceRequestBody) GetCustomDNS() string`

GetCustomDNS returns the CustomDNS field if non-nil, zero value otherwise.

### GetCustomDNSOk

`func (o *AddCustomDNSToResourceInstanceRequestBody) GetCustomDNSOk() (*string, bool)`

GetCustomDNSOk returns a tuple with the CustomDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDNS

`func (o *AddCustomDNSToResourceInstanceRequestBody) SetCustomDNS(v string)`

SetCustomDNS sets CustomDNS field to given value.


### GetTargetPort

`func (o *AddCustomDNSToResourceInstanceRequestBody) GetTargetPort() int64`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *AddCustomDNSToResourceInstanceRequestBody) GetTargetPortOk() (*int64, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *AddCustomDNSToResourceInstanceRequestBody) SetTargetPort(v int64)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *AddCustomDNSToResourceInstanceRequestBody) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


