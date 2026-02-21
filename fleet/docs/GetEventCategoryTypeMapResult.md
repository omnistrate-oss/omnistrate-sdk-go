# GetEventCategoryTypeMapResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypeMap** | **map[string][]string** | Supported event category to event types mapping for notification subscriptions | 

## Methods

### NewGetEventCategoryTypeMapResult

`func NewGetEventCategoryTypeMapResult(categoryTypeMap map[string][]string, ) *GetEventCategoryTypeMapResult`

NewGetEventCategoryTypeMapResult instantiates a new GetEventCategoryTypeMapResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventCategoryTypeMapResultWithDefaults

`func NewGetEventCategoryTypeMapResultWithDefaults() *GetEventCategoryTypeMapResult`

NewGetEventCategoryTypeMapResultWithDefaults instantiates a new GetEventCategoryTypeMapResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTypeMap

`func (o *GetEventCategoryTypeMapResult) GetCategoryTypeMap() map[string][]string`

GetCategoryTypeMap returns the CategoryTypeMap field if non-nil, zero value otherwise.

### GetCategoryTypeMapOk

`func (o *GetEventCategoryTypeMapResult) GetCategoryTypeMapOk() (*map[string][]string, bool)`

GetCategoryTypeMapOk returns a tuple with the CategoryTypeMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypeMap

`func (o *GetEventCategoryTypeMapResult) SetCategoryTypeMap(v map[string][]string)`

SetCategoryTypeMap sets CategoryTypeMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


