# ResourceInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**CloudProvider** | **string** | Name of the Infra Provider | 
**ConsumptionResourceInstanceResult** | [**DescribeResourceInstanceResult**](DescribeResourceInstanceResult.md) |  | 
**DefaultSubscription** | **bool** | Whether the subscription is the default subscription | 
**DeploymentCellID** | Pointer to **string** | ID of a Host Cluster | [optional] 
**EnvironmentId** | **string** | ID of a Service Environment | 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**InputParams** | **interface{}** | Custom input parameters | 
**InstanceDebugCommands** | **[]string** | The debug commands to access the instance | 
**IntegrationsStatus** | [**[]IntegrationStatus**](IntegrationStatus.md) | List of individual integrations and their statuses for the instance | 
**MaintenanceTasks** | Pointer to **map[string]interface{}** | Pending actions or maintenance tasks for the resource instance, with action type as key and reference key as value. | [optional] 
**ManagedResourceType** | Pointer to **string** | The managed resource type of instance | [optional] 
**ManualOverride** | Pointer to [**ManualOverride**](ManualOverride.md) |  | [optional] 
**OrganizationId** | **string** | ID of an Org | 
**OrganizationName** | **string** | The organization name of the resource instance. | 
**PortsRegistrationStatus** | Pointer to **map[string][]int64** | The ports registration status of ports based proxy instance | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The product tier name | 
**ProductTierType** | **string** | The product tier type | 
**ProxyType** | Pointer to **string** | The proxy type of instance | [optional] 
**ResourceVersionSummaries** | [**[]ResourceVersionSummary**](ResourceVersionSummary.md) | Associated internal/external resources deployed for the instance, the corresponding versions deployed. | 
**ServiceEnvName** | **string** | The service environment name | 
**ServiceId** | **string** | ID of a Service | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 
**ServiceModelId** | **string** | ID of a Service Model | 
**ServiceModelName** | **string** | The service model name | 
**ServiceModelType** | **string** | The service model type | 
**ServiceName** | **string** | The service name | 
**SubscriptionId** | **string** | ID of a Subscription | 
**SubscriptionOwnerName** | **string** | The subscription owner name | 
**SubscriptionStatus** | Pointer to **string** | Subscription Status | [optional] 
**TierVersion** | **string** | The tier version of the resource instance. | 
**TierVersionReleasedAt** | **string** | The timestamp when the version set was released. | 
**TierVersionReleasedByUserId** | **string** | ID of a User | 
**TierVersionReleasedByUserName** | **string** | The name of the user who released the version set. | 
**TierVersionStatus** | **string** | The tier version set status. | 

## Methods

### NewResourceInstance

`func NewResourceInstance(cloudProvider string, consumptionResourceInstanceResult DescribeResourceInstanceResult, defaultSubscription bool, environmentId string, inputParams interface{}, instanceDebugCommands []string, integrationsStatus []IntegrationStatus, organizationId string, organizationName string, productTierId string, productTierName string, productTierType string, resourceVersionSummaries []ResourceVersionSummary, serviceEnvName string, serviceId string, serviceModelId string, serviceModelName string, serviceModelType string, serviceName string, subscriptionId string, subscriptionOwnerName string, tierVersion string, tierVersionReleasedAt string, tierVersionReleasedByUserId string, tierVersionReleasedByUserName string, tierVersionStatus string, ) *ResourceInstance`

NewResourceInstance instantiates a new ResourceInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInstanceWithDefaults

`func NewResourceInstanceWithDefaults() *ResourceInstance`

NewResourceInstanceWithDefaults instantiates a new ResourceInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *ResourceInstance) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *ResourceInstance) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *ResourceInstance) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *ResourceInstance) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ResourceInstance) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ResourceInstance) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ResourceInstance) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetConsumptionResourceInstanceResult

`func (o *ResourceInstance) GetConsumptionResourceInstanceResult() DescribeResourceInstanceResult`

GetConsumptionResourceInstanceResult returns the ConsumptionResourceInstanceResult field if non-nil, zero value otherwise.

### GetConsumptionResourceInstanceResultOk

`func (o *ResourceInstance) GetConsumptionResourceInstanceResultOk() (*DescribeResourceInstanceResult, bool)`

GetConsumptionResourceInstanceResultOk returns a tuple with the ConsumptionResourceInstanceResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionResourceInstanceResult

`func (o *ResourceInstance) SetConsumptionResourceInstanceResult(v DescribeResourceInstanceResult)`

SetConsumptionResourceInstanceResult sets ConsumptionResourceInstanceResult field to given value.


