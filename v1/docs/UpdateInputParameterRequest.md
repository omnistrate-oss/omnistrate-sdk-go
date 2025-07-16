# UpdateInputParameterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | Default value to use for an optional input parameter represented as a string | [optional] 
**Description** | Pointer to **string** | A brief description of the input parameter | [optional] 
**HasOptions** | Pointer to **bool** | Marks the input parameter to be selectable from a list of values | [optional] 
**Id** | **string** | ID of an Input Parameter | 
**LabeledOptions** | Pointer to **map[string]string** | A map for labeled options. The key is the label and the value is the option. When the option is selected, the label will be displayed to the end customer. Specify either options or labeledOptions when defining the input parameter. | [optional] 
**Limits** | Pointer to [**Limits**](Limits.md) |  | [optional] 
**Modifiable** | Pointer to **bool** | Marks the input parameter as immutable | [optional] 
**Name** | Pointer to **string** | External name for the input parameter | [optional] 
**Options** | Pointer to **[]string** | A list of options to restrict the value of the input parameter to (represented as a string) | [optional] 
**Regex** | Pointer to **string** | Regular expression pattern for validating the input parameter value | [optional] 
**Required** | Pointer to **bool** | Enforces the input parameter as a required parameter | [optional] 
**ServiceId** | **string** | ID of a Service | 
**TabIndex** | Pointer to **int64** | Index for parameter ordering in the SaaS portal | [optional] [default to 0]
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateInputParameterRequest

`func NewUpdateInputParameterRequest(id string, serviceId string, token string, ) *UpdateInputParameterRequest`

NewUpdateInputParameterRequest instantiates a new UpdateInputParameterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInputParameterRequestWithDefaults

`func NewUpdateInputParameterRequestWithDefaults() *UpdateInputParameterRequest`

NewUpdateInputParameterRequestWithDefaults instantiates a new UpdateInputParameterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *UpdateInputParameterRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UpdateInputParameterRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UpdateInputParameterRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *UpdateInputParameterRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInputParameterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInputParameterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInputParameterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInputParameterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasOptions

`func (o *UpdateInputParameterRequest) GetHasOptions() bool`

GetHasOptions returns the HasOptions field if non-nil, zero value otherwise.

### GetHasOptionsOk

`func (o *UpdateInputParameterRequest) GetHasOptionsOk() (*bool, bool)`

GetHasOptionsOk returns a tuple with the HasOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOptions

`func (o *UpdateInputParameterRequest) SetHasOptions(v bool)`

SetHasOptions sets HasOptions field to given value.

### HasHasOptions

`func (o *UpdateInputParameterRequest) HasHasOptions() bool`

HasHasOptions returns a boolean if a field has been set.

### GetId

`func (o *UpdateInputParameterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInputParameterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInputParameterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLabeledOptions

`func (o *UpdateInputParameterRequest) GetLabeledOptions() map[string]string`

GetLabeledOptions returns the LabeledOptions field if non-nil, zero value otherwise.

### GetLabeledOptionsOk

`func (o *UpdateInputParameterRequest) GetLabeledOptionsOk() (*map[string]string, bool)`

GetLabeledOptionsOk returns a tuple with the LabeledOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabeledOptions

`func (o *UpdateInputParameterRequest) SetLabeledOptions(v map[string]string)`

SetLabeledOptions sets LabeledOptions field to given value.

### HasLabeledOptions

`func (o *UpdateInputParameterRequest) HasLabeledOptions() bool`

HasLabeledOptions returns a boolean if a field has been set.

### GetLimits

`func (o *UpdateInputParameterRequest) GetLimits() Limits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateInputParameterRequest) GetLimitsOk() (*Limits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *UpdateInputParameterRequest) SetLimits(v Limits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *UpdateInputParameterRequest) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetModifiable

`func (o *UpdateInputParameterRequest) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *UpdateInputParameterRequest) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *UpdateInputParameterRequest) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.

### HasModifiable

`func (o *UpdateInputParameterRequest) HasModifiable() bool`

HasModifiable returns a boolean if a field has been set.

### GetName

`func (o *UpdateInputParameterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInputParameterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInputParameterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInputParameterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateInputParameterRequest) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateInputParameterRequest) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateInputParameterRequest) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateInputParameterRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRegex

`func (o *UpdateInputParameterRequest) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *UpdateInputParameterRequest) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *UpdateInputParameterRequest) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *UpdateInputParameterRequest) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetRequired

`func (o *UpdateInputParameterRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UpdateInputParameterRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UpdateInputParameterRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UpdateInputParameterRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateInputParameterRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateInputParameterRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateInputParameterRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTabIndex

`func (o *UpdateInputParameterRequest) GetTabIndex() int64`

GetTabIndex returns the TabIndex field if non-nil, zero value otherwise.

### GetTabIndexOk

`func (o *UpdateInputParameterRequest) GetTabIndexOk() (*int64, bool)`

GetTabIndexOk returns a tuple with the TabIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabIndex

`func (o *UpdateInputParameterRequest) SetTabIndex(v int64)`

SetTabIndex sets TabIndex field to given value.

### HasTabIndex

`func (o *UpdateInputParameterRequest) HasTabIndex() bool`

HasTabIndex returns a boolean if a field has been set.

### GetToken

`func (o *UpdateInputParameterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateInputParameterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateInputParameterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


