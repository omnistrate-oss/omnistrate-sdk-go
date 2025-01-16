# DescribePendingChangesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Service API | 
**ResourceChangeSets** | **map[string]interface{}** | The difference for the service API configuration per resource | 
**ServiceId** | **string** | ID of a Service | 

## Methods

### NewDescribePendingChangesResult

`func NewDescribePendingChangesResult(id string, resourceChangeSets map[string]interface{}, serviceId string, ) *DescribePendingChangesResult`

NewDescribePendingChangesResult instantiates a new DescribePendingChangesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePendingChangesResultWithDefaults

`func NewDescribePendingChangesResultWithDefaults() *DescribePendingChangesResult`

NewDescribePendingChangesResultWithDefaults instantiates a new DescribePendingChangesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescribePendingChangesResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribePendingChangesResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribePendingChangesResult) SetId(v string)`

SetId sets Id field to given value.


### GetResourceChangeSets

`func (o *DescribePendingChangesResult) GetResourceChangeSets() map[string]interface{}`

GetResourceChangeSets returns the ResourceChangeSets field if non-nil, zero value otherwise.

### GetResourceChangeSetsOk

`func (o *DescribePendingChangesResult) GetResourceChangeSetsOk() (*map[string]interface{}, bool)`

GetResourceChangeSetsOk returns a tuple with the ResourceChangeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceChangeSets

`func (o *DescribePendingChangesResult) SetResourceChangeSets(v map[string]interface{})`

SetResourceChangeSets sets ResourceChangeSets field to given value.


### GetServiceId

`func (o *DescribePendingChangesResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribePendingChangesResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribePendingChangesResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