### GetDefaultSubscription

`func (o *ResourceInstance) GetDefaultSubscription() bool`

GetDefaultSubscription returns the DefaultSubscription field if non-nil, zero value otherwise.

### GetDefaultSubscriptionOk

`func (o *ResourceInstance) GetDefaultSubscriptionOk() (*bool, bool)`

GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubscription

`func (o *ResourceInstance) SetDefaultSubscription(v bool)`

SetDefaultSubscription sets DefaultSubscription field to given value.


### GetDeploymentCellID

`func (o *ResourceInstance) GetDeploymentCellID() string`

GetDeploymentCellID returns the DeploymentCellID field if non-nil, zero value otherwise.

### GetDeploymentCellIDOk

`func (o *ResourceInstance) GetDeploymentCellIDOk() (*string, bool)`

GetDeploymentCellIDOk returns a tuple with the DeploymentCellID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellID

`func (o *ResourceInstance) SetDeploymentCellID(v string)`

SetDeploymentCellID sets DeploymentCellID field to given value.

### HasDeploymentCellID

`func (o *ResourceInstance) HasDeploymentCellID() bool`

HasDeploymentCellID returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ResourceInstance) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ResourceInstance) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ResourceInstance) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetGcpProjectID

`func (o *ResourceInstance) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *ResourceInstance) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *ResourceInstance) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *ResourceInstance) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetInputParams

`func (o *ResourceInstance) GetInputParams() interface{}`

GetInputParams returns the InputParams field if non-nil, zero value otherwise.

### GetInputParamsOk

`func (o *ResourceInstance) GetInputParamsOk() (*interface{}, bool)`

GetInputParamsOk returns a tuple with the InputParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParams

`func (o *ResourceInstance) SetInputParams(v interface{})`

SetInputParams sets InputParams field to given value.


### SetInputParamsNil

`func (o *ResourceInstance) SetInputParamsNil(b bool)`

 SetInputParamsNil sets the value for InputParams to be an explicit nil

### UnsetInputParams
`func (o *ResourceInstance) UnsetInputParams()`

UnsetInputParams ensures that no value is present for InputParams, not even an explicit nil
### GetInstanceDebugCommands

`func (o *ResourceInstance) GetInstanceDebugCommands() []string`

GetInstanceDebugCommands returns the InstanceDebugCommands field if non-nil, zero value otherwise.

### GetInstanceDebugCommandsOk

`func (o *ResourceInstance) GetInstanceDebugCommandsOk() (*[]string, bool)`

GetInstanceDebugCommandsOk returns a tuple with the InstanceDebugCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDebugCommands

`func (o *ResourceInstance) SetInstanceDebugCommands(v []string)`

SetInstanceDebugCommands sets InstanceDebugCommands field to given value.


### GetIntegrationsStatus

`func (o *ResourceInstance) GetIntegrationsStatus() []IntegrationStatus`

GetIntegrationsStatus returns the IntegrationsStatus field if non-nil, zero value otherwise.

### GetIntegrationsStatusOk

`func (o *ResourceInstance) GetIntegrationsStatusOk() (*[]IntegrationStatus, bool)`

GetIntegrationsStatusOk returns a tuple with the IntegrationsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsStatus

`func (o *ResourceInstance) SetIntegrationsStatus(v []IntegrationStatus)`

SetIntegrationsStatus sets IntegrationsStatus field to given value.


### GetMaintenanceTasks

`func (o *ResourceInstance) GetMaintenanceTasks() map[string]interface{}`

GetMaintenanceTasks returns the MaintenanceTasks field if non-nil, zero value otherwise.

### GetMaintenanceTasksOk

`func (o *ResourceInstance) GetMaintenanceTasksOk() (*map[string]interface{}, bool)`

GetMaintenanceTasksOk returns a tuple with the MaintenanceTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceTasks

`func (o *ResourceInstance) SetMaintenanceTasks(v map[string]interface{})`

SetMaintenanceTasks sets MaintenanceTasks field to given value.

### HasMaintenanceTasks

`func (o *ResourceInstance) HasMaintenanceTasks() bool`

HasMaintenanceTasks returns a boolean if a field has been set.

### GetManagedResourceType

`func (o *ResourceInstance) GetManagedResourceType() string`

GetManagedResourceType returns the ManagedResourceType field if non-nil, zero value otherwise.

### GetManagedResourceTypeOk

`func (o *ResourceInstance) GetManagedResourceTypeOk() (*string, bool)`

GetManagedResourceTypeOk returns a tuple with the ManagedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedResourceType

`func (o *ResourceInstance) SetManagedResourceType(v string)`

SetManagedResourceType sets ManagedResourceType field to given value.

### HasManagedResourceType

`func (o *ResourceInstance) HasManagedResourceType() bool`

HasManagedResourceType returns a boolean if a field has been set.

### GetManualOverride

`func (o *ResourceInstance) GetManualOverride() ManualOverride`

GetManualOverride returns the ManualOverride field if non-nil, zero value otherwise.

### GetManualOverrideOk

`func (o *ResourceInstance) GetManualOverrideOk() (*ManualOverride, bool)`

GetManualOverrideOk returns a tuple with the ManualOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverride

`func (o *ResourceInstance) SetManualOverride(v ManualOverride)`

SetManualOverride sets ManualOverride field to given value.

### HasManualOverride

`func (o *ResourceInstance) HasManualOverride() bool`

HasManualOverride returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ResourceInstance) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ResourceInstance) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ResourceInstance) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *ResourceInstance) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ResourceInstance) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ResourceInstance) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetPortsRegistrationStatus

`func (o *ResourceInstance) GetPortsRegistrationStatus() map[string][]int64`

GetPortsRegistrationStatus returns the PortsRegistrationStatus field if non-nil, zero value otherwise.

### GetPortsRegistrationStatusOk

`func (o *ResourceInstance) GetPortsRegistrationStatusOk() (*map[string][]int64, bool)`

GetPortsRegistrationStatusOk returns a tuple with the PortsRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsRegistrationStatus

`func (o *ResourceInstance) SetPortsRegistrationStatus(v map[string][]int64)`

SetPortsRegistrationStatus sets PortsRegistrationStatus field to given value.

### HasPortsRegistrationStatus

`func (o *ResourceInstance) HasPortsRegistrationStatus() bool`

HasPortsRegistrationStatus returns a boolean if a field has been set.

### GetProductTierId

`func (o *ResourceInstance) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ResourceInstance) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ResourceInstance) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierName

