# FleetUpdateResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tags for the resource instance. | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**RequestParams** | Pointer to **interface{}** | The request parameters | [optional] 
**ResourceId** | **string** | The resource ID. | 

## Methods

### NewFleetUpdateResourceInstanceRequest2

`func NewFleetUpdateResourceInstanceRequest2(resourceId string, ) *FleetUpdateResourceInstanceRequest2`

NewFleetUpdateResourceInstanceRequest2 instantiates a new FleetUpdateResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateResourceInstanceRequest2WithDefaults

`func NewFleetUpdateResourceInstanceRequest2WithDefaults() *FleetUpdateResourceInstanceRequest2`

NewFleetUpdateResourceInstanceRequest2WithDefaults instantiates a new FleetUpdateResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTags

`func (o *FleetUpdateResourceInstanceRequest2) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *FleetUpdateResourceInstanceRequest2) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *FleetUpdateResourceInstanceRequest2) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *FleetUpdateResourceInstanceRequest2) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetNetworkType

`func (o *FleetUpdateResourceInstanceRequest2) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *FleetUpdateResourceInstanceRequest2) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *FleetUpdateResourceInstanceRequest2) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *FleetUpdateResourceInstanceRequest2) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetRequestParams

`func (o *FleetUpdateResourceInstanceRequest2) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *FleetUpdateResourceInstanceRequest2) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *FleetUpdateResourceInstanceRequest2) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *FleetUpdateResourceInstanceRequest2) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *FleetUpdateResourceInstanceRequest2) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *FleetUpdateResourceInstanceRequest2) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil
### GetResourceId

`func (o *FleetUpdateResourceInstanceRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetUpdateResourceInstanceRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetUpdateResourceInstanceRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


