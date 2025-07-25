# FleetDescribeHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amenities** | Pointer to [**[]Amenity**](Amenity.md) | The amenities available in the host cluster | [optional] 
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**AzureSubscriptionID** | Pointer to **string** | The Azure subscription ID | [optional] 
**CloudProvider** | **string** | Name of the Infra Provider | 
**DashboardEndpoint** | Pointer to **string** | The endpoint to access the dashboard | [optional] 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**HasPendingChanges** | Pointer to **bool** | Whether the host cluster has pending changes | [optional] 
**Id** | **string** | ID of a Host Cluster | 
**PendingAmenities** | Pointer to [**[]Amenity**](Amenity.md) | The pending amenities for the host cluster | [optional] 
**Region** | **string** | The region of the host cluster | 
**Status** | **string** | The status of an operation | 
**Type** | **string** |  | 

## Methods

### NewFleetDescribeHostClusterResult

`func NewFleetDescribeHostClusterResult(cloudProvider string, id string, region string, status string, type_ string, ) *FleetDescribeHostClusterResult`

NewFleetDescribeHostClusterResult instantiates a new FleetDescribeHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeHostClusterResultWithDefaults

`func NewFleetDescribeHostClusterResultWithDefaults() *FleetDescribeHostClusterResult`

NewFleetDescribeHostClusterResultWithDefaults instantiates a new FleetDescribeHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmenities

`func (o *FleetDescribeHostClusterResult) GetAmenities() []Amenity`

GetAmenities returns the Amenities field if non-nil, zero value otherwise.

### GetAmenitiesOk

`func (o *FleetDescribeHostClusterResult) GetAmenitiesOk() (*[]Amenity, bool)`

GetAmenitiesOk returns a tuple with the Amenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenities

`func (o *FleetDescribeHostClusterResult) SetAmenities(v []Amenity)`

SetAmenities sets Amenities field to given value.

### HasAmenities

`func (o *FleetDescribeHostClusterResult) HasAmenities() bool`

HasAmenities returns a boolean if a field has been set.

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

### GetAzureSubscriptionID

`func (o *FleetDescribeHostClusterResult) GetAzureSubscriptionID() string`

GetAzureSubscriptionID returns the AzureSubscriptionID field if non-nil, zero value otherwise.

### GetAzureSubscriptionIDOk

`func (o *FleetDescribeHostClusterResult) GetAzureSubscriptionIDOk() (*string, bool)`

GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionID

`func (o *FleetDescribeHostClusterResult) SetAzureSubscriptionID(v string)`

SetAzureSubscriptionID sets AzureSubscriptionID field to given value.

### HasAzureSubscriptionID

`func (o *FleetDescribeHostClusterResult) HasAzureSubscriptionID() bool`

HasAzureSubscriptionID returns a boolean if a field has been set.

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

### GetHasPendingChanges

`func (o *FleetDescribeHostClusterResult) GetHasPendingChanges() bool`

GetHasPendingChanges returns the HasPendingChanges field if non-nil, zero value otherwise.

### GetHasPendingChangesOk

`func (o *FleetDescribeHostClusterResult) GetHasPendingChangesOk() (*bool, bool)`

GetHasPendingChangesOk returns a tuple with the HasPendingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingChanges

`func (o *FleetDescribeHostClusterResult) SetHasPendingChanges(v bool)`

SetHasPendingChanges sets HasPendingChanges field to given value.

### HasHasPendingChanges

`func (o *FleetDescribeHostClusterResult) HasHasPendingChanges() bool`

HasHasPendingChanges returns a boolean if a field has been set.

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


### GetPendingAmenities

`func (o *FleetDescribeHostClusterResult) GetPendingAmenities() []Amenity`

GetPendingAmenities returns the PendingAmenities field if non-nil, zero value otherwise.

### GetPendingAmenitiesOk

`func (o *FleetDescribeHostClusterResult) GetPendingAmenitiesOk() (*[]Amenity, bool)`

GetPendingAmenitiesOk returns a tuple with the PendingAmenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmenities

`func (o *FleetDescribeHostClusterResult) SetPendingAmenities(v []Amenity)`

SetPendingAmenities sets PendingAmenities field to given value.

### HasPendingAmenities

`func (o *FleetDescribeHostClusterResult) HasPendingAmenities() bool`

HasPendingAmenities returns a boolean if a field has been set.

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


