# DebugHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomHelmExecutionLogs** | Pointer to **map[string]*os.File** | Custom Helm execution logs for the host cluster, keyed by namespace | [optional] 

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

### GetCustomHelmExecutionLogs

`func (o *DebugHostClusterResult) GetCustomHelmExecutionLogs() map[string]*os.File`

GetCustomHelmExecutionLogs returns the CustomHelmExecutionLogs field if non-nil, zero value otherwise.

### GetCustomHelmExecutionLogsOk

`func (o *DebugHostClusterResult) GetCustomHelmExecutionLogsOk() (*map[string]*os.File, bool)`

GetCustomHelmExecutionLogsOk returns a tuple with the CustomHelmExecutionLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHelmExecutionLogs

`func (o *DebugHostClusterResult) SetCustomHelmExecutionLogs(v map[string]*os.File)`

SetCustomHelmExecutionLogs sets CustomHelmExecutionLogs field to given value.

### HasCustomHelmExecutionLogs

`func (o *DebugHostClusterResult) HasCustomHelmExecutionLogs() bool`

HasCustomHelmExecutionLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


