# OneOffPatchResourceInstanceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The resource ID. | 
**TargetTierVersion** | Pointer to **string** | The target resource version. | [optional] 

## Methods

### NewOneOffPatchResourceInstanceRequestBody

`func NewOneOffPatchResourceInstanceRequestBody(resourceId string, ) *OneOffPatchResourceInstanceRequestBody`

NewOneOffPatchResourceInstanceRequestBody instantiates a new OneOffPatchResourceInstanceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneOffPatchResourceInstanceRequestBodyWithDefaults

`func NewOneOffPatchResourceInstanceRequestBodyWithDefaults() *OneOffPatchResourceInstanceRequestBody`

NewOneOffPatchResourceInstanceRequestBodyWithDefaults instantiates a new OneOffPatchResourceInstanceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *OneOffPatchResourceInstanceRequestBody) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OneOffPatchResourceInstanceRequestBody) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OneOffPatchResourceInstanceRequestBody) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequestBody) GetTargetTierVersion() string`

GetTargetTierVersion returns the TargetTierVersion field if non-nil, zero value otherwise.

### GetTargetTierVersionOk

`func (o *OneOffPatchResourceInstanceRequestBody) GetTargetTierVersionOk() (*string, bool)`

GetTargetTierVersionOk returns a tuple with the TargetTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequestBody) SetTargetTierVersion(v string)`

SetTargetTierVersion sets TargetTierVersion field to given value.

### HasTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequestBody) HasTargetTierVersion() bool`

HasTargetTierVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


