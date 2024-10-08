# UpdateInputParameterRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | Default value to use for an optional input parameter represented as a string | [optional] 
**Description** | Pointer to **string** | A brief description of the input parameter | [optional] 
**HasOptions** | Pointer to **bool** | Marks the input parameter to be selectable from a list of values | [optional] 
**LabeledOptions** | Pointer to **map[string]string** | A map for labeled options. The key is the label and the value is the option. When the option is selected, the label will be displayed to the end customer. Specify either options or labeledOptions when defining the input parameter. | [optional] 
**Limits** | Pointer to [**Limits**](Limits.md) |  | [optional] 
**Modifiable** | Pointer to **bool** | Marks the input parameter as immutable | [optional] 
**Name** | Pointer to **string** | External name for the input parameter | [optional] 
**Options** | Pointer to **[]string** | A list of options to restrict the value of the input parameter to (represented as a string) | [optional] 
**Required** | Pointer to **bool** | Enforces the input parameter as a required parameter | [optional] 

## Methods

### NewUpdateInputParameterRequestBody

`func NewUpdateInputParameterRequestBody() *UpdateInputParameterRequestBody`

NewUpdateInputParameterRequestBody instantiates a new UpdateInputParameterRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInputParameterRequestBodyWithDefaults

`func NewUpdateInputParameterRequestBodyWithDefaults() *UpdateInputParameterRequestBody`

NewUpdateInputParameterRequestBodyWithDefaults instantiates a new UpdateInputParameterRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *UpdateInputParameterRequestBody) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UpdateInputParameterRequestBody) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UpdateInputParameterRequestBody) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *UpdateInputParameterRequestBody) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInputParameterRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInputParameterRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInputParameterRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInputParameterRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasOptions

`func (o *UpdateInputParameterRequestBody) GetHasOptions() bool`

GetHasOptions returns the HasOptions field if non-nil, zero value otherwise.

### GetHasOptionsOk

`func (o *UpdateInputParameterRequestBody) GetHasOptionsOk() (*bool, bool)`

GetHasOptionsOk returns a tuple with the HasOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOptions

`func (o *UpdateInputParameterRequestBody) SetHasOptions(v bool)`

SetHasOptions sets HasOptions field to given value.

### HasHasOptions

`func (o *UpdateInputParameterRequestBody) HasHasOptions() bool`

HasHasOptions returns a boolean if a field has been set.

### GetLabeledOptions

`func (o *UpdateInputParameterRequestBody) GetLabeledOptions() map[string]string`

GetLabeledOptions returns the LabeledOptions field if non-nil, zero value otherwise.

### GetLabeledOptionsOk

`func (o *UpdateInputParameterRequestBody) GetLabeledOptionsOk() (*map[string]string, bool)`

GetLabeledOptionsOk returns a tuple with the LabeledOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabeledOptions

`func (o *UpdateInputParameterRequestBody) SetLabeledOptions(v map[string]string)`

SetLabeledOptions sets LabeledOptions field to given value.

### HasLabeledOptions

`func (o *UpdateInputParameterRequestBody) HasLabeledOptions() bool`

HasLabeledOptions returns a boolean if a field has been set.

### GetLimits

`func (o *UpdateInputParameterRequestBody) GetLimits() Limits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateInputParameterRequestBody) GetLimitsOk() (*Limits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *UpdateInputParameterRequestBody) SetLimits(v Limits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *UpdateInputParameterRequestBody) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetModifiable

`func (o *UpdateInputParameterRequestBody) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *UpdateInputParameterRequestBody) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *UpdateInputParameterRequestBody) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.

### HasModifiable

`func (o *UpdateInputParameterRequestBody) HasModifiable() bool`

HasModifiable returns a boolean if a field has been set.

### GetName

`func (o *UpdateInputParameterRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInputParameterRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInputParameterRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInputParameterRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateInputParameterRequestBody) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateInputParameterRequestBody) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateInputParameterRequestBody) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateInputParameterRequestBody) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRequired

`func (o *UpdateInputParameterRequestBody) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UpdateInputParameterRequestBody) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UpdateInputParameterRequestBody) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UpdateInputParameterRequestBody) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


