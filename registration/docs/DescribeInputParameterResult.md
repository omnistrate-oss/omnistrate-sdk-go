# DescribeInputParameterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | Default value to use for an optional input parameter represented as a string | [optional] 
**DependentResourceId** | Pointer to **string** | The ID of the resource whose instance this input parameter depends on | [optional] 
**Description** | **string** | A brief description of the input parameter | 
**HasOptions** | Pointer to **bool** | Marks the input parameter to be selectable from a list of values | [optional] [default to false]
**Id** | **string** | ID of the input parameter | 
**IsList** | **bool** | Marks the input parameter as a list of values | [default to false]
**Key** | **string** | Key of the input parameter | 
**LabeledOptions** | Pointer to **map[string]string** | A map for labeled options. The key is the label and the value is the option. When the option is selected, the label will be displayed to the end customer. | [optional] 
**Limits** | Pointer to [**Limits**](Limits.md) |  | [optional] 
**Modifiable** | **bool** | Marks the input parameter as immutable | 
**Name** | **string** | External name for the input parameter | 
**Options** | Pointer to **[]string** | A list of options to restrict the value of the input parameter to (represented as a string) | [optional] 
**Required** | **bool** | Enforces the input parameter as a required parameter | 
**ResourceId** | **string** | The ID of the resource that this input parameter belongs to | 
**ServiceId** | **string** | The ID of the service that this output parameter belongs to | 
**Type** | **string** |  | 

## Methods

### NewDescribeInputParameterResult

`func NewDescribeInputParameterResult(description string, id string, isList bool, key string, modifiable bool, name string, required bool, resourceId string, serviceId string, type_ string, ) *DescribeInputParameterResult`

NewDescribeInputParameterResult instantiates a new DescribeInputParameterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeInputParameterResultWithDefaults

`func NewDescribeInputParameterResultWithDefaults() *DescribeInputParameterResult`

NewDescribeInputParameterResultWithDefaults instantiates a new DescribeInputParameterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *DescribeInputParameterResult) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DescribeInputParameterResult) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DescribeInputParameterResult) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DescribeInputParameterResult) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDependentResourceId

`func (o *DescribeInputParameterResult) GetDependentResourceId() string`

GetDependentResourceId returns the DependentResourceId field if non-nil, zero value otherwise.

### GetDependentResourceIdOk

`func (o *DescribeInputParameterResult) GetDependentResourceIdOk() (*string, bool)`

GetDependentResourceIdOk returns a tuple with the DependentResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentResourceId

`func (o *DescribeInputParameterResult) SetDependentResourceId(v string)`

SetDependentResourceId sets DependentResourceId field to given value.

### HasDependentResourceId

`func (o *DescribeInputParameterResult) HasDependentResourceId() bool`

HasDependentResourceId returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeInputParameterResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeInputParameterResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeInputParameterResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHasOptions

`func (o *DescribeInputParameterResult) GetHasOptions() bool`

GetHasOptions returns the HasOptions field if non-nil, zero value otherwise.

### GetHasOptionsOk

`func (o *DescribeInputParameterResult) GetHasOptionsOk() (*bool, bool)`

GetHasOptionsOk returns a tuple with the HasOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOptions

`func (o *DescribeInputParameterResult) SetHasOptions(v bool)`

SetHasOptions sets HasOptions field to given value.

### HasHasOptions

`func (o *DescribeInputParameterResult) HasHasOptions() bool`

HasHasOptions returns a boolean if a field has been set.

### GetId

`func (o *DescribeInputParameterResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeInputParameterResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeInputParameterResult) SetId(v string)`

SetId sets Id field to given value.


### GetIsList

`func (o *DescribeInputParameterResult) GetIsList() bool`

GetIsList returns the IsList field if non-nil, zero value otherwise.

### GetIsListOk

`func (o *DescribeInputParameterResult) GetIsListOk() (*bool, bool)`

GetIsListOk returns a tuple with the IsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsList

`func (o *DescribeInputParameterResult) SetIsList(v bool)`

SetIsList sets IsList field to given value.


### GetKey

`func (o *DescribeInputParameterResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DescribeInputParameterResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DescribeInputParameterResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabeledOptions

`func (o *DescribeInputParameterResult) GetLabeledOptions() map[string]string`

GetLabeledOptions returns the LabeledOptions field if non-nil, zero value otherwise.

### GetLabeledOptionsOk

`func (o *DescribeInputParameterResult) GetLabeledOptionsOk() (*map[string]string, bool)`

GetLabeledOptionsOk returns a tuple with the LabeledOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabeledOptions

`func (o *DescribeInputParameterResult) SetLabeledOptions(v map[string]string)`

SetLabeledOptions sets LabeledOptions field to given value.

### HasLabeledOptions

`func (o *DescribeInputParameterResult) HasLabeledOptions() bool`

HasLabeledOptions returns a boolean if a field has been set.

### GetLimits

`func (o *DescribeInputParameterResult) GetLimits() Limits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *DescribeInputParameterResult) GetLimitsOk() (*Limits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *DescribeInputParameterResult) SetLimits(v Limits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *DescribeInputParameterResult) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetModifiable

`func (o *DescribeInputParameterResult) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *DescribeInputParameterResult) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *DescribeInputParameterResult) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.


### GetName

`func (o *DescribeInputParameterResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeInputParameterResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeInputParameterResult) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *DescribeInputParameterResult) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DescribeInputParameterResult) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DescribeInputParameterResult) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DescribeInputParameterResult) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRequired

`func (o *DescribeInputParameterResult) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DescribeInputParameterResult) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DescribeInputParameterResult) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetResourceId

`func (o *DescribeInputParameterResult) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DescribeInputParameterResult) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DescribeInputParameterResult) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *DescribeInputParameterResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeInputParameterResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeInputParameterResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetType

`func (o *DescribeInputParameterResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescribeInputParameterResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescribeInputParameterResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


