# SearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Name of the Infra Provider | [optional] 
**Description** | Pointer to **string** | The description of the record. | [optional] 
**EnvironmentId** | Pointer to **string** | ID of a Service Environment | [optional] 
**EnvironmentKey** | Pointer to **string** | The environment key of the record. | [optional] 
**Id** | **string** | The ID of the record. | 
**Name** | Pointer to **string** | The name of the record. | [optional] 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**OrgName** | Pointer to **string** | The Organization Name of the record, if it&#39;s a user. | [optional] 
**RegionCode** | Pointer to **string** | The region code of the record, if it&#39;s a deployment cell. | [optional] 
**ResourceId** | Pointer to **string** | The resource ID of the record, if it&#39;s an instance. | [optional] 
**ResourceName** | Pointer to **string** | The resource name for this record, if it&#39;s a workflow. | [optional] 
**ServiceEnvironmentName** | Pointer to **string** | The service environment name of this record. | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**ServiceName** | Pointer to **string** | The service name of this record. | [optional] 
**ServicePlanName** | Pointer to **string** | The service plan name of this record, if it&#39;s a subscription. | [optional] 
**Status** | Pointer to **string** | The status of an operation | [optional] 
**StatusDescription** | Pointer to **string** | The status description of the record. | [optional] 
**TargetResourceName** | Pointer to **string** | The Target Resource Name of the record, if it&#39;s a proxy instance. | [optional] 
**UserEmail** | Pointer to **string** | The user email of this record, if it&#39;s a subscription. | [optional] 
**UserID** | Pointer to **string** | ID of a User | [optional] 
**Version** | Pointer to **string** | The version of this record. | [optional] 

## Methods

### NewSearchRecord

`func NewSearchRecord(id string, ) *SearchRecord`

NewSearchRecord instantiates a new SearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRecordWithDefaults

`func NewSearchRecordWithDefaults() *SearchRecord`

NewSearchRecordWithDefaults instantiates a new SearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *SearchRecord) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SearchRecord) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SearchRecord) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *SearchRecord) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetDescription

`func (o *SearchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SearchRecord) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SearchRecord) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SearchRecord) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SearchRecord) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *SearchRecord) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *SearchRecord) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *SearchRecord) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *SearchRecord) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetId

`func (o *SearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SearchRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *SearchRecord) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SearchRecord) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SearchRecord) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SearchRecord) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *SearchRecord) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *SearchRecord) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *SearchRecord) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *SearchRecord) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetRegionCode

`func (o *SearchRecord) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *SearchRecord) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *SearchRecord) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *SearchRecord) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetResourceId

`func (o *SearchRecord) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *SearchRecord) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *SearchRecord) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *SearchRecord) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *SearchRecord) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *SearchRecord) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *SearchRecord) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *SearchRecord) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetServiceEnvironmentName

`func (o *SearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *SearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *SearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.

### HasServiceEnvironmentName

`func (o *SearchRecord) HasServiceEnvironmentName() bool`

HasServiceEnvironmentName returns a boolean if a field has been set.

### GetServiceId

`func (o *SearchRecord) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SearchRecord) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SearchRecord) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SearchRecord) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *SearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *SearchRecord) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServicePlanName

`func (o *SearchRecord) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *SearchRecord) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *SearchRecord) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *SearchRecord) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetStatus

`func (o *SearchRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *SearchRecord) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *SearchRecord) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *SearchRecord) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *SearchRecord) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetTargetResourceName

`func (o *SearchRecord) GetTargetResourceName() string`

GetTargetResourceName returns the TargetResourceName field if non-nil, zero value otherwise.

### GetTargetResourceNameOk

`func (o *SearchRecord) GetTargetResourceNameOk() (*string, bool)`

GetTargetResourceNameOk returns a tuple with the TargetResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceName

`func (o *SearchRecord) SetTargetResourceName(v string)`

SetTargetResourceName sets TargetResourceName field to given value.

### HasTargetResourceName

`func (o *SearchRecord) HasTargetResourceName() bool`

HasTargetResourceName returns a boolean if a field has been set.

### GetUserEmail

`func (o *SearchRecord) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *SearchRecord) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *SearchRecord) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *SearchRecord) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserID

`func (o *SearchRecord) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *SearchRecord) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *SearchRecord) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *SearchRecord) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetVersion

`func (o *SearchRecord) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SearchRecord) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SearchRecord) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SearchRecord) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