`func (o *ResourceInstance) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *ResourceInstance) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *ResourceInstance) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetProductTierType

`func (o *ResourceInstance) GetProductTierType() string`

GetProductTierType returns the ProductTierType field if non-nil, zero value otherwise.

### GetProductTierTypeOk

`func (o *ResourceInstance) GetProductTierTypeOk() (*string, bool)`

GetProductTierTypeOk returns a tuple with the ProductTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierType

`func (o *ResourceInstance) SetProductTierType(v string)`

SetProductTierType sets ProductTierType field to given value.


### GetProxyType

`func (o *ResourceInstance) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *ResourceInstance) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *ResourceInstance) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *ResourceInstance) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetResourceVersionSummaries

`func (o *ResourceInstance) GetResourceVersionSummaries() []ResourceVersionSummary`

GetResourceVersionSummaries returns the ResourceVersionSummaries field if non-nil, zero value otherwise.

### GetResourceVersionSummariesOk

`func (o *ResourceInstance) GetResourceVersionSummariesOk() (*[]ResourceVersionSummary, bool)`

GetResourceVersionSummariesOk returns a tuple with the ResourceVersionSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersionSummaries

`func (o *ResourceInstance) SetResourceVersionSummaries(v []ResourceVersionSummary)`

SetResourceVersionSummaries sets ResourceVersionSummaries field to given value.


### GetServiceEnvName

`func (o *ResourceInstance) GetServiceEnvName() string`

GetServiceEnvName returns the ServiceEnvName field if non-nil, zero value otherwise.

### GetServiceEnvNameOk

`func (o *ResourceInstance) GetServiceEnvNameOk() (*string, bool)`

GetServiceEnvNameOk returns a tuple with the ServiceEnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvName

`func (o *ResourceInstance) SetServiceEnvName(v string)`

SetServiceEnvName sets ServiceEnvName field to given value.


### GetServiceId

`func (o *ResourceInstance) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ResourceInstance) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ResourceInstance) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceLogoURL

`func (o *ResourceInstance) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *ResourceInstance) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *ResourceInstance) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *ResourceInstance) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetServiceModelId

