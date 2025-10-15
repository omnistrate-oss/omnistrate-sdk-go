# UpdateResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tag for the resource instance | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**RequestParams** | Pointer to **interface{}** | The request parameters | [optional] 

## Methods

### NewUpdateResourceInstanceRequest2

`func NewUpdateResourceInstanceRequest2() *UpdateResourceInstanceRequest2`

NewUpdateResourceInstanceRequest2 instantiates a new UpdateResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceInstanceRequest2WithDefaults

`func NewUpdateResourceInstanceRequest2WithDefaults() *UpdateResourceInstanceRequest2`

NewUpdateResourceInstanceRequest2WithDefaults instantiates a new UpdateResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTags

`func (o *UpdateResourceInstanceRequest2) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *UpdateResourceInstanceRequest2) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *UpdateResourceInstanceRequest2) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *UpdateResourceInstanceRequest2) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetNetworkType

`func (o *UpdateResourceInstanceRequest2) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *UpdateResourceInstanceRequest2) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *UpdateResourceInstanceRequest2) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *UpdateResourceInstanceRequest2) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetRequestParams

`func (o *UpdateResourceInstanceRequest2) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *UpdateResourceInstanceRequest2) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *UpdateResourceInstanceRequest2) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *UpdateResourceInstanceRequest2) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *UpdateResourceInstanceRequest2) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *UpdateResourceInstanceRequest2) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


