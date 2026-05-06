# OnboardingResourceInputVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | The default value. | [optional] 
**Key** | **string** | The variable key. | 
**SourceInputVariableName** | Pointer to **string** | Source input variable name for cross-resource references. | [optional] 
**SourceResourceName** | Pointer to **string** | Source resource name for cross-resource references. | [optional] 

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

`func (o *OnboardingResourceInputVariable) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *OnboardingResourceInputVariable) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *OnboardingResourceInputVariable) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *OnboardingResourceInputVariable) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