`func (o *ResourceInstance) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *ResourceInstance) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *ResourceInstance) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetServiceModelName

`func (o *ResourceInstance) GetServiceModelName() string`

GetServiceModelName returns the ServiceModelName field if non-nil, zero value otherwise.

### GetServiceModelNameOk

`func (o *ResourceInstance) GetServiceModelNameOk() (*string, bool)`

GetServiceModelNameOk returns a tuple with the ServiceModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelName

`func (o *ResourceInstance) SetServiceModelName(v string)`

SetServiceModelName sets ServiceModelName field to given value.


### GetServiceModelType

`func (o *ResourceInstance) GetServiceModelType() string`

GetServiceModelType returns the ServiceModelType field if non-nil, zero value otherwise.

### GetServiceModelTypeOk

`func (o *ResourceInstance) GetServiceModelTypeOk() (*string, bool)`

GetServiceModelTypeOk returns a tuple with the ServiceModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelType

`func (o *ResourceInstance) SetServiceModelType(v string)`

SetServiceModelType sets ServiceModelType field to given value.


### GetServiceName

`func (o *ResourceInstance) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourceInstance) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourceInstance) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetSubscriptionId

`func (o *ResourceInstance) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ResourceInstance) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ResourceInstance) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetSubscriptionOwnerName

`func (o *ResourceInstance) GetSubscriptionOwnerName() string`

GetSubscriptionOwnerName returns the SubscriptionOwnerName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerNameOk

`func (o *ResourceInstance) GetSubscriptionOwnerNameOk() (*string, bool)`

GetSubscriptionOwnerNameOk returns a tuple with the SubscriptionOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerName

`func (o *ResourceInstance) SetSubscriptionOwnerName(v string)`

SetSubscriptionOwnerName sets SubscriptionOwnerName field to given value.


### GetSubscriptionStatus

`func (o *ResourceInstance) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *ResourceInstance) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *ResourceInstance) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *ResourceInstance) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetTierVersion

`func (o *ResourceInstance) GetTierVersion() string`

GetTierVersion returns the TierVersion field if non-nil, zero value otherwise.

### GetTierVersionOk

`func (o *ResourceInstance) GetTierVersionOk() (*string, bool)`

GetTierVersionOk returns a tuple with the TierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierVersion

`func (o *ResourceInstance) SetTierVersion(v string)`

SetTierVersion sets TierVersion field to given value.


### GetTierVersionReleasedAt

`func (o *ResourceInstance) GetTierVersionReleasedAt() string`

GetTierVersionReleasedAt returns the TierVersionReleasedAt field if non-nil, zero value otherwise.

### GetTierVersionReleasedAtOk

`func (o *ResourceInstance) GetTierVersionReleasedAtOk() (*string, bool)`

GetTierVersionReleasedAtOk returns a tuple with the TierVersionReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierVersionReleasedAt

`func (o *ResourceInstance) SetTierVersionReleasedAt(v string)`

SetTierVersionReleasedAt sets TierVersionReleasedAt field to given value.


### GetTierVersionReleasedByUserId

`func (o *ResourceInstance) GetTierVersionReleasedByUserId() string`

GetTierVersionReleasedByUserId returns the TierVersionReleasedByUserId field if non-nil, zero value otherwise.

### GetTierVersionReleasedByUserIdOk

`func (o *ResourceInstance) GetTierVersionReleasedByUserIdOk() (*string, bool)`

GetTierVersionReleasedByUserIdOk returns a tuple with the TierVersionReleasedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierVersionReleasedByUserId

`func (o *ResourceInstance) SetTierVersionReleasedByUserId(v string)`

SetTierVersionReleasedByUserId sets TierVersionReleasedByUserId field to given value.


### GetTierVersionReleasedByUserName

`func (o *ResourceInstance) GetTierVersionReleasedByUserName() string`

GetTierVersionReleasedByUserName returns the TierVersionReleasedByUserName field if non-nil, zero value otherwise.

### GetTierVersionReleasedByUserNameOk

`func (o *ResourceInstance) GetTierVersionReleasedByUserNameOk() (*string, bool)`

GetTierVersionReleasedByUserNameOk returns a tuple with the TierVersionReleasedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierVersionReleasedByUserName

`func (o *ResourceInstance) SetTierVersionReleasedByUserName(v string)`

SetTierVersionReleasedByUserName sets TierVersionReleasedByUserName field to given value.


### GetTierVersionStatus

`func (o *ResourceInstance) GetTierVersionStatus() string`

GetTierVersionStatus returns the TierVersionStatus field if non-nil, zero value otherwise.

### GetTierVersionStatusOk

`func (o *ResourceInstance) GetTierVersionStatusOk() (*string, bool)`

GetTierVersionStatusOk returns a tuple with the TierVersionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierVersionStatus

`func (o *ResourceInstance) SetTierVersionStatus(v string)`

SetTierVersionStatus sets TierVersionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


