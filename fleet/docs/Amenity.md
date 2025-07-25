# Amenity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the amenity. | [optional] 
**IsManaged** | Pointer to **bool** | Whether the amenity is managed by the system. | [optional] 
**Modifiable** | Pointer to **bool** | Whether the amenity can be modified. | [optional] 
**Name** | Pointer to **string** | The name of the amenity. | [optional] 
**Properties** | Pointer to **map[string]interface{}** | The properties of the amenity. | [optional] 
**Type** | Pointer to **string** | The type of the amenity. | [optional] 

## Methods

### NewAmenity

`func NewAmenity() *Amenity`

NewAmenity instantiates a new Amenity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmenityWithDefaults

`func NewAmenityWithDefaults() *Amenity`

NewAmenityWithDefaults instantiates a new Amenity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Amenity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Amenity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Amenity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Amenity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsManaged

`func (o *Amenity) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *Amenity) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *Amenity) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *Amenity) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetModifiable

`func (o *Amenity) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *Amenity) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *Amenity) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.

### HasModifiable

`func (o *Amenity) HasModifiable() bool`

HasModifiable returns a boolean if a field has been set.

### GetName

`func (o *Amenity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Amenity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Amenity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Amenity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *Amenity) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Amenity) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Amenity) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Amenity) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *Amenity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Amenity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Amenity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Amenity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


