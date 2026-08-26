# OnboardingResourceInputVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **interface{}** | The typed override value, when one is provided. | [optional] 
**InitialValue** | Pointer to **interface{}** | The typed initial value inferred from the artifact default, when one exists. | [optional] 
**Key** | **string** | The variable key. | 
**Required** | Pointer to **bool** | Whether the input variable must be supplied because the artifact has no default value. | [optional] 
**SourceInputVariableName** | Pointer to **string** | Source input variable name for cross-resource references. | [optional] 
**SourceResourceName** | Pointer to **string** | Source resource name for cross-resource references. | [optional] 
**Type** | Pointer to **string** | Normalized input variable type inferred from the onboarding artifact. | [optional] 

## Methods

### NewOnboardingResourceInputVariable

`func NewOnboardingResourceInputVariable(key string, ) *OnboardingResourceInputVariable`

NewOnboardingResourceInputVariable instantiates a new OnboardingResourceInputVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingResourceInputVariableWithDefaults

`func NewOnboardingResourceInputVariableWithDefaults() *OnboardingResourceInputVariable`

NewOnboardingResourceInputVariableWithDefaults instantiates a new OnboardingResourceInputVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *OnboardingResourceInputVariable) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *OnboardingResourceInputVariable) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *OnboardingResourceInputVariable) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *OnboardingResourceInputVariable) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *OnboardingResourceInputVariable) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *OnboardingResourceInputVariable) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetInitialValue

`func (o *OnboardingResourceInputVariable) GetInitialValue() interface{}`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *OnboardingResourceInputVariable) GetInitialValueOk() (*interface{}, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *OnboardingResourceInputVariable) SetInitialValue(v interface{})`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *OnboardingResourceInputVariable) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### SetInitialValueNil

`func (o *OnboardingResourceInputVariable) SetInitialValueNil(b bool)`

 SetInitialValueNil sets the value for InitialValue to be an explicit nil

### UnsetInitialValue
`func (o *OnboardingResourceInputVariable) UnsetInitialValue()`

UnsetInitialValue ensures that no value is present for InitialValue, not even an explicit nil
### GetKey

`func (o *OnboardingResourceInputVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OnboardingResourceInputVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OnboardingResourceInputVariable) SetKey(v string)`

SetKey sets Key field to given value.


### GetRequired

`func (o *OnboardingResourceInputVariable) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *OnboardingResourceInputVariable) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *OnboardingResourceInputVariable) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *OnboardingResourceInputVariable) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSourceInputVariableName

`func (o *OnboardingResourceInputVariable) GetSourceInputVariableName() string`

GetSourceInputVariableName returns the SourceInputVariableName field if non-nil, zero value otherwise.

### GetSourceInputVariableNameOk

`func (o *OnboardingResourceInputVariable) GetSourceInputVariableNameOk() (*string, bool)`

GetSourceInputVariableNameOk returns a tuple with the SourceInputVariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInputVariableName

`func (o *OnboardingResourceInputVariable) SetSourceInputVariableName(v string)`

SetSourceInputVariableName sets SourceInputVariableName field to given value.

### HasSourceInputVariableName

`func (o *OnboardingResourceInputVariable) HasSourceInputVariableName() bool`

HasSourceInputVariableName returns a boolean if a field has been set.

### GetSourceResourceName

`func (o *OnboardingResourceInputVariable) GetSourceResourceName() string`

GetSourceResourceName returns the SourceResourceName field if non-nil, zero value otherwise.

### GetSourceResourceNameOk

`func (o *OnboardingResourceInputVariable) GetSourceResourceNameOk() (*string, bool)`

GetSourceResourceNameOk returns a tuple with the SourceResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceResourceName

`func (o *OnboardingResourceInputVariable) SetSourceResourceName(v string)`

SetSourceResourceName sets SourceResourceName field to given value.

### HasSourceResourceName

`func (o *OnboardingResourceInputVariable) HasSourceResourceName() bool`

HasSourceResourceName returns a boolean if a field has been set.

### GetType

`func (o *OnboardingResourceInputVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnboardingResourceInputVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnboardingResourceInputVariable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OnboardingResourceInputVariable) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


