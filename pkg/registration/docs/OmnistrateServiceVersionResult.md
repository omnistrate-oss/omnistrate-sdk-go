# OmnistrateServiceVersionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Omnistrate API version | 
**BuildCommitSHA** | **string** | The commit hash of the build | 
**BuildTimestamp** | **string** | The timestamp of the build | 

## Methods

### NewOmnistrateServiceVersionResult

`func NewOmnistrateServiceVersionResult(apiVersion string, buildCommitSHA string, buildTimestamp string, ) *OmnistrateServiceVersionResult`

NewOmnistrateServiceVersionResult instantiates a new OmnistrateServiceVersionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmnistrateServiceVersionResultWithDefaults

`func NewOmnistrateServiceVersionResultWithDefaults() *OmnistrateServiceVersionResult`

NewOmnistrateServiceVersionResultWithDefaults instantiates a new OmnistrateServiceVersionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OmnistrateServiceVersionResult) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OmnistrateServiceVersionResult) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OmnistrateServiceVersionResult) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetBuildCommitSHA

`func (o *OmnistrateServiceVersionResult) GetBuildCommitSHA() string`

GetBuildCommitSHA returns the BuildCommitSHA field if non-nil, zero value otherwise.

### GetBuildCommitSHAOk

`func (o *OmnistrateServiceVersionResult) GetBuildCommitSHAOk() (*string, bool)`

GetBuildCommitSHAOk returns a tuple with the BuildCommitSHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCommitSHA

`func (o *OmnistrateServiceVersionResult) SetBuildCommitSHA(v string)`

SetBuildCommitSHA sets BuildCommitSHA field to given value.


### GetBuildTimestamp

`func (o *OmnistrateServiceVersionResult) GetBuildTimestamp() string`

GetBuildTimestamp returns the BuildTimestamp field if non-nil, zero value otherwise.

### GetBuildTimestampOk

`func (o *OmnistrateServiceVersionResult) GetBuildTimestampOk() (*string, bool)`

GetBuildTimestampOk returns a tuple with the BuildTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimestamp

`func (o *OmnistrateServiceVersionResult) SetBuildTimestamp(v string)`

SetBuildTimestamp sets BuildTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


