# CreateInputParameterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | Default value to use for an optional input parameter represented as a string | [optional] 
**DependentResourceId** | Pointer to **string** | ID of a resource | [optional] 
**Description** | **string** | A brief description of the input parameter | 
**HasOptions** | Pointer to **bool** | Marks the input parameter to be selectable from a list of values | [optional] [default to false]
**IsList** | Pointer to **bool** | Marks the input parameter as a list of values | [optional] [default to false]
**Key** | **string** | Key of the input parameter | 
**LabeledOptions** | Pointer to **map[string]string** | A map for labeled options. The key is the label and the value is the option. When the option is selected, the label will be displayed to the end customer. Specify either options or labeledOptions when defining the input parameter. | [optional] 
**Limits** | Pointer to [**Limits**](Limits.md) |  | [optional] 
**Modifiable** | **bool** | Marks the input parameter as immutable | 
**Name** | **string** | External name for the input parameter | 
**Options** | Pointer to **[]string** | A list of options to restrict the value of the input parameter to (represented as a string) | [optional] 
**Regex** | Pointer to **string** | Regular expression pattern for validating the input parameter value | [optional] 
**Required** | **bool** | Enforces the input parameter as a required parameter | 
**ResourceId** | **string** | ID of a resource | 
**Scope** | Pointer to [**InputParameterScope**](InputParameterScope.md) |  | [optional] 
**ServiceId** | **string** | ID of a Service | 
**TabIndex** | Pointer to **int64** | Index for parameter ordering in the SaaS portal | [optional] [default to 0]
**Token** | **string** | JWT token used to perform authorization | 
**Type** | **string** | Type of the variable encoding the value | 

## Methods

### NewCreateInputParameterRequest

`func NewCreateInputParameterRequest(description string, key string, modifiable bool, name string, required bool, resourceId string, serviceId string, token string, type_ string, ) *CreateInputParameterRequest`

NewCreateInputParameterRequest instantiates a new CreateInputParameterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInputParameterRequestWithDefaults

`func NewCreateInputParameterRequestWithDefaults() *CreateInputParameterRequest`

NewCreateInputParameterRequestWithDefaults instantiates a new CreateInputParameterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *CreateInputParameterRequest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CreateInputParameterRequest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CreateInputParameterRequest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CreateInputParameterRequest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDependentResourceId

`func (o *CreateInputParameterRequest) GetDependentResourceId() string`

GetDependentResourceId returns the DependentResourceId field if non-nil, zero value otherwise.

### GetDependentResourceIdOk

`func (o *CreateInputParameterRequest) GetDependentResourceIdOk() (*string, bool)`

GetDependentResourceIdOk returns a tuple with the DependentResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentResourceId

`func (o *CreateInputParameterRequest) SetDependentResourceId(v string)`

SetDependentResourceId sets DependentResourceId field to given value.

### HasDependentResourceId

`func (o *CreateInputParameterRequest) HasDependentResourceId() bool`

HasDependentResourceId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateInputParameterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInputParameterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInputParameterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHasOptions

`func (o *CreateInputParameterRequest) GetHasOptions() bool`

GetHasOptions returns the HasOptions field if non-nil, zero value otherwise.

### GetHasOptionsOk

`func (o *CreateInputParameterRequest) GetHasOptionsOk() (*bool, bool)`

GetHasOptionsOk returns a tuple with the HasOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOptions

`func (o *CreateInputParameterRequest) SetHasOptions(v bool)`

SetHasOptions sets HasOptions field to given value.

### HasHasOptions

`func (o *CreateInputParameterRequest) HasHasOptions() bool`

HasHasOptions returns a boolean if a field has been set.

### GetIsList

`func (o *CreateInputParameterRequest) GetIsList() bool`

GetIsList returns the IsList field if non-nil, zero value otherwise.

### GetIsListOk

`func (o *CreateInputParameterRequest) GetIsListOk() (*bool, bool)`

GetIsListOk returns a tuple with the IsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsList

`func (o *CreateInputParameterRequest) SetIsList(v bool)`

SetIsList sets IsList field to given value.

### HasIsList

`func (o *CreateInputParameterRequest) HasIsList() bool`

HasIsList returns a boolean if a field has been set.

### GetKey

`func (o *CreateInputParameterRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateInputParameterRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateInputParameterRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabeledOptions

`func (o *CreateInputParameterRequest) GetLabeledOptions() map[string]string`

GetLabeledOptions returns the LabeledOptions field if non-nil, zero value otherwise.

### GetLabeledOptionsOk

`func (o *CreateInputParameterRequest) GetLabeledOptionsOk() (*map[string]string, bool)`

GetLabeledOptionsOk returns a tuple with the LabeledOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabeledOptions

`func (o *CreateInputParameterRequest) SetLabeledOptions(v map[string]string)`

SetLabeledOptions sets LabeledOptions field to given value.

### HasLabeledOptions

`func (o *CreateInputParameterRequest) HasLabeledOptions() bool`

HasLabeledOptions returns a boolean if a field has been set.

### GetLimits

`func (o *CreateInputParameterRequest) GetLimits() Limits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CreateInputParameterRequest) GetLimitsOk() (*Limits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *CreateInputParameterRequest) SetLimits(v Limits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *CreateInputParameterRequest) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetModifiable

`func (o *CreateInputParameterRequest) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *CreateInputParameterRequest) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *CreateInputParameterRequest) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.


### GetName

`func (o *CreateInputParameterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInputParameterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInputParameterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *CreateInputParameterRequest) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateInputParameterRequest) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateInputParameterRequest) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateInputParameterRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRegex

`func (o *CreateInputParameterRequest) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *CreateInputParameterRequest) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *CreateInputParameterRequest) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *CreateInputParameterRequest) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetRequired

`func (o *CreateInputParameterRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CreateInputParameterRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CreateInputParameterRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetResourceId

`func (o *CreateInputParameterRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateInputParameterRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateInputParameterRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetScope

`func (o *CreateInputParameterRequest) GetScope() InputParameterScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateInputParameterRequest) GetScopeOk() (*InputParameterScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateInputParameterRequest) SetScope(v InputParameterScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateInputParameterRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateInputParameterRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateInputParameterRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateInputParameterRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTabIndex

`func (o *CreateInputParameterRequest) GetTabIndex() int64`

GetTabIndex returns the TabIndex field if non-nil, zero value otherwise.

### GetTabIndexOk

`func (o *CreateInputParameterRequest) GetTabIndexOk() (*int64, bool)`

GetTabIndexOk returns a tuple with the TabIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabIndex

`func (o *CreateInputParameterRequest) SetTabIndex(v int64)`

SetTabIndex sets TabIndex field to given value.

### HasTabIndex

`func (o *CreateInputParameterRequest) HasTabIndex() bool`

HasTabIndex returns a boolean if a field has been set.

### GetToken

`func (o *CreateInputParameterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateInputParameterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateInputParameterRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *CreateInputParameterRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateInputParameterRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateInputParameterRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


