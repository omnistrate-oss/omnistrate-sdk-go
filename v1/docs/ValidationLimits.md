# ValidationLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxArchiveEntries** | **int64** | Maximum number of archive entries across all supplied artifacts | 
**MaxArchiveMemberPathBytes** | **int64** | Maximum length in UTF-8 bytes of a canonical archive member path | 
**MaxArtifacts** | **int64** | Maximum number of distinct supplied artifacts | 
**MaxConcurrentContentValidations** | **int64** | Number of simultaneous content-heavy validations admitted per server process. Requests beyond this bound are rejected rather than queued without bound. | 
**MaxRequestBodyBytes** | **int64** | Maximum HTTP JSON request body size, enforced before the body is decoded | 
**MaxTotalCompressedArtifactBytes** | **int64** | Maximum sum of decoded compressed archive bytes across all supplied artifacts | 
**MaxTotalExtractedBytes** | **int64** | Maximum sum of extracted regular-file bytes across all supplied artifacts | 
**MaxTotalSpecBytes** | **int64** | Maximum sum of the decoded specification YAML plus decoded configs plus decoded secrets | 
**RequestDeadlineSeconds** | **int64** | Overall request deadline in seconds | 

## Methods

### NewValidationLimits

`func NewValidationLimits(maxArchiveEntries int64, maxArchiveMemberPathBytes int64, maxArtifacts int64, maxConcurrentContentValidations int64, maxRequestBodyBytes int64, maxTotalCompressedArtifactBytes int64, maxTotalExtractedBytes int64, maxTotalSpecBytes int64, requestDeadlineSeconds int64, ) *ValidationLimits`

NewValidationLimits instantiates a new ValidationLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationLimitsWithDefaults

`func NewValidationLimitsWithDefaults() *ValidationLimits`

NewValidationLimitsWithDefaults instantiates a new ValidationLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxArchiveEntries

`func (o *ValidationLimits) GetMaxArchiveEntries() int64`

GetMaxArchiveEntries returns the MaxArchiveEntries field if non-nil, zero value otherwise.

### GetMaxArchiveEntriesOk

`func (o *ValidationLimits) GetMaxArchiveEntriesOk() (*int64, bool)`

GetMaxArchiveEntriesOk returns a tuple with the MaxArchiveEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxArchiveEntries

`func (o *ValidationLimits) SetMaxArchiveEntries(v int64)`

SetMaxArchiveEntries sets MaxArchiveEntries field to given value.


### GetMaxArchiveMemberPathBytes

`func (o *ValidationLimits) GetMaxArchiveMemberPathBytes() int64`

GetMaxArchiveMemberPathBytes returns the MaxArchiveMemberPathBytes field if non-nil, zero value otherwise.

### GetMaxArchiveMemberPathBytesOk

`func (o *ValidationLimits) GetMaxArchiveMemberPathBytesOk() (*int64, bool)`

GetMaxArchiveMemberPathBytesOk returns a tuple with the MaxArchiveMemberPathBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxArchiveMemberPathBytes

`func (o *ValidationLimits) SetMaxArchiveMemberPathBytes(v int64)`

SetMaxArchiveMemberPathBytes sets MaxArchiveMemberPathBytes field to given value.


### GetMaxArtifacts

`func (o *ValidationLimits) GetMaxArtifacts() int64`

GetMaxArtifacts returns the MaxArtifacts field if non-nil, zero value otherwise.

### GetMaxArtifactsOk

`func (o *ValidationLimits) GetMaxArtifactsOk() (*int64, bool)`

GetMaxArtifactsOk returns a tuple with the MaxArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxArtifacts

`func (o *ValidationLimits) SetMaxArtifacts(v int64)`

SetMaxArtifacts sets MaxArtifacts field to given value.


### GetMaxConcurrentContentValidations

`func (o *ValidationLimits) GetMaxConcurrentContentValidations() int64`

GetMaxConcurrentContentValidations returns the MaxConcurrentContentValidations field if non-nil, zero value otherwise.

