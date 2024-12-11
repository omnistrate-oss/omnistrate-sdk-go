# AddCapacityToResourceInstanceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityToBeAdded** | **int64** | Number of replicas to be added | 
**ResourceId** | **string** | The resource ID. | 

## Methods

### NewAddCapacityToResourceInstanceRequestBody

`func NewAddCapacityToResourceInstanceRequestBody(capacityToBeAdded int64, resourceId string, ) *AddCapacityToResourceInstanceRequestBody`

NewAddCapacityToResourceInstanceRequestBody instantiates a new AddCapacityToResourceInstanceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCapacityToResourceInstanceRequestBodyWithDefaults

`func NewAddCapacityToResourceInstanceRequestBodyWithDefaults() *AddCapacityToResourceInstanceRequestBody`

NewAddCapacityToResourceInstanceRequestBodyWithDefaults instantiates a new AddCapacityToResourceInstanceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityToBeAdded

`func (o *AddCapacityToResourceInstanceRequestBody) GetCapacityToBeAdded() int64`

GetCapacityToBeAdded returns the CapacityToBeAdded field if non-nil, zero value otherwise.

### GetCapacityToBeAddedOk

`func (o *AddCapacityToResourceInstanceRequestBody) GetCapacityToBeAddedOk() (*int64, bool)`

GetCapacityToBeAddedOk returns a tuple with the CapacityToBeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityToBeAdded

`func (o *AddCapacityToResourceInstanceRequestBody) SetCapacityToBeAdded(v int64)`

SetCapacityToBeAdded sets CapacityToBeAdded field to given value.


### GetResourceId

`func (o *AddCapacityToResourceInstanceRequestBody) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AddCapacityToResourceInstanceRequestBody) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AddCapacityToResourceInstanceRequestBody) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


