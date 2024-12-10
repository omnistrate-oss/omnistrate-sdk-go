# FleetDescribeHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**CloudProvider** | **string** | Name of the Infra Provider | 
**DashboardEndpoint** | Pointer to **string** | The endpoint to access the dashboard | [optional] 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**Id** | **string** | ID of a Host Cluster | 
**InstanceID** | **string** | ID of a Resource Instance | 
**Region** | **string** | The region of the host cluster | 
**Status** | **string** | The status of an operation | 
**Type** | **string** |  | 

## Methods

### NewFleetDescribeHostClusterResult

`func NewFleetDescribeHostClusterResult(cloudProvider string, id string, instanceID string, region string, status string, type_ string, ) *FleetDescribeHostClusterResult`

NewFleetDescribeHostClusterResult instantiates a new FleetDescribeHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeHostClusterResultWithDefaults

`func NewFleetDescribeHostClusterResultWithDefaults() *FleetDescribeHostClusterResult`

NewFleetDescribeHostClusterResultWithDefaults instantiates a new FleetDescribeHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *FleetDescribeHostClusterResult) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *FleetDescribeHostClusterResult) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *FleetDescribeHostClusterResult) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *FleetDescribeHostClusterResult) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetCloudProvider

`func (o *FleetDescribeHostClusterResult) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *FleetDescribeHostClusterResult) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *FleetDescribeHostClusterResult) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetDashboardEndpoint

`func (o *FleetDescribeHostClusterResult) GetDashboardEndpoint() string`

GetDashboardEndpoint returns the DashboardEndpoint field if non-nil, zero value otherwise.

### GetDashboardEndpointOk

`func (o *FleetDescribeHostClusterResult) GetDashboardEndpointOk() (*string, bool)`

GetDashboardEndpointOk returns a tuple with the DashboardEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardEndpoint

`func (o *FleetDescribeHostClusterResult) SetDashboardEndpoint(v string)`

SetDashboardEndpoint sets DashboardEndpoint field to given value.

### HasDashboardEndpoint

`func (o *FleetDescribeHostClusterResult) HasDashboardEndpoint() bool`

HasDashboardEndpoint returns a boolean if a field has been set.

### GetGcpProjectID

`func (o *FleetDescribeHostClusterResult) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *FleetDescribeHostClusterResult) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *FleetDescribeHostClusterResult) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *FleetDescribeHostClusterResult) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetId

`func (o *FleetDescribeHostClusterResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetDescribeHostClusterResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetDescribeHostClusterResult) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceID

`func (o *FleetDescribeHostClusterResult) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *FleetDescribeHostClusterResult) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *FleetDescribeHostClusterResult) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.


### GetRegion

`func (o *FleetDescribeHostClusterResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetDescribeHostClusterResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetDescribeHostClusterResult) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetStatus

`func (o *FleetDescribeHostClusterResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetDescribeHostClusterResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetDescribeHostClusterResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *FleetDescribeHostClusterResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FleetDescribeHostClusterResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FleetDescribeHostClusterResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


