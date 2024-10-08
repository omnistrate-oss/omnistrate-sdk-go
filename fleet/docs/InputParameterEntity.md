# InputParameterEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | **bool** | Whether the parameter is custom | 
**DefaultValue** | Pointer to **string** | The parameter default value | [optional] 
**DependentResourceID** | Pointer to **string** | The parameter dependent resource ID | [optional] 
**Description** | **string** | The parameter description | 
**DisplayName** | **string** | The parameter display name | 
**IsList** | **bool** | Whether the parameter is a list | 
**Key** | **string** | The parameter key | 
**Modifiable** | **bool** | Whether the parameter is modifiable | 
**Options** | Pointer to **[]string** | The parameter options | [optional] 
**Required** | **bool** | Whether the parameter is required | 
**Type** | **string** | The parameter type | 

## Methods

### NewInputParameterEntity

`func NewInputParameterEntity(custom bool, description string, displayName string, isList bool, key string, modifiable bool, required bool, type_ string, ) *InputParameterEntity`

NewInputParameterEntity instantiates a new InputParameterEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputParameterEntityWithDefaults

`func NewInputParameterEntityWithDefaults() *InputParameterEntity`

NewInputParameterEntityWithDefaults instantiates a new InputParameterEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *InputParameterEntity) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *InputParameterEntity) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *InputParameterEntity) SetCustom(v bool)`

SetCustom sets Custom field to given value.


### GetDefaultValue

`func (o *InputParameterEntity) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *InputParameterEntity) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *InputParameterEntity) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *InputParameterEntity) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDependentResourceID

`func (o *InputParameterEntity) GetDependentResourceID() string`

GetDependentResourceID returns the DependentResourceID field if non-nil, zero value otherwise.

### GetDependentResourceIDOk

`func (o *InputParameterEntity) GetDependentResourceIDOk() (*string, bool)`

GetDependentResourceIDOk returns a tuple with the DependentResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentResourceID

`func (o *InputParameterEntity) SetDependentResourceID(v string)`

SetDependentResourceID sets DependentResourceID field to given value.

### HasDependentResourceID

`func (o *InputParameterEntity) HasDependentResourceID() bool`

HasDependentResourceID returns a boolean if a field has been set.

### GetDescription

`func (o *InputParameterEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputParameterEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputParameterEntity) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *InputParameterEntity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InputParameterEntity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InputParameterEntity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsList

`func (o *InputParameterEntity) GetIsList() bool`

GetIsList returns the IsList field if non-nil, zero value otherwise.

### GetIsListOk

`func (o *InputParameterEntity) GetIsListOk() (*bool, bool)`

GetIsListOk returns a tuple with the IsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsList

`func (o *InputParameterEntity) SetIsList(v bool)`

SetIsList sets IsList field to given value.


### GetKey

`func (o *InputParameterEntity) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InputParameterEntity) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InputParameterEntity) SetKey(v string)`

SetKey sets Key field to given value.


### GetModifiable

`func (o *InputParameterEntity) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *InputParameterEntity) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *InputParameterEntity) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.


### GetOptions

`func (o *InputParameterEntity) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InputParameterEntity) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InputParameterEntity) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InputParameterEntity) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRequired

`func (o *InputParameterEntity) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *InputParameterEntity) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *InputParameterEntity) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetType

`func (o *InputParameterEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputParameterEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputParameterEntity) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


