# InfrastructurePDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisruptionsAllowed** | Pointer to **int64** |  | [optional] 
**MaxUnavailable** | Pointer to **string** |  | [optional] 
**MinAvailable** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**Selector** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructurePDB

`func NewInfrastructurePDB() *InfrastructurePDB`

NewInfrastructurePDB instantiates a new InfrastructurePDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructurePDBWithDefaults

`func NewInfrastructurePDBWithDefaults() *InfrastructurePDB`

NewInfrastructurePDBWithDefaults instantiates a new InfrastructurePDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisruptionsAllowed

`func (o *InfrastructurePDB) GetDisruptionsAllowed() int64`

GetDisruptionsAllowed returns the DisruptionsAllowed field if non-nil, zero value otherwise.

### GetDisruptionsAllowedOk

`func (o *InfrastructurePDB) GetDisruptionsAllowedOk() (*int64, bool)`

GetDisruptionsAllowedOk returns a tuple with the DisruptionsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptionsAllowed

`func (o *InfrastructurePDB) SetDisruptionsAllowed(v int64)`

SetDisruptionsAllowed sets DisruptionsAllowed field to given value.

### HasDisruptionsAllowed

`func (o *InfrastructurePDB) HasDisruptionsAllowed() bool`

HasDisruptionsAllowed returns a boolean if a field has been set.

### GetMaxUnavailable

`func (o *InfrastructurePDB) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *InfrastructurePDB) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *InfrastructurePDB) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *InfrastructurePDB) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetMinAvailable

`func (o *InfrastructurePDB) GetMinAvailable() string`

GetMinAvailable returns the MinAvailable field if non-nil, zero value otherwise.

### GetMinAvailableOk

`func (o *InfrastructurePDB) GetMinAvailableOk() (*string, bool)`

GetMinAvailableOk returns a tuple with the MinAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailable

`func (o *InfrastructurePDB) SetMinAvailable(v string)`

SetMinAvailable sets MinAvailable field to given value.

### HasMinAvailable

`func (o *InfrastructurePDB) HasMinAvailable() bool`

HasMinAvailable returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructurePDB) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructurePDB) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructurePDB) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructurePDB) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSelector

`func (o *InfrastructurePDB) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *InfrastructurePDB) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *InfrastructurePDB) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *InfrastructurePDB) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


