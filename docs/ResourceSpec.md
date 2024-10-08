# ResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to [**ResourceSpecLimits**](ResourceSpecLimits.md) |  | [optional] 
**Requests** | Pointer to [**ResourceSpecRequests**](ResourceSpecRequests.md) |  | [optional] 

## Methods

### NewResourceSpec

`func NewResourceSpec() *ResourceSpec`

NewResourceSpec instantiates a new ResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSpecWithDefaults

`func NewResourceSpecWithDefaults() *ResourceSpec`

NewResourceSpecWithDefaults instantiates a new ResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *ResourceSpec) GetLimits() ResourceSpecLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ResourceSpec) GetLimitsOk() (*ResourceSpecLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ResourceSpec) SetLimits(v ResourceSpecLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ResourceSpec) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetRequests

`func (o *ResourceSpec) GetRequests() ResourceSpecRequests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ResourceSpec) GetRequestsOk() (*ResourceSpecRequests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ResourceSpec) SetRequests(v ResourceSpecRequests)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ResourceSpec) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


