# HelmChartAffinityControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableInjection** | Pointer to **bool** | Enable auto injection of affinity rules | [optional] 
**EnableSharedHost** | Pointer to **bool** | Enable shared host support for the resource | [optional] 

## Methods

### NewHelmChartAffinityControl

`func NewHelmChartAffinityControl() *HelmChartAffinityControl`

NewHelmChartAffinityControl instantiates a new HelmChartAffinityControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartAffinityControlWithDefaults

`func NewHelmChartAffinityControlWithDefaults() *HelmChartAffinityControl`

NewHelmChartAffinityControlWithDefaults instantiates a new HelmChartAffinityControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableInjection

`func (o *HelmChartAffinityControl) GetEnableInjection() bool`

GetEnableInjection returns the EnableInjection field if non-nil, zero value otherwise.

### GetEnableInjectionOk

`func (o *HelmChartAffinityControl) GetEnableInjectionOk() (*bool, bool)`

GetEnableInjectionOk returns a tuple with the EnableInjection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInjection

`func (o *HelmChartAffinityControl) SetEnableInjection(v bool)`

SetEnableInjection sets EnableInjection field to given value.

### HasEnableInjection

`func (o *HelmChartAffinityControl) HasEnableInjection() bool`

HasEnableInjection returns a boolean if a field has been set.

### GetEnableSharedHost

`func (o *HelmChartAffinityControl) GetEnableSharedHost() bool`

GetEnableSharedHost returns the EnableSharedHost field if non-nil, zero value otherwise.

### GetEnableSharedHostOk

`func (o *HelmChartAffinityControl) GetEnableSharedHostOk() (*bool, bool)`

GetEnableSharedHostOk returns a tuple with the EnableSharedHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSharedHost

`func (o *HelmChartAffinityControl) SetEnableSharedHost(v bool)`

SetEnableSharedHost sets EnableSharedHost field to given value.

### HasEnableSharedHost

`func (o *HelmChartAffinityControl) HasEnableSharedHost() bool`

HasEnableSharedHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


