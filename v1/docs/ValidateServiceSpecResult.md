# ValidateServiceSpecResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | [**[]ValidationCheck**](ValidationCheck.md) | The outcome of every validation check category. Every category is always present. Always serialized as an array, using an empty array rather than null. | 
**Diagnostics** | [**[]ValidationDiagnostic**](ValidationDiagnostic.md) | Every observation produced by the checks, in a stable order. Always serialized as an array, using an empty array rather than null. | 
**ExistingTarget** | Pointer to [**ValidationExistingTarget**](ValidationExistingTarget.md) |  | [optional] 
**InputDigest** | **string** | Lowercase hex SHA-256 over a deterministic representation of the effective request fields, excluding the token and the supplied artifacts. It lets a client confirm that a discovery response and a final response describe the same input. It is not a service revision lock and not an authorization token. | 
**Limits** | Pointer to [**ValidationLimits**](ValidationLimits.md) |  | [optional] 
**ObservedAt** | Pointer to **time.Time** | When the read-only snapshot of existing state was acquired. Omitted when validation stopped before any snapshot was read. It is not a promise that a later build observes that state. | [optional] 
**RequiredArtifacts** | [**[]ArtifactRequirement**](ArtifactRequirement.md) | Local artifact content the submitted specification needs and that this request did not supply, sorted by canonical path. Always serialized as an array, using an empty array rather than null. | 
**Status** | **string** | The overall outcome of a specification validation. VALID means every applicable, currently supported build-time check completed successfully; it is not a promise that a later deployment succeeds. INVALID means at least one diagnostic has severity &#39;error&#39;. INCOMPLETE means no error was found but at least one required check could not be completed, for example because local artifact content, a remote dependency or a supported validation implementation was unavailable. An unfinished check is never reported as VALID. | 
**ValidatedArtifacts** | [**[]ValidatedArtifact**](ValidatedArtifact.md) | Supplied artifact content whose applicable content checks finished, sorted by canonical path. Always serialized as an array, using an empty array rather than null. | 
**ValidationVersion** | **string** | Version of the validation contract that produced this result | 

## Methods

### NewValidateServiceSpecResult

`func NewValidateServiceSpecResult(checks []ValidationCheck, diagnostics []ValidationDiagnostic, inputDigest string, requiredArtifacts []ArtifactRequirement, status string, validatedArtifacts []ValidatedArtifact, validationVersion string, ) *ValidateServiceSpecResult`

NewValidateServiceSpecResult instantiates a new ValidateServiceSpecResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateServiceSpecResultWithDefaults

`func NewValidateServiceSpecResultWithDefaults() *ValidateServiceSpecResult`

NewValidateServiceSpecResultWithDefaults instantiates a new ValidateServiceSpecResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *ValidateServiceSpecResult) GetChecks() []ValidationCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ValidateServiceSpecResult) GetChecksOk() (*[]ValidationCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ValidateServiceSpecResult) SetChecks(v []ValidationCheck)`

SetChecks sets Checks field to given value.


### GetDiagnostics

`func (o *ValidateServiceSpecResult) GetDiagnostics() []ValidationDiagnostic`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *ValidateServiceSpecResult) GetDiagnosticsOk() (*[]ValidationDiagnostic, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *ValidateServiceSpecResult) SetDiagnostics(v []ValidationDiagnostic)`

SetDiagnostics sets Diagnostics field to given value.


### GetExistingTarget

`func (o *ValidateServiceSpecResult) GetExistingTarget() ValidationExistingTarget`

GetExistingTarget returns the ExistingTarget field if non-nil, zero value otherwise.

### GetExistingTargetOk

`func (o *ValidateServiceSpecResult) GetExistingTargetOk() (*ValidationExistingTarget, bool)`

GetExistingTargetOk returns a tuple with the ExistingTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingTarget

`func (o *ValidateServiceSpecResult) SetExistingTarget(v ValidationExistingTarget)`

SetExistingTarget sets ExistingTarget field to given value.

### HasExistingTarget

`func (o *ValidateServiceSpecResult) HasExistingTarget() bool`

HasExistingTarget returns a boolean if a field has been set.

### GetInputDigest

`func (o *ValidateServiceSpecResult) GetInputDigest() string`

GetInputDigest returns the InputDigest field if non-nil, zero value otherwise.

### GetInputDigestOk

`func (o *ValidateServiceSpecResult) GetInputDigestOk() (*string, bool)`

GetInputDigestOk returns a tuple with the InputDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDigest

`func (o *ValidateServiceSpecResult) SetInputDigest(v string)`

SetInputDigest sets InputDigest field to given value.


### GetLimits

`func (o *ValidateServiceSpecResult) GetLimits() ValidationLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ValidateServiceSpecResult) GetLimitsOk() (*ValidationLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ValidateServiceSpecResult) SetLimits(v ValidationLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ValidateServiceSpecResult) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetObservedAt

`func (o *ValidateServiceSpecResult) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *ValidateServiceSpecResult) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *ValidateServiceSpecResult) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *ValidateServiceSpecResult) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### GetRequiredArtifacts

`func (o *ValidateServiceSpecResult) GetRequiredArtifacts() []ArtifactRequirement`

GetRequiredArtifacts returns the RequiredArtifacts field if non-nil, zero value otherwise.

### GetRequiredArtifactsOk

`func (o *ValidateServiceSpecResult) GetRequiredArtifactsOk() (*[]ArtifactRequirement, bool)`

GetRequiredArtifactsOk returns a tuple with the RequiredArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredArtifacts

`func (o *ValidateServiceSpecResult) SetRequiredArtifacts(v []ArtifactRequirement)`

SetRequiredArtifacts sets RequiredArtifacts field to given value.


### GetStatus

`func (o *ValidateServiceSpecResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidateServiceSpecResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidateServiceSpecResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetValidatedArtifacts

`func (o *ValidateServiceSpecResult) GetValidatedArtifacts() []ValidatedArtifact`

GetValidatedArtifacts returns the ValidatedArtifacts field if non-nil, zero value otherwise.

### GetValidatedArtifactsOk

`func (o *ValidateServiceSpecResult) GetValidatedArtifactsOk() (*[]ValidatedArtifact, bool)`

GetValidatedArtifactsOk returns a tuple with the ValidatedArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedArtifacts

`func (o *ValidateServiceSpecResult) SetValidatedArtifacts(v []ValidatedArtifact)`

SetValidatedArtifacts sets ValidatedArtifacts field to given value.


### GetValidationVersion

`func (o *ValidateServiceSpecResult) GetValidationVersion() string`

GetValidationVersion returns the ValidationVersion field if non-nil, zero value otherwise.

### GetValidationVersionOk

`func (o *ValidateServiceSpecResult) GetValidationVersionOk() (*string, bool)`

GetValidationVersionOk returns a tuple with the ValidationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationVersion

`func (o *ValidateServiceSpecResult) SetValidationVersion(v string)`

SetValidationVersion sets ValidationVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


