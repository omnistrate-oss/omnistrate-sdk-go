# ListEntitiesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]Entity**](Entity.md) | Array of entities matching the criteria | [optional] 

## Methods

### NewListEntitiesResult

`func NewListEntitiesResult() *ListEntitiesResult`

NewListEntitiesResult instantiates a new ListEntitiesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitiesResultWithDefaults

`func NewListEntitiesResultWithDefaults() *ListEntitiesResult`

NewListEntitiesResultWithDefaults instantiates a new ListEntitiesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ListEntitiesResult) GetEntities() []Entity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ListEntitiesResult) GetEntitiesOk() (*[]Entity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ListEntitiesResult) SetEntities(v []Entity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ListEntitiesResult) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


