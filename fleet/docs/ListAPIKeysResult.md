# ListAPIKeysResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | [**[]APIKeyMetadata**](APIKeyMetadata.md) | The list of API keys for the caller&#39;s org. Every entry includes lastUsedAt (nullable) so callers can sort, filter, and surface staleness without a follow-up describe call. Returns metadata only — no plaintext, no hashes. | 

## Methods

### NewListAPIKeysResult

`func NewListAPIKeysResult(apiKeys []APIKeyMetadata, ) *ListAPIKeysResult`

NewListAPIKeysResult instantiates a new ListAPIKeysResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAPIKeysResultWithDefaults

`func NewListAPIKeysResultWithDefaults() *ListAPIKeysResult`

NewListAPIKeysResultWithDefaults instantiates a new ListAPIKeysResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *ListAPIKeysResult) GetApiKeys() []APIKeyMetadata`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *ListAPIKeysResult) GetApiKeysOk() (*[]APIKeyMetadata, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *ListAPIKeysResult) SetApiKeys(v []APIKeyMetadata)`

SetApiKeys sets ApiKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


