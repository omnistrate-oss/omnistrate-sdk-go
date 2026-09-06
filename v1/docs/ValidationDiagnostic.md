# ValidationDiagnostic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Stable machine-readable diagnostic code. Callers branch on this value rather than on message text. Envelope-level failures (malformed JSON, invalid base64, wrong types, missing required fields, request limits, admission exhaustion, authentication, authorization, timeout) are reported through the typed HTTP error responses of the endpoint instead of through these diagnostics. | 
**Message** | **string** | Human readable, redacted explanation | 
**Path** | Pointer to **string** | JSON pointer into the submitted specification identifying the user input this observation is about. Omitted when the observation cannot be attributed to a specific input location. | [optional] 
**ResourceKey** | Pointer to **string** | The canonical resource key from the submitted specification this observation is about. This is a specification key, never a persisted resource ID. | [optional] 
**Severity** | **string** | The severity of a validation diagnostic. &#39;error&#39; makes the overall status INVALID. &#39;incomplete&#39; reports something that could not be validated and makes the overall status INCOMPLETE unless an error is also present. &#39;warning&#39; is informational only and never changes the overall status. Content that is required but missing is reported as &#39;incomplete&#39;, never as a warning. | 

## Methods

### NewValidationDiagnostic

`func NewValidationDiagnostic(code string, message string, severity string, ) *ValidationDiagnostic`

NewValidationDiagnostic instantiates a new ValidationDiagnostic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationDiagnosticWithDefaults

`func NewValidationDiagnosticWithDefaults() *ValidationDiagnostic`

NewValidationDiagnosticWithDefaults instantiates a new ValidationDiagnostic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ValidationDiagnostic) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationDiagnostic) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationDiagnostic) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ValidationDiagnostic) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationDiagnostic) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationDiagnostic) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPath

`func (o *ValidationDiagnostic) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ValidationDiagnostic) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ValidationDiagnostic) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ValidationDiagnostic) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetResourceKey

`func (o *ValidationDiagnostic) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ValidationDiagnostic) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ValidationDiagnostic) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.

### HasResourceKey

`func (o *ValidationDiagnostic) HasResourceKey() bool`

HasResourceKey returns a boolean if a field has been set.

### GetSeverity

`func (o *ValidationDiagnostic) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ValidationDiagnostic) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ValidationDiagnostic) SetSeverity(v string)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


