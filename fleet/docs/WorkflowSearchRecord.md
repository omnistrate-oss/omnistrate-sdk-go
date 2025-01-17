# WorkflowSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the workflow. | 
**ResourceName** | **string** | The name of the resource associated with the workflow. | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceEnvironmentName** | **string** | The service environment name of the workflow. | 
**ServiceEnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceName** | **string** | The service name of the workflow. | 
**Status** | **string** | The status of an operation | 
**Type** | **string** | The workflow type. | 

## Methods

### NewWorkflowSearchRecord

`func NewWorkflowSearchRecord(id string, resourceName string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceName string, status string, type_ string, ) *WorkflowSearchRecord`

NewWorkflowSearchRecord instantiates a new WorkflowSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSearchRecordWithDefaults

`func NewWorkflowSearchRecordWithDefaults() *WorkflowSearchRecord`

NewWorkflowSearchRecordWithDefaults instantiates a new WorkflowSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetResourceName

`func (o *WorkflowSearchRecord) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *WorkflowSearchRecord) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *WorkflowSearchRecord) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetServiceEnvironmentId

`func (o *WorkflowSearchRecord) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *WorkflowSearchRecord) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *WorkflowSearchRecord) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceEnvironmentName

`func (o *WorkflowSearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *WorkflowSearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *WorkflowSearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *WorkflowSearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *WorkflowSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *WorkflowSearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *WorkflowSearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceId

`func (o *WorkflowSearchRecord) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *WorkflowSearchRecord) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *WorkflowSearchRecord) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *WorkflowSearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *WorkflowSearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *WorkflowSearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatus

`func (o *WorkflowSearchRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowSearchRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowSearchRecord) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *WorkflowSearchRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowSearchRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowSearchRecord) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


