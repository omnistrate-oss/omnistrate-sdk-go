# ResourceHealthSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodesHealth** | [**map[string]NodeHealthSummary**](NodeHealthSummary.md) | The health summary of the nodes by node name | 
**ResourceID** | **string** | ID of a resource | 
**ResourceKey** | **string** | The key of the resource | 
**ResourceType** | Pointer to **string** | The type of the resource | [optional] 
**Status** | **string** | The heath status of a resource | 

## Methods

### NewResourceHealthSummary

`func NewResourceHealthSummary(nodesHealth map[string]NodeHealthSummary, resourceID string, resourceKey string, status string, ) *ResourceHealthSummary`

NewResourceHealthSummary instantiates a new ResourceHealthSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceHealthSummaryWithDefaults

`func NewResourceHealthSummaryWithDefaults() *ResourceHealthSummary`

NewResourceHealthSummaryWithDefaults instantiates a new ResourceHealthSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodesHealth

`func (o *ResourceHealthSummary) GetNodesHealth() map[string]NodeHealthSummary`

GetNodesHealth returns the NodesHealth field if non-nil, zero value otherwise.

### GetNodesHealthOk

`func (o *ResourceHealthSummary) GetNodesHealthOk() (*map[string]NodeHealthSummary, bool)`

GetNodesHealthOk returns a tuple with the NodesHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesHealth

`func (o *ResourceHealthSummary) SetNodesHealth(v map[string]NodeHealthSummary)`

SetNodesHealth sets NodesHealth field to given value.


### GetResourceID

`func (o *ResourceHealthSummary) GetResourceID() string`

GetResourceID returns the ResourceID field if non-nil, zero value otherwise.

### GetResourceIDOk

`func (o *ResourceHealthSummary) GetResourceIDOk() (*string, bool)`

GetResourceIDOk returns a tuple with the ResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceID

`func (o *ResourceHealthSummary) SetResourceID(v string)`

SetResourceID sets ResourceID field to given value.


### GetResourceKey

`func (o *ResourceHealthSummary) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ResourceHealthSummary) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ResourceHealthSummary) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceType

`func (o *ResourceHealthSummary) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceHealthSummary) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceHealthSummary) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceHealthSummary) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceHealthSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceHealthSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceHealthSummary) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


