# DetailedNodeHealthResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectivityStatus** | Pointer to **string** | The heath status of a resource | [optional] 
**DiskHealth** | Pointer to **string** | The heath status of a resource | [optional] 
**LoadHealth** | Pointer to **string** | The load status of a pod | [optional] 
**NodeHealth** | Pointer to **string** | The heath status of a resource | [optional] 
**ProcessHealth** | Pointer to **string** | The heath status of a resource | [optional] 
**ProcessLiveness** | Pointer to **string** | The heath status of a resource | [optional] 

## Methods

### NewDetailedNodeHealthResult

`func NewDetailedNodeHealthResult() *DetailedNodeHealthResult`

NewDetailedNodeHealthResult instantiates a new DetailedNodeHealthResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedNodeHealthResultWithDefaults

`func NewDetailedNodeHealthResultWithDefaults() *DetailedNodeHealthResult`

NewDetailedNodeHealthResultWithDefaults instantiates a new DetailedNodeHealthResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectivityStatus

`func (o *DetailedNodeHealthResult) GetConnectivityStatus() string`

GetConnectivityStatus returns the ConnectivityStatus field if non-nil, zero value otherwise.

### GetConnectivityStatusOk

`func (o *DetailedNodeHealthResult) GetConnectivityStatusOk() (*string, bool)`

GetConnectivityStatusOk returns a tuple with the ConnectivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityStatus

`func (o *DetailedNodeHealthResult) SetConnectivityStatus(v string)`

SetConnectivityStatus sets ConnectivityStatus field to given value.

### HasConnectivityStatus

`func (o *DetailedNodeHealthResult) HasConnectivityStatus() bool`

HasConnectivityStatus returns a boolean if a field has been set.

### GetDiskHealth

`func (o *DetailedNodeHealthResult) GetDiskHealth() string`

GetDiskHealth returns the DiskHealth field if non-nil, zero value otherwise.

### GetDiskHealthOk

`func (o *DetailedNodeHealthResult) GetDiskHealthOk() (*string, bool)`

GetDiskHealthOk returns a tuple with the DiskHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskHealth

`func (o *DetailedNodeHealthResult) SetDiskHealth(v string)`

SetDiskHealth sets DiskHealth field to given value.

### HasDiskHealth

`func (o *DetailedNodeHealthResult) HasDiskHealth() bool`

HasDiskHealth returns a boolean if a field has been set.

### GetLoadHealth

`func (o *DetailedNodeHealthResult) GetLoadHealth() string`

GetLoadHealth returns the LoadHealth field if non-nil, zero value otherwise.

### GetLoadHealthOk

`func (o *DetailedNodeHealthResult) GetLoadHealthOk() (*string, bool)`

GetLoadHealthOk returns a tuple with the LoadHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadHealth

`func (o *DetailedNodeHealthResult) SetLoadHealth(v string)`

SetLoadHealth sets LoadHealth field to given value.

### HasLoadHealth

`func (o *DetailedNodeHealthResult) HasLoadHealth() bool`

HasLoadHealth returns a boolean if a field has been set.

### GetNodeHealth

`func (o *DetailedNodeHealthResult) GetNodeHealth() string`

GetNodeHealth returns the NodeHealth field if non-nil, zero value otherwise.

### GetNodeHealthOk

`func (o *DetailedNodeHealthResult) GetNodeHealthOk() (*string, bool)`

GetNodeHealthOk returns a tuple with the NodeHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeHealth

`func (o *DetailedNodeHealthResult) SetNodeHealth(v string)`

SetNodeHealth sets NodeHealth field to given value.

### HasNodeHealth

`func (o *DetailedNodeHealthResult) HasNodeHealth() bool`

HasNodeHealth returns a boolean if a field has been set.

### GetProcessHealth

`func (o *DetailedNodeHealthResult) GetProcessHealth() string`

GetProcessHealth returns the ProcessHealth field if non-nil, zero value otherwise.

### GetProcessHealthOk

`func (o *DetailedNodeHealthResult) GetProcessHealthOk() (*string, bool)`

GetProcessHealthOk returns a tuple with the ProcessHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessHealth

`func (o *DetailedNodeHealthResult) SetProcessHealth(v string)`

SetProcessHealth sets ProcessHealth field to given value.

### HasProcessHealth

`func (o *DetailedNodeHealthResult) HasProcessHealth() bool`

HasProcessHealth returns a boolean if a field has been set.

### GetProcessLiveness

`func (o *DetailedNodeHealthResult) GetProcessLiveness() string`

GetProcessLiveness returns the ProcessLiveness field if non-nil, zero value otherwise.

### GetProcessLivenessOk

`func (o *DetailedNodeHealthResult) GetProcessLivenessOk() (*string, bool)`

GetProcessLivenessOk returns a tuple with the ProcessLiveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessLiveness

`func (o *DetailedNodeHealthResult) SetProcessLiveness(v string)`

SetProcessLiveness sets ProcessLiveness field to given value.

### HasProcessLiveness

`func (o *DetailedNodeHealthResult) HasProcessLiveness() bool`

HasProcessLiveness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