### GetMaxConcurrentContentValidationsOk

`func (o *ValidationLimits) GetMaxConcurrentContentValidationsOk() (*int64, bool)`

GetMaxConcurrentContentValidationsOk returns a tuple with the MaxConcurrentContentValidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentContentValidations

`func (o *ValidationLimits) SetMaxConcurrentContentValidations(v int64)`

SetMaxConcurrentContentValidations sets MaxConcurrentContentValidations field to given value.


### GetMaxRequestBodyBytes

`func (o *ValidationLimits) GetMaxRequestBodyBytes() int64`

GetMaxRequestBodyBytes returns the MaxRequestBodyBytes field if non-nil, zero value otherwise.

### GetMaxRequestBodyBytesOk

`func (o *ValidationLimits) GetMaxRequestBodyBytesOk() (*int64, bool)`

GetMaxRequestBodyBytesOk returns a tuple with the MaxRequestBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestBodyBytes

`func (o *ValidationLimits) SetMaxRequestBodyBytes(v int64)`

SetMaxRequestBodyBytes sets MaxRequestBodyBytes field to given value.


### GetMaxTotalCompressedArtifactBytes

`func (o *ValidationLimits) GetMaxTotalCompressedArtifactBytes() int64`

GetMaxTotalCompressedArtifactBytes returns the MaxTotalCompressedArtifactBytes field if non-nil, zero value otherwise.

### GetMaxTotalCompressedArtifactBytesOk

`func (o *ValidationLimits) GetMaxTotalCompressedArtifactBytesOk() (*int64, bool)`

GetMaxTotalCompressedArtifactBytesOk returns a tuple with the MaxTotalCompressedArtifactBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalCompressedArtifactBytes

`func (o *ValidationLimits) SetMaxTotalCompressedArtifactBytes(v int64)`

SetMaxTotalCompressedArtifactBytes sets MaxTotalCompressedArtifactBytes field to given value.


### GetMaxTotalExtractedBytes

`func (o *ValidationLimits) GetMaxTotalExtractedBytes() int64`

GetMaxTotalExtractedBytes returns the MaxTotalExtractedBytes field if non-nil, zero value otherwise.

### GetMaxTotalExtractedBytesOk

`func (o *ValidationLimits) GetMaxTotalExtractedBytesOk() (*int64, bool)`

GetMaxTotalExtractedBytesOk returns a tuple with the MaxTotalExtractedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalExtractedBytes

`func (o *ValidationLimits) SetMaxTotalExtractedBytes(v int64)`

SetMaxTotalExtractedBytes sets MaxTotalExtractedBytes field to given value.


### GetMaxTotalSpecBytes

`func (o *ValidationLimits) GetMaxTotalSpecBytes() int64`

GetMaxTotalSpecBytes returns the MaxTotalSpecBytes field if non-nil, zero value otherwise.

### GetMaxTotalSpecBytesOk

`func (o *ValidationLimits) GetMaxTotalSpecBytesOk() (*int64, bool)`

GetMaxTotalSpecBytesOk returns a tuple with the MaxTotalSpecBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalSpecBytes

`func (o *ValidationLimits) SetMaxTotalSpecBytes(v int64)`

SetMaxTotalSpecBytes sets MaxTotalSpecBytes field to given value.


### GetRequestDeadlineSeconds

`func (o *ValidationLimits) GetRequestDeadlineSeconds() int64`

GetRequestDeadlineSeconds returns the RequestDeadlineSeconds field if non-nil, zero value otherwise.

### GetRequestDeadlineSecondsOk

`func (o *ValidationLimits) GetRequestDeadlineSecondsOk() (*int64, bool)`

GetRequestDeadlineSecondsOk returns a tuple with the RequestDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDeadlineSeconds

`func (o *ValidationLimits) SetRequestDeadlineSeconds(v int64)`

SetRequestDeadlineSeconds sets RequestDeadlineSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


