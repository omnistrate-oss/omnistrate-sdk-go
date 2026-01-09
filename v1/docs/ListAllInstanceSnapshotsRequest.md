# ListAllInstanceSnapshotsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**SnapshotType** | Pointer to **string** | The snapshot type | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAllInstanceSnapshotsRequest

`func NewListAllInstanceSnapshotsRequest(token string, ) *ListAllInstanceSnapshotsRequest`

NewListAllInstanceSnapshotsRequest instantiates a new ListAllInstanceSnapshotsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllInstanceSnapshotsRequestWithDefaults

`func NewListAllInstanceSnapshotsRequestWithDefaults() *ListAllInstanceSnapshotsRequest`

NewListAllInstanceSnapshotsRequestWithDefaults instantiates a new ListAllInstanceSnapshotsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *ListAllInstanceSnapshotsRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ListAllInstanceSnapshotsRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ListAllInstanceSnapshotsRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ListAllInstanceSnapshotsRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetSnapshotType

`func (o *ListAllInstanceSnapshotsRequest) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *ListAllInstanceSnapshotsRequest) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *ListAllInstanceSnapshotsRequest) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *ListAllInstanceSnapshotsRequest) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetToken

`func (o *ListAllInstanceSnapshotsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAllInstanceSnapshotsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAllInstanceSnapshotsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


