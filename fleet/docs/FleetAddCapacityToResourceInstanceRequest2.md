# FleetAddCapacityToResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityToBeAdded** | **int64** | Number of replicas to be added | 
**ResourceId** | **string** | The resource ID. | 

## Methods

### NewFleetAddCapacityToResourceInstanceRequest2

`func NewFleetAddCapacityToResourceInstanceRequest2(capacityToBeAdded int64, resourceId string, ) *FleetAddCapacityToResourceInstanceRequest2`

NewFleetAddCapacityToResourceInstanceRequest2 instantiates a new FleetAddCapacityToResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAddCapacityToResourceInstanceRequest2WithDefaults

`func NewFleetAddCapacityToResourceInstanceRequest2WithDefaults() *FleetAddCapacityToResourceInstanceRequest2`

NewFleetAddCapacityToResourceInstanceRequest2WithDefaults instantiates a new FleetAddCapacityToResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityToBeAdded

`func (o *FleetAddCapacityToResourceInstanceRequest2) GetCapacityToBeAdded() int64`

GetCapacityToBeAdded returns the CapacityToBeAdded field if non-nil, zero value otherwise.

### GetCapacityToBeAddedOk

`func (o *FleetAddCapacityToResourceInstanceRequest2) GetCapacityToBeAddedOk() (*int64, bool)`

GetCapacityToBeAddedOk returns a tuple with the CapacityToBeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityToBeAdded

`func (o *FleetAddCapacityToResourceInstanceRequest2) SetCapacityToBeAdded(v int64)`

SetCapacityToBeAdded sets CapacityToBeAdded field to given value.


### GetResourceId

`func (o *FleetAddCapacityToResourceInstanceRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetAddCapacityToResourceInstanceRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetAddCapacityToResourceInstanceRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


