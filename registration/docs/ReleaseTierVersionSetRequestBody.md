# ReleaseTierVersionSetRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPreferred** | Pointer to **bool** | Indicates whether this version set is preferred. | [optional] 
**Name** | Pointer to **string** | The name of the product-tier version set. | [optional] 

## Methods

### NewReleaseTierVersionSetRequestBody

`func NewReleaseTierVersionSetRequestBody() *ReleaseTierVersionSetRequestBody`

NewReleaseTierVersionSetRequestBody instantiates a new ReleaseTierVersionSetRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseTierVersionSetRequestBodyWithDefaults

`func NewReleaseTierVersionSetRequestBodyWithDefaults() *ReleaseTierVersionSetRequestBody`

NewReleaseTierVersionSetRequestBodyWithDefaults instantiates a new ReleaseTierVersionSetRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPreferred

`func (o *ReleaseTierVersionSetRequestBody) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *ReleaseTierVersionSetRequestBody) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *ReleaseTierVersionSetRequestBody) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *ReleaseTierVersionSetRequestBody) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetName

`func (o *ReleaseTierVersionSetRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleaseTierVersionSetRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleaseTierVersionSetRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReleaseTierVersionSetRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


