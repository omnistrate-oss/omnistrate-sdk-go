# DebugHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomHelmExecutionLogsBase64** | Pointer to **map[string]string** | Custom Helm execution logs for the host cluster, keyed by namespace | [optional] 

## Methods

### NewDebugHostClusterResult

`func NewDebugHostClusterResult() *DebugHostClusterResult`

NewDebugHostClusterResult instantiates a new DebugHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugHostClusterResultWithDefaults

`func NewDebugHostClusterResultWithDefaults() *DebugHostClusterResult`

NewDebugHostClusterResultWithDefaults instantiates a new DebugHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomHelmExecutionLogsBase64

`func (o *DebugHostClusterResult) GetCustomHelmExecutionLogsBase64() map[string]string`

GetCustomHelmExecutionLogsBase64 returns the CustomHelmExecutionLogsBase64 field if non-nil, zero value otherwise.

### GetCustomHelmExecutionLogsBase64Ok

`func (o *DebugHostClusterResult) GetCustomHelmExecutionLogsBase64Ok() (*map[string]string, bool)`

GetCustomHelmExecutionLogsBase64Ok returns a tuple with the CustomHelmExecutionLogsBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHelmExecutionLogsBase64

`func (o *DebugHostClusterResult) SetCustomHelmExecutionLogsBase64(v map[string]string)`

SetCustomHelmExecutionLogsBase64 sets CustomHelmExecutionLogsBase64 field to given value.

### HasCustomHelmExecutionLogsBase64

`func (o *DebugHostClusterResult) HasCustomHelmExecutionLogsBase64() bool`

HasCustomHelmExecutionLogsBase64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


