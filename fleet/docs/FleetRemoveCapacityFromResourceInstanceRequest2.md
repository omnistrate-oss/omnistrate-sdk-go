# FleetRemoveCapacityFromResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityToBeRemoved** | **int64** | Number of replicas to be removed | 
**ResourceId** | **string** | The resource ID. | 

## Methods

### NewFleetRemoveCapacityFromResourceInstanceRequest2

`func NewFleetRemoveCapacityFromResourceInstanceRequest2(capacityToBeRemoved int64, resourceId string, ) *FleetRemoveCapacityFromResourceInstanceRequest2`

NewFleetRemoveCapacityFromResourceInstanceRequest2 instantiates a new FleetRemoveCapacityFromResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetRemoveCapacityFromResourceInstanceRequest2WithDefaults

`func NewFleetRemoveCapacityFromResourceInstanceRequest2WithDefaults() *FleetRemoveCapacityFromResourceInstanceRequest2`

NewFleetRemoveCapacityFromResourceInstanceRequest2WithDefaults instantiates a new FleetRemoveCapacityFromResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityToBeRemoved

`func (o *FleetRemoveCapacityFromResourceInstanceRequest2) GetCapacityToBeRemoved() int64`

GetCapacityToBeRemoved returns the CapacityToBeRemoved field if non-nil, zero value otherwise.

### GetCapacityToBeRemovedOk

`func (o *FleetRemoveCapacityFromResourceInstanceRequest2) GetCapacityToBeRemovedOk() (*int64, bool)`

GetCapacityToBeRemovedOk returns a tuple with the CapacityToBeRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityToBeRemoved

`func (o *FleetRemoveCapacityFromResourceInstanceRequest2) SetCapacityToBeRemoved(v int64)`

SetCapacityToBeRemoved sets CapacityToBeRemoved field to given value.


### GetResourceId

`func (o *FleetRemoveCapacityFromResourceInstanceRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetRemoveCapacityFromResourceInstanceRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetRemoveCapacityFromResourceInstanceRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


