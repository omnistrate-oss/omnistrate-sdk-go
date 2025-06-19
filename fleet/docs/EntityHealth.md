# EntityHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Optional error message if the entity is not in sync | [optional] 
**Identifier** | **string** | Unique identifier for the entity | 
**SyncStatus** | **string** | Current synchronization status of the entity | 
**Type** | **string** | Type of the entity (e.g., NAMESPACE, SERVICE, POD, etc.) | 

## Methods

### NewEntityHealth

`func NewEntityHealth(identifier string, syncStatus string, type_ string, ) *EntityHealth`

NewEntityHealth instantiates a new EntityHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityHealthWithDefaults

`func NewEntityHealthWithDefaults() *EntityHealth`

NewEntityHealthWithDefaults instantiates a new EntityHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *EntityHealth) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EntityHealth) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EntityHealth) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EntityHealth) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIdentifier

`func (o *EntityHealth) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EntityHealth) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EntityHealth) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSyncStatus

`func (o *EntityHealth) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *EntityHealth) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *EntityHealth) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.


### GetType

`func (o *EntityHealth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityHealth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityHealth) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


