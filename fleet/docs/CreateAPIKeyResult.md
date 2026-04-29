# CreateAPIKeyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of an API Key | 
**Key** | **string** | The plaintext API key. Returned exactly once on creation and never persisted in this form. Callers MUST capture and store this value at creation time; it cannot be retrieved later. | 
**Metadata** | [**APIKeyMetadata**](APIKeyMetadata.md) |  | 

## Methods

### NewCreateAPIKeyResult

`func NewCreateAPIKeyResult(id string, key string, metadata APIKeyMetadata, ) *CreateAPIKeyResult`

NewCreateAPIKeyResult instantiates a new CreateAPIKeyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAPIKeyResultWithDefaults

`func NewCreateAPIKeyResultWithDefaults() *CreateAPIKeyResult`

NewCreateAPIKeyResultWithDefaults instantiates a new CreateAPIKeyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAPIKeyResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAPIKeyResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAPIKeyResult) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *CreateAPIKeyResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateAPIKeyResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateAPIKeyResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetMetadata

`func (o *CreateAPIKeyResult) GetMetadata() APIKeyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateAPIKeyResult) GetMetadataOk() (*APIKeyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateAPIKeyResult) SetMetadata(v APIKeyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


