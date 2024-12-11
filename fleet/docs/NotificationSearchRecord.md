# NotificationSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The notification description. | 
**Id** | **string** | The Event ID of the notification. | 
**Name** | **string** | The notification name. | 
**Priority** | **string** | The priority of the notification. | 
**ResourceName** | **string** | The resource name for the notification. | 
**ServiceEnvironmentID** | **string** | The service environment ID of the notification. | 
**ServiceEnvironmentName** | **string** | The service environment name of the notification. | 
**ServiceEnvironmentType** | Pointer to **string** | The service environment type of the notification. | [optional] 
**ServiceID** | **string** | The service ID of the notification. | 
**ServiceName** | **string** | The service name of the notification. | 
**Time** | **string** | The event time of the notification. | 
**Type** | **string** | The notification type. | 

## Methods

### NewNotificationSearchRecord

`func NewNotificationSearchRecord(description string, id string, name string, priority string, resourceName string, serviceEnvironmentID string, serviceEnvironmentName string, serviceID string, serviceName string, time string, type_ string, ) *NotificationSearchRecord`

NewNotificationSearchRecord instantiates a new NotificationSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSearchRecordWithDefaults

`func NewNotificationSearchRecordWithDefaults() *NotificationSearchRecord`

NewNotificationSearchRecordWithDefaults instantiates a new NotificationSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NotificationSearchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationSearchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationSearchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *NotificationSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NotificationSearchRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationSearchRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationSearchRecord) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *NotificationSearchRecord) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotificationSearchRecord) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotificationSearchRecord) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetResourceName

`func (o *NotificationSearchRecord) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *NotificationSearchRecord) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *NotificationSearchRecord) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetServiceEnvironmentID

`func (o *NotificationSearchRecord) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *NotificationSearchRecord) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *NotificationSearchRecord) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.


### GetServiceEnvironmentName

`func (o *NotificationSearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *NotificationSearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *NotificationSearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *NotificationSearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *NotificationSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *NotificationSearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *NotificationSearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceID

`func (o *NotificationSearchRecord) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *NotificationSearchRecord) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *NotificationSearchRecord) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetServiceName

`func (o *NotificationSearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NotificationSearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *NotificationSearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetTime

`func (o *NotificationSearchRecord) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotificationSearchRecord) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotificationSearchRecord) SetTime(v string)`

SetTime sets Time field to given value.


### GetType

`func (o *NotificationSearchRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationSearchRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationSearchRecord) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


