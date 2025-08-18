# DebugHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Host Cluster | 
**IncludeAmenitiesInstallationLogs** | Pointer to **bool** | Whether to include the installation logs of the amenities in the debug information | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDebugHostClusterRequest

`func NewDebugHostClusterRequest(id string, token string, ) *DebugHostClusterRequest`

NewDebugHostClusterRequest instantiates a new DebugHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugHostClusterRequestWithDefaults

`func NewDebugHostClusterRequestWithDefaults() *DebugHostClusterRequest`

NewDebugHostClusterRequestWithDefaults instantiates a new DebugHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DebugHostClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DebugHostClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DebugHostClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIncludeAmenitiesInstallationLogs

`func (o *DebugHostClusterRequest) GetIncludeAmenitiesInstallationLogs() bool`

GetIncludeAmenitiesInstallationLogs returns the IncludeAmenitiesInstallationLogs field if non-nil, zero value otherwise.

### GetIncludeAmenitiesInstallationLogsOk

`func (o *DebugHostClusterRequest) GetIncludeAmenitiesInstallationLogsOk() (*bool, bool)`

GetIncludeAmenitiesInstallationLogsOk returns a tuple with the IncludeAmenitiesInstallationLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAmenitiesInstallationLogs

`func (o *DebugHostClusterRequest) SetIncludeAmenitiesInstallationLogs(v bool)`

SetIncludeAmenitiesInstallationLogs sets IncludeAmenitiesInstallationLogs field to given value.

### HasIncludeAmenitiesInstallationLogs

`func (o *DebugHostClusterRequest) HasIncludeAmenitiesInstallationLogs() bool`

HasIncludeAmenitiesInstallationLogs returns a boolean if a field has been set.

### GetToken

`func (o *DebugHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DebugHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DebugHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


