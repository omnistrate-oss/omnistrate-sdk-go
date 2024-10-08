# ResourceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoscalingEnabled** | Pointer to **bool** | Whether the resource has autoscaling enabled | [optional] 
**IsBackupEnabled** | **bool** | Whether the resource has backup enabled | 
**IsDeprecated** | **bool** | Whether the service offering is deprecated | [default to false]
**Name** | **string** | The resource name | 
**ResourceId** | **string** | The resource ID | 
**ResourceType** | Pointer to **string** | The resource type | [optional] 
**UrlKey** | **string** | The resource URL key | 

## Methods

### NewResourceEntity

`func NewResourceEntity(isBackupEnabled bool, isDeprecated bool, name string, resourceId string, urlKey string, ) *ResourceEntity`

NewResourceEntity instantiates a new ResourceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEntityWithDefaults

`func NewResourceEntityWithDefaults() *ResourceEntity`

NewResourceEntityWithDefaults instantiates a new ResourceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoscalingEnabled

`func (o *ResourceEntity) GetIsAutoscalingEnabled() bool`

GetIsAutoscalingEnabled returns the IsAutoscalingEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingEnabledOk

`func (o *ResourceEntity) GetIsAutoscalingEnabledOk() (*bool, bool)`

GetIsAutoscalingEnabledOk returns a tuple with the IsAutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingEnabled

`func (o *ResourceEntity) SetIsAutoscalingEnabled(v bool)`

SetIsAutoscalingEnabled sets IsAutoscalingEnabled field to given value.

### HasIsAutoscalingEnabled

`func (o *ResourceEntity) HasIsAutoscalingEnabled() bool`

HasIsAutoscalingEnabled returns a boolean if a field has been set.

### GetIsBackupEnabled

`func (o *ResourceEntity) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ResourceEntity) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ResourceEntity) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.


### GetIsDeprecated

`func (o *ResourceEntity) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *ResourceEntity) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *ResourceEntity) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetName

`func (o *ResourceEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceEntity) SetName(v string)`

SetName sets Name field to given value.


### GetResourceId

`func (o *ResourceEntity) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceEntity) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceEntity) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ResourceEntity) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceEntity) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceEntity) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceEntity) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetUrlKey

`func (o *ResourceEntity) GetUrlKey() string`

GetUrlKey returns the UrlKey field if non-nil, zero value otherwise.

### GetUrlKeyOk

`func (o *ResourceEntity) GetUrlKeyOk() (*string, bool)`

GetUrlKeyOk returns a tuple with the UrlKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlKey

`func (o *ResourceEntity) SetUrlKey(v string)`

SetUrlKey sets UrlKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


