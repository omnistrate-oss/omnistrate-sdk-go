# DescribeServicesOrchestrationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicesTopology** | Pointer to [**[]ServiceDeploymentDetails**](ServiceDeploymentDetails.md) | The services deployment details | [optional] 
**CreatedAt** | **string** | The time the services orchestration was created | 
**Id** | **string** | The ID of the services orchestration | 
**OrchestrationFailedReason** | Pointer to **string** | The reason why the orchestration failed | [optional] 
**ResultParams** | Pointer to **interface{}** | Custom result parameters of the services orchestration | [optional] 
**Status** | **string** | The status of the services orchestration | 
**UpdatedAt** | **string** | The time the services orchestration was last updated | 

## Methods

### NewDescribeServicesOrchestrationResult

`func NewDescribeServicesOrchestrationResult(createdAt string, id string, status string, updatedAt string, ) *DescribeServicesOrchestrationResult`

NewDescribeServicesOrchestrationResult instantiates a new DescribeServicesOrchestrationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServicesOrchestrationResultWithDefaults

`func NewDescribeServicesOrchestrationResultWithDefaults() *DescribeServicesOrchestrationResult`

NewDescribeServicesOrchestrationResultWithDefaults instantiates a new DescribeServicesOrchestrationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicesTopology

`func (o *DescribeServicesOrchestrationResult) GetServicesTopology() []ServiceDeploymentDetails`

GetServicesTopology returns the ServicesTopology field if non-nil, zero value otherwise.

### GetServicesTopologyOk

`func (o *DescribeServicesOrchestrationResult) GetServicesTopologyOk() (*[]ServiceDeploymentDetails, bool)`

GetServicesTopologyOk returns a tuple with the ServicesTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesTopology

`func (o *DescribeServicesOrchestrationResult) SetServicesTopology(v []ServiceDeploymentDetails)`

SetServicesTopology sets ServicesTopology field to given value.

### HasServicesTopology

`func (o *DescribeServicesOrchestrationResult) HasServicesTopology() bool`

HasServicesTopology returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DescribeServicesOrchestrationResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DescribeServicesOrchestrationResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DescribeServicesOrchestrationResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *DescribeServicesOrchestrationResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeServicesOrchestrationResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeServicesOrchestrationResult) SetId(v string)`

SetId sets Id field to given value.


### GetOrchestrationFailedReason

`func (o *DescribeServicesOrchestrationResult) GetOrchestrationFailedReason() string`

GetOrchestrationFailedReason returns the OrchestrationFailedReason field if non-nil, zero value otherwise.

### GetOrchestrationFailedReasonOk

`func (o *DescribeServicesOrchestrationResult) GetOrchestrationFailedReasonOk() (*string, bool)`

GetOrchestrationFailedReasonOk returns a tuple with the OrchestrationFailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationFailedReason

`func (o *DescribeServicesOrchestrationResult) SetOrchestrationFailedReason(v string)`

SetOrchestrationFailedReason sets OrchestrationFailedReason field to given value.

### HasOrchestrationFailedReason

`func (o *DescribeServicesOrchestrationResult) HasOrchestrationFailedReason() bool`

HasOrchestrationFailedReason returns a boolean if a field has been set.

### GetResultParams

`func (o *DescribeServicesOrchestrationResult) GetResultParams() interface{}`

GetResultParams returns the ResultParams field if non-nil, zero value otherwise.

### GetResultParamsOk

`func (o *DescribeServicesOrchestrationResult) GetResultParamsOk() (*interface{}, bool)`

GetResultParamsOk returns a tuple with the ResultParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultParams

`func (o *DescribeServicesOrchestrationResult) SetResultParams(v interface{})`

SetResultParams sets ResultParams field to given value.

### HasResultParams

`func (o *DescribeServicesOrchestrationResult) HasResultParams() bool`

HasResultParams returns a boolean if a field has been set.

### SetResultParamsNil

`func (o *DescribeServicesOrchestrationResult) SetResultParamsNil(b bool)`

 SetResultParamsNil sets the value for ResultParams to be an explicit nil

### UnsetResultParams
`func (o *DescribeServicesOrchestrationResult) UnsetResultParams()`

UnsetResultParams ensures that no value is present for ResultParams, not even an explicit nil
### GetStatus

`func (o *DescribeServicesOrchestrationResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeServicesOrchestrationResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeServicesOrchestrationResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *DescribeServicesOrchestrationResult) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DescribeServicesOrchestrationResult) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DescribeServicesOrchestrationResult) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


