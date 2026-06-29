# ResourceInstanceSearchFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Predicates** | Pointer to [**[]ResourceInstanceFilterGroup**](ResourceInstanceFilterGroup.md) | Resource instance predicate groups. Predicates within a group are combined with AND; groups are combined with OR. | [optional] 
**Tags** | Pointer to [**[]ResourceInstanceTagFilter**](ResourceInstanceTagFilter.md) | Custom tag predicates. All tag predicates must match. | [optional] 

## Methods

### NewResourceInstanceSearchFilters

`func NewResourceInstanceSearchFilters() *ResourceInstanceSearchFilters`

NewResourceInstanceSearchFilters instantiates a new ResourceInstanceSearchFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInstanceSearchFiltersWithDefaults

`func NewResourceInstanceSearchFiltersWithDefaults() *ResourceInstanceSearchFilters`

NewResourceInstanceSearchFiltersWithDefaults instantiates a new ResourceInstanceSearchFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredicates

`func (o *ResourceInstanceSearchFilters) GetPredicates() []ResourceInstanceFilterGroup`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *ResourceInstanceSearchFilters) GetPredicatesOk() (*[]ResourceInstanceFilterGroup, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *ResourceInstanceSearchFilters) SetPredicates(v []ResourceInstanceFilterGroup)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *ResourceInstanceSearchFilters) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetTags

`func (o *ResourceInstanceSearchFilters) GetTags() []ResourceInstanceTagFilter`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResourceInstanceSearchFilters) GetTagsOk() (*[]ResourceInstanceTagFilter, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResourceInstanceSearchFilters) SetTags(v []ResourceInstanceTagFilter)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResourceInstanceSearchFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


