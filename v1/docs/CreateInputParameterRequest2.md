# CreateInputParameterRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | Default value to use for an optional input parameter represented as a string | [optional] 
**DependentResourceId** | Pointer to **string** | The ID of the resource whose instance this input parameter depends on | [optional] 
**Description** | **string** | A brief description of the input parameter | 
**HasOptions** | Pointer to **bool** | Marks the input parameter to be selectable from a list of values | [optional] [default to false]
**IsList** | Pointer to **bool** | Marks the input parameter as a list of values | [optional] [default to false]
**Key** | **string** | Key of the input parameter | 
**LabeledOptions** | Pointer to **map[string]string** | A map for labeled options. The key is the label and the value is the option. When the option is selected, the label will be displayed to the end customer. Specify either options or labeledOptions when defining the input parameter. | [optional] 
**Limits** | Pointer to [**Limits**](Limits.md) |  | [optional] 
**Modifiable** | **bool** | Marks the input parameter as immutable | 
**Name** | **string** | External name for the input parameter | 
**Options** | Pointer to **[]string** | A list of options to restrict the value of the input parameter to (represented as a string) | [optional] 
**Required** | **bool** | Enforces the input parameter as a required parameter | 
**ResourceId** | **string** | The ID of the resource that this input parameter belongs to | 
**TabIndex** | Pointer to **int64** | Index for parameter ordering in the SaaS portal | [optional] [default to 0]
**Type** | **string** |  | 

## Methods

### NewCreateInputParameterRequest2

`func NewCreateInputParameterRequest2(description string, key string, modifiable bool, name string, required bool, resourceId string, type_ string, ) *CreateInputParameterRequest2`

NewCreateInputParameterRequest2 instantiates a new CreateInputParameterRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInputParameterRequest2WithDefaults

`func NewCreateInputParameterRequest2WithDefaults() *CreateInputParameterRequest2`

NewCreateInputParameterRequest2WithDefaults instantiates a new CreateInputParameterRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *CreateInputParameterRequest2) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CreateInputParameterRequest2) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CreateInputParameterRequest2) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CreateInputParameterRequest2) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDependentResourceId

`func (o *CreateInputParameterRequest2) GetDependentResourceId() string`

GetDependentResourceId returns the DependentResourceId field if non-nil, zero value otherwise.

### GetDependentResourceIdOk

`func (o *CreateInputParameterRequest2) GetDependentResourceIdOk() (*string, bool)`

GetDependentResourceIdOk returns a tuple with the DependentResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentResourceId

`func (o *CreateInputParameterRequest2) SetDependentResourceId(v string)`

SetDependentResourceId sets DependentResourceId field to given value.

### HasDependentResourceId

`func (o *CreateInputParameterRequest2) HasDependentResourceId() bool`

HasDependentResourceId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateInputParameterRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInputParameterRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInputParameterRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHasOptions

`func (o *CreateInputParameterRequest2) GetHasOptions() bool`

GetHasOptions returns the HasOptions field if non-nil, zero value otherwise.

### GetHasOptionsOk

`func (o *CreateInputParameterRequest2) GetHasOptionsOk() (*bool, bool)`

GetHasOptionsOk returns a tuple with the HasOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOptions

`func (o *CreateInputParameterRequest2) SetHasOptions(v bool)`

SetHasOptions sets HasOptions field to given value.

### HasHasOptions

`func (o *CreateInputParameterRequest2) HasHasOptions() bool`

HasHasOptions returns a boolean if a field has been set.

### GetIsList

`func (o *CreateInputParameterRequest2) GetIsList() bool`

GetIsList returns the IsList field if non-nil, zero value otherwise.

### GetIsListOk

`func (o *CreateInputParameterRequest2) GetIsListOk() (*bool, bool)`

GetIsListOk returns a tuple with the IsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsList

`func (o *CreateInputParameterRequest2) SetIsList(v bool)`

SetIsList sets IsList field to given value.

### HasIsList

`func (o *CreateInputParameterRequest2) HasIsList() bool`

HasIsList returns a boolean if a field has been set.

### GetKey

`func (o *CreateInputParameterRequest2) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateInputParameterRequest2) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateInputParameterRequest2) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabeledOptions

`func (o *CreateInputParameterRequest2) GetLabeledOptions() map[string]string`

GetLabeledOptions returns the LabeledOptions field if non-nil, zero value otherwise.

### GetLabeledOptionsOk

`func (o *CreateInputParameterRequest2) GetLabeledOptionsOk() (*map[string]string, bool)`

GetLabeledOptionsOk returns a tuple with the LabeledOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabeledOptions

`func (o *CreateInputParameterRequest2) SetLabeledOptions(v map[string]string)`

SetLabeledOptions sets LabeledOptions field to given value.

### HasLabeledOptions

`func (o *CreateInputParameterRequest2) HasLabeledOptions() bool`

HasLabeledOptions returns a boolean if a field has been set.

### GetLimits

`func (o *CreateInputParameterRequest2) GetLimits() Limits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CreateInputParameterRequest2) GetLimitsOk() (*Limits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *CreateInputParameterRequest2) SetLimits(v Limits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *CreateInputParameterRequest2) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetModifiable

`func (o *CreateInputParameterRequest2) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *CreateInputParameterRequest2) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *CreateInputParameterRequest2) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.


### GetName

`func (o *CreateInputParameterRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInputParameterRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInputParameterRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *CreateInputParameterRequest2) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateInputParameterRequest2) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateInputParameterRequest2) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateInputParameterRequest2) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRequired

`func (o *CreateInputParameterRequest2) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CreateInputParameterRequest2) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CreateInputParameterRequest2) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetResourceId

`func (o *CreateInputParameterRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateInputParameterRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateInputParameterRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetTabIndex

`func (o *CreateInputParameterRequest2) GetTabIndex() int64`

GetTabIndex returns the TabIndex field if non-nil, zero value otherwise.

### GetTabIndexOk

`func (o *CreateInputParameterRequest2) GetTabIndexOk() (*int64, bool)`

GetTabIndexOk returns a tuple with the TabIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabIndex

`func (o *CreateInputParameterRequest2) SetTabIndex(v int64)`

SetTabIndex sets TabIndex field to given value.

### HasTabIndex

`func (o *CreateInputParameterRequest2) HasTabIndex() bool`

HasTabIndex returns a boolean if a field has been set.

### GetType

`func (o *CreateInputParameterRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateInputParameterRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateInputParameterRequest2) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


