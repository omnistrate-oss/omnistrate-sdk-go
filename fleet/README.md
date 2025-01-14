# Go API client for fleet

REST API for Omnistrate Fleet

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2022-09-01-00
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import fleet "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `fleet.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), fleet.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `fleet.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), fleet.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `fleet.ContextOperationServerIndices` and `fleet.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), fleet.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), fleet.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.omnistrate.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditEventsApiAPI* | [**AuditEventsApiAuditEvents**](docs/AuditEventsApiAPI.md#auditeventsapiauditevents) | **Get** /2022-09-01-00/fleet/audit-events | AuditEvents audit-events-api
*CustomerOnboardingsApiAPI* | [**CustomerOnboardingsApiCreateCustomerOnboarding**](docs/CustomerOnboardingsApiAPI.md#customeronboardingsapicreatecustomeronboarding) | **Post** /2022-09-01-00/fleet/customer-onboarding | CreateCustomerOnboarding customer-onboardings-api
*CustomerOnboardingsApiAPI* | [**CustomerOnboardingsApiDeleteCustomerOnboarding**](docs/CustomerOnboardingsApiAPI.md#customeronboardingsapideletecustomeronboarding) | **Delete** /2022-09-01-00/fleet/customer-onboarding/{id} | DeleteCustomerOnboarding customer-onboardings-api
*CustomerOnboardingsApiAPI* | [**CustomerOnboardingsApiDescribeCustomerOnboarding**](docs/CustomerOnboardingsApiAPI.md#customeronboardingsapidescribecustomeronboarding) | **Get** /2022-09-01-00/fleet/customer-onboarding/{id} | DescribeCustomerOnboarding customer-onboardings-api
*CustomerOnboardingsApiAPI* | [**CustomerOnboardingsApiListCustomerOnboardingStages**](docs/CustomerOnboardingsApiAPI.md#customeronboardingsapilistcustomeronboardingstages) | **Get** /2022-09-01-00/fleet/customer-onboarding-stages | ListCustomerOnboardingStages customer-onboardings-api
*CustomerOnboardingsApiAPI* | [**CustomerOnboardingsApiListCustomerOnboardings**](docs/CustomerOnboardingsApiAPI.md#customeronboardingsapilistcustomeronboardings) | **Get** /2022-09-01-00/fleet/customer-onboarding | ListCustomerOnboardings customer-onboardings-api
*CustomerOnboardingsApiAPI* | [**CustomerOnboardingsApiUpdateCustomerOnboarding**](docs/CustomerOnboardingsApiAPI.md#customeronboardingsapiupdatecustomeronboarding) | **Patch** /2022-09-01-00/fleet/customer-onboarding/{id} | UpdateCustomerOnboarding customer-onboardings-api
*EventsApiAPI* | [**EventsApiAcknowledgeEvent**](docs/EventsApiAPI.md#eventsapiacknowledgeevent) | **Delete** /2022-09-01-00/fleet/events/{id} | AcknowledgeEvent events-api
*EventsApiAPI* | [**EventsApiListEvents**](docs/EventsApiAPI.md#eventsapilistevents) | **Get** /2022-09-01-00/fleet/events | ListEvents events-api
*FleetCustomNetworkApiAPI* | [**FleetCustomNetworkApiCreateCustomNetwork**](docs/FleetCustomNetworkApiAPI.md#fleetcustomnetworkapicreatecustomnetwork) | **Post** /2022-09-01-00/fleet/custom-network | CreateCustomNetwork fleet-custom-network-api
*FleetCustomNetworkApiAPI* | [**FleetCustomNetworkApiDeleteCustomNetwork**](docs/FleetCustomNetworkApiAPI.md#fleetcustomnetworkapideletecustomnetwork) | **Delete** /2022-09-01-00/fleet/custom-network/{id} | DeleteCustomNetwork fleet-custom-network-api
*FleetCustomNetworkApiAPI* | [**FleetCustomNetworkApiDescribeCustomNetwork**](docs/FleetCustomNetworkApiAPI.md#fleetcustomnetworkapidescribecustomnetwork) | **Get** /2022-09-01-00/fleet/custom-network/{id} | DescribeCustomNetwork fleet-custom-network-api
*FleetCustomNetworkApiAPI* | [**FleetCustomNetworkApiListCustomNetworks**](docs/FleetCustomNetworkApiAPI.md#fleetcustomnetworkapilistcustomnetworks) | **Get** /2022-09-01-00/fleet/custom-network | ListCustomNetworks fleet-custom-network-api
*FleetCustomNetworkApiAPI* | [**FleetCustomNetworkApiUpdateCustomNetwork**](docs/FleetCustomNetworkApiAPI.md#fleetcustomnetworkapiupdatecustomnetwork) | **Patch** /2022-09-01-00/fleet/custom-network/{id} | UpdateCustomNetwork fleet-custom-network-api
*FleetFeaturesApiAPI* | [**FleetFeaturesApiDescribeFleetFeature**](docs/FleetFeaturesApiAPI.md#fleetfeaturesapidescribefleetfeature) | **Get** /2022-09-01-00/fleet/feature/{feature} | DescribeFleetFeature fleet-features-api
*FleetFeaturesApiAPI* | [**FleetFeaturesApiDisableFleetFeature**](docs/FleetFeaturesApiAPI.md#fleetfeaturesapidisablefleetfeature) | **Delete** /2022-09-01-00/fleet/feature | DisableFleetFeature fleet-features-api
*FleetFeaturesApiAPI* | [**FleetFeaturesApiEnableFleetFeature**](docs/FleetFeaturesApiAPI.md#fleetfeaturesapienablefleetfeature) | **Put** /2022-09-01-00/fleet/feature | EnableFleetFeature fleet-features-api
*FleetFeaturesApiAPI* | [**FleetFeaturesApiListFleetFeatures**](docs/FleetFeaturesApiAPI.md#fleetfeaturesapilistfleetfeatures) | **Get** /2022-09-01-00/fleet/features | ListFleetFeatures fleet-features-api
*FleetWorkflowsApiAPI* | [**FleetWorkflowsApiDescribeServiceWorkflow**](docs/FleetWorkflowsApiAPI.md#fleetworkflowsapidescribeserviceworkflow) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows/{id} | DescribeServiceWorkflow fleet-workflows-api
*FleetWorkflowsApiAPI* | [**FleetWorkflowsApiDescribeServiceWorkflowSummary**](docs/FleetWorkflowsApiAPI.md#fleetworkflowsapidescribeserviceworkflowsummary) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows-summary | DescribeServiceWorkflowSummary fleet-workflows-api
*FleetWorkflowsApiAPI* | [**FleetWorkflowsApiGetWorkflowEvents**](docs/FleetWorkflowsApiAPI.md#fleetworkflowsapigetworkflowevents) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows/{id}/events | GetWorkflowEvents fleet-workflows-api
*FleetWorkflowsApiAPI* | [**FleetWorkflowsApiListServiceWorkflows**](docs/FleetWorkflowsApiAPI.md#fleetworkflowsapilistserviceworkflows) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows | ListServiceWorkflows fleet-workflows-api
*FleetWorkflowsApiAPI* | [**FleetWorkflowsApiTerminateServiceWorkflow**](docs/FleetWorkflowsApiAPI.md#fleetworkflowsapiterminateserviceworkflow) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows/{id} | TerminateServiceWorkflow fleet-workflows-api
*FleetWorkflowsApiAPI* | [**FleetWorkflowsApiUpdateServiceWorkflow**](docs/FleetWorkflowsApiAPI.md#fleetworkflowsapiupdateserviceworkflow) | **Patch** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-workflows/{id} | UpdateServiceWorkflow fleet-workflows-api
*HelmPackageApiAPI* | [**HelmPackageApiListHelmPackageInstallations**](docs/HelmPackageApiAPI.md#helmpackageapilisthelmpackageinstallations) | **Get** /2022-09-01-00/fleet/helm-package-installations | ListHelmPackageInstallations helm-package-api
*HostclusterApiAPI* | [**HostclusterApiListHostClusters**](docs/HostclusterApiAPI.md#hostclusterapilisthostclusters) | **Get** /2022-09-01-00/fleet/host-clusters | ListHostClusters hostcluster-api
*InventoryApiAPI* | [**InventoryApiAddCapacityToResourceInstance**](docs/InventoryApiAPI.md#inventoryapiaddcapacitytoresourceinstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/add-capacity | AddCapacityToResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiAddCustomDNSToResourceInstance**](docs/InventoryApiAPI.md#inventoryapiaddcustomdnstoresourceinstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/{resourceKey}/instance/{instanceId}/custom-dns | AddCustomDNSToResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiApproveSubscriptionRequest**](docs/InventoryApiAPI.md#inventoryapiapprovesubscriptionrequest) | **Put** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/request/{id} | ApproveSubscriptionRequest inventory-api
*InventoryApiAPI* | [**InventoryApiCancelUpgradePath**](docs/InventoryApiAPI.md#inventoryapicancelupgradepath) | **Post** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-path/{upgradePathId}/cancel | CancelUpgradePath inventory-api
*InventoryApiAPI* | [**InventoryApiCreateConsumptionUser**](docs/InventoryApiAPI.md#inventoryapicreateconsumptionuser) | **Post** /2022-09-01-00/fleet/user | CreateConsumptionUser inventory-api
*InventoryApiAPI* | [**InventoryApiCreateProxyResourceInstance**](docs/InventoryApiAPI.md#inventoryapicreateproxyresourceinstance) | **Post** /2022-09-01-00/fleet/proxy-resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{sourceResourceKey} | CreateProxyResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiCreateResourceInstance**](docs/InventoryApiAPI.md#inventoryapicreateresourceinstance) | **Post** /2022-09-01-00/fleet/resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{resourceKey} | CreateResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiCreateResourceInstanceSnapshot**](docs/InventoryApiAPI.md#inventoryapicreateresourceinstancesnapshot) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/snapshot | CreateResourceInstanceSnapshot inventory-api
*InventoryApiAPI* | [**InventoryApiCreateServicesOrchestration**](docs/InventoryApiAPI.md#inventoryapicreateservicesorchestration) | **Post** /2022-09-01-00/fleet/services-orchestration | CreateServicesOrchestration inventory-api
*InventoryApiAPI* | [**InventoryApiCreateUpgradePath**](docs/InventoryApiAPI.md#inventoryapicreateupgradepath) | **Post** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-path | CreateUpgradePath inventory-api
*InventoryApiAPI* | [**InventoryApiDeleteProxyResourceInstance**](docs/InventoryApiAPI.md#inventoryapideleteproxyresourceinstance) | **Delete** /2022-09-01-00/fleet/proxy-resource-instance/{serviceProviderId}/{serviceKey}/{serviceAPIVersion}/{serviceEnvironmentKey}/{serviceModelKey}/{productTierKey}/{id} | DeleteProxyResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiDeleteResourceInstance**](docs/InventoryApiAPI.md#inventoryapideleteresourceinstance) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId} | DeleteResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiDeleteResourceInstanceSnapshot**](docs/InventoryApiAPI.md#inventoryapideleteresourceinstancesnapshot) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/snapshot/{snapshotId} | DeleteResourceInstanceSnapshot inventory-api
*InventoryApiAPI* | [**InventoryApiDeleteServicesOrchestration**](docs/InventoryApiAPI.md#inventoryapideleteservicesorchestration) | **Delete** /2022-09-01-00/fleet/services-orchestration/{id} | DeleteServicesOrchestration inventory-api
*InventoryApiAPI* | [**InventoryApiDeleteUser**](docs/InventoryApiAPI.md#inventoryapideleteuser) | **Delete** /2022-09-01-00/fleet/user/{userId} | DeleteUser inventory-api
*InventoryApiAPI* | [**InventoryApiDenySubscriptionRequest**](docs/InventoryApiAPI.md#inventoryapidenysubscriptionrequest) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/request/{id} | DenySubscriptionRequest inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeHostCluster**](docs/InventoryApiAPI.md#inventoryapidescribehostcluster) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/host-cluster/{id} | DescribeHostCluster inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeInstanceEvent**](docs/InventoryApiAPI.md#inventoryapidescribeinstanceevent) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/event/{id} | DescribeInstanceEvent inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeInventorySummary**](docs/InventoryApiAPI.md#inventoryapidescribeinventorysummary) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/service-inventory-summary | DescribeInventorySummary inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeOrgUser**](docs/InventoryApiAPI.md#inventoryapidescribeorguser) | **Get** /2022-09-01-00/fleet/user/{userId} | DescribeOrgUser inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeOrganization**](docs/InventoryApiAPI.md#inventoryapidescribeorganization) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/organization/{organizationId} | DescribeOrganization inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeResource**](docs/InventoryApiAPI.md#inventoryapidescriberesource) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/model/{serviceModelId}/productTier/{productTierId}/resource/{resourceId} | DescribeResource inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeResourceInstance**](docs/InventoryApiAPI.md#inventoryapidescriberesourceinstance) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId} | DescribeResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeResourceInstanceSnapshotFromTime**](docs/InventoryApiAPI.md#inventoryapidescriberesourceinstancesnapshotfromtime) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/restore | DescribeResourceInstanceSnapshotFromTime inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeServiceOffering**](docs/InventoryApiAPI.md#inventoryapidescribeserviceoffering) | **Get** /2022-09-01-00/fleet/service-offering/{serviceId} | DescribeServiceOffering inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeServiceOfferingResource**](docs/InventoryApiAPI.md#inventoryapidescribeserviceofferingresource) | **Get** /2022-09-01-00/fleet/service-offering/{serviceId}/resource/{resourceId}/instance/{instanceId} | DescribeServiceOfferingResource inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeServicesOrchestration**](docs/InventoryApiAPI.md#inventoryapidescribeservicesorchestration) | **Get** /2022-09-01-00/fleet/services-orchestration/{id} | DescribeServicesOrchestration inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeSubscription**](docs/InventoryApiAPI.md#inventoryapidescribesubscription) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/{id} | DescribeSubscription inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeSubscriptionRequest**](docs/InventoryApiAPI.md#inventoryapidescribesubscriptionrequest) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/request/{id} | DescribeSubscriptionRequest inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeUpgradePath**](docs/InventoryApiAPI.md#inventoryapidescribeupgradepath) | **Get** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-path/{upgradePathId} | DescribeUpgradePath inventory-api
*InventoryApiAPI* | [**InventoryApiDescribeUser**](docs/InventoryApiAPI.md#inventoryapidescribeuser) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/user/{userId} | DescribeUser inventory-api
*InventoryApiAPI* | [**InventoryApiFailoverResourceInstance**](docs/InventoryApiAPI.md#inventoryapifailoverresourceinstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/failover | FailoverResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiGenerateTokenForHostClusterDashboard**](docs/InventoryApiAPI.md#inventoryapigeneratetokenforhostclusterdashboard) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/host-cluster/{id}/dashboard/token | GenerateTokenForHostClusterDashboard inventory-api
*InventoryApiAPI* | [**InventoryApiListActiveOrganizations**](docs/InventoryApiAPI.md#inventoryapilistactiveorganizations) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/organizations | ListActiveOrganizations inventory-api
*InventoryApiAPI* | [**InventoryApiListAllUsers**](docs/InventoryApiAPI.md#inventoryapilistallusers) | **Get** /2022-09-01-00/fleet/users | ListAllUsers inventory-api
*InventoryApiAPI* | [**InventoryApiListDependentComponents**](docs/InventoryApiAPI.md#inventoryapilistdependentcomponents) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/dependent-components | ListDependentComponents inventory-api
*InventoryApiAPI* | [**InventoryApiListEligibleInstancesPerUpgrade**](docs/InventoryApiAPI.md#inventoryapilisteligibleinstancesperupgrade) | **Get** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-path/{upgradePathId}/eligible-instances | ListEligibleInstancesPerUpgrade inventory-api
*InventoryApiAPI* | [**InventoryApiListHostClusters**](docs/InventoryApiAPI.md#inventoryapilisthostclusters) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/host-clusters | ListHostClusters inventory-api
*InventoryApiAPI* | [**InventoryApiListInstanceEvents**](docs/InventoryApiAPI.md#inventoryapilistinstanceevents) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/events | ListInstanceEvents inventory-api
*InventoryApiAPI* | [**InventoryApiListLinkedInstances**](docs/InventoryApiAPI.md#inventoryapilistlinkedinstances) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/linked-instances | ListLinkedInstances inventory-api
*InventoryApiAPI* | [**InventoryApiListResourceInstanceSnapshots**](docs/InventoryApiAPI.md#inventoryapilistresourceinstancesnapshots) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/snapshot | ListResourceInstanceSnapshots inventory-api
*InventoryApiAPI* | [**InventoryApiListResourceInstances**](docs/InventoryApiAPI.md#inventoryapilistresourceinstances) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instances/ | ListResourceInstances inventory-api
*InventoryApiAPI* | [**InventoryApiListResources**](docs/InventoryApiAPI.md#inventoryapilistresources) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/model/{serviceModelId}/productTier/{productTierId}/resources | ListResources inventory-api
*InventoryApiAPI* | [**InventoryApiListServiceOfferings**](docs/InventoryApiAPI.md#inventoryapilistserviceofferings) | **Get** /2022-09-01-00/fleet/service-offering | ListServiceOfferings inventory-api
*InventoryApiAPI* | [**InventoryApiListServicesOrchestrations**](docs/InventoryApiAPI.md#inventoryapilistservicesorchestrations) | **Get** /2022-09-01-00/fleet/services-orchestration | ListServicesOrchestrations inventory-api
*InventoryApiAPI* | [**InventoryApiListSubscription**](docs/InventoryApiAPI.md#inventoryapilistsubscription) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription | ListSubscription inventory-api
*InventoryApiAPI* | [**InventoryApiListSubscriptionRequests**](docs/InventoryApiAPI.md#inventoryapilistsubscriptionrequests) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/request | ListSubscriptionRequests inventory-api
*InventoryApiAPI* | [**InventoryApiListUpgradePaths**](docs/InventoryApiAPI.md#inventoryapilistupgradepaths) | **Get** /2022-09-01-00/fleet/service/{serviceId}/productTier/{productTierId}/upgrade-paths | ListUpgradePaths inventory-api
*InventoryApiAPI* | [**InventoryApiListUsers**](docs/InventoryApiAPI.md#inventoryapilistusers) | **Get** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/users | ListUsers inventory-api
*InventoryApiAPI* | [**InventoryApiModifyServicesOrchestration**](docs/InventoryApiAPI.md#inventoryapimodifyservicesorchestration) | **Patch** /2022-09-01-00/fleet/services-orchestration/{id} | ModifyServicesOrchestration inventory-api
*InventoryApiAPI* | [**InventoryApiRemoveCapacityFromResourceInstance**](docs/InventoryApiAPI.md#inventoryapiremovecapacityfromresourceinstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/remove-capacity | RemoveCapacityFromResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiRemoveCustomDNSFromResourceInstance**](docs/InventoryApiAPI.md#inventoryapiremovecustomdnsfromresourceinstance) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/{resourceKey}/instance/{instanceId}/custom-dns | RemoveCustomDNSFromResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiRestartResourceInstance**](docs/InventoryApiAPI.md#inventoryapirestartresourceinstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/restart | RestartResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiRestoreResourceInstance**](docs/InventoryApiAPI.md#inventoryapirestoreresourceinstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/restore | RestoreResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiRestoreResourceInstanceFromSnapshot**](docs/InventoryApiAPI.md#inventoryapirestoreresourceinstancefromsnapshot) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/snapshot/{snapshotId}/restore | RestoreResourceInstanceFromSnapshot inventory-api
*InventoryApiAPI* | [**InventoryApiResumeSubscription**](docs/InventoryApiAPI.md#inventoryapiresumesubscription) | **Put** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/{id}/resume | ResumeSubscription inventory-api
*InventoryApiAPI* | [**InventoryApiSearchInventory**](docs/InventoryApiAPI.md#inventoryapisearchinventory) | **Post** /2022-09-01-00/fleet/search-inventory | SearchInventory inventory-api
*InventoryApiAPI* | [**InventoryApiSearchServiceInventory**](docs/InventoryApiAPI.md#inventoryapisearchserviceinventory) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/search-inventory | SearchServiceInventory inventory-api
*InventoryApiAPI* | [**InventoryApiStartResourceInstance**](docs/InventoryApiAPI.md#inventoryapistartresourceinstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/start | StartResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiStopResourceInstance**](docs/InventoryApiAPI.md#inventoryapistopresourceinstance) | **Post** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId}/stop | StopResourceInstance inventory-api
*InventoryApiAPI* | [**InventoryApiSuspendSubscription**](docs/InventoryApiAPI.md#inventoryapisuspendsubscription) | **Put** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/{id}/suspend | SuspendSubscription inventory-api
*InventoryApiAPI* | [**InventoryApiSuspendUser**](docs/InventoryApiAPI.md#inventoryapisuspenduser) | **Put** /2022-09-01-00/fleet/user/{userId}/suspend | SuspendUser inventory-api
*InventoryApiAPI* | [**InventoryApiTerminateSubscription**](docs/InventoryApiAPI.md#inventoryapiterminatesubscription) | **Delete** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/subscription/{id} | TerminateSubscription inventory-api
*InventoryApiAPI* | [**InventoryApiUnsuspendUser**](docs/InventoryApiAPI.md#inventoryapiunsuspenduser) | **Put** /2022-09-01-00/fleet/user/{userId}/unsuspend | UnsuspendUser inventory-api
*InventoryApiAPI* | [**InventoryApiUpdateResourceInstance**](docs/InventoryApiAPI.md#inventoryapiupdateresourceinstance) | **Patch** /2022-09-01-00/fleet/service/{serviceId}/environment/{environmentId}/instance/{instanceId} | UpdateResourceInstance inventory-api
*NotificationsApiAPI* | [**NotificationsApiCreateNotificationChannel**](docs/NotificationsApiAPI.md#notificationsapicreatenotificationchannel) | **Post** /2022-09-01-00/fleet/notification-channel | CreateNotificationChannel notifications-api
*NotificationsApiAPI* | [**NotificationsApiDeleteNotificationChannel**](docs/NotificationsApiAPI.md#notificationsapideletenotificationchannel) | **Delete** /2022-09-01-00/fleet/notification-channel/{id} | DeleteNotificationChannel notifications-api
*NotificationsApiAPI* | [**NotificationsApiDescribeNotificationChannel**](docs/NotificationsApiAPI.md#notificationsapidescribenotificationchannel) | **Get** /2022-09-01-00/fleet/notification-channel/{id} | DescribeNotificationChannel notifications-api
*NotificationsApiAPI* | [**NotificationsApiListNotificationChannels**](docs/NotificationsApiAPI.md#notificationsapilistnotificationchannels) | **Get** /2022-09-01-00/fleet/notification-channel | ListNotificationChannels notifications-api
*NotificationsApiAPI* | [**NotificationsApiUpdateNotificationChannel**](docs/NotificationsApiAPI.md#notificationsapiupdatenotificationchannel) | **Patch** /2022-09-01-00/fleet/notification-channel/{id} | UpdateNotificationChannel notifications-api
*OperationsApiAPI* | [**OperationsApiDeploymentCellHealth**](docs/OperationsApiAPI.md#operationsapideploymentcellhealth) | **Get** /2022-09-01-00/fleet/operations/deployment-cell-health | DeploymentCellHealth operations-api
*OperationsApiAPI* | [**OperationsApiListEvents**](docs/OperationsApiAPI.md#operationsapilistevents) | **Get** /2022-09-01-00/fleet/operations/events | ListEvents operations-api
*OperationsApiAPI* | [**OperationsApiServiceHealth**](docs/OperationsApiAPI.md#operationsapiservicehealth) | **Get** /2022-09-01-00/fleet/operations/service-health | ServiceHealth operations-api
*WebhooksApiAPI* | [**WebhooksApiReceiveWebhook**](docs/WebhooksApiAPI.md#webhooksapireceivewebhook) | **Post** /2022-09-01-00/fleet/hooks/{id} | ReceiveWebhook webhooks-api


## Documentation For Models

 - [APIEntity](docs/APIEntity.md)
 - [AccessSideUser](docs/AccessSideUser.md)
 - [AddCapacityToResourceInstanceRequestBody](docs/AddCapacityToResourceInstanceRequestBody.md)
 - [AddCustomDNSToResourceInstanceRequestBody](docs/AddCustomDNSToResourceInstanceRequestBody.md)
 - [BackupStatus](docs/BackupStatus.md)
 - [Channel](docs/Channel.md)
 - [ChannelSubscription](docs/ChannelSubscription.md)
 - [CloudProviderHealthSummary](docs/CloudProviderHealthSummary.md)
 - [ClusterEndpoint](docs/ClusterEndpoint.md)
 - [CreateConsumptionUserRequestBody](docs/CreateConsumptionUserRequestBody.md)
 - [CreateCustomNetworkRequestBody](docs/CreateCustomNetworkRequestBody.md)
 - [CreateCustomerOnboardingRequestBody](docs/CreateCustomerOnboardingRequestBody.md)
 - [CreateNotificationChannelRequestBody](docs/CreateNotificationChannelRequestBody.md)
 - [CreateProxyResourceInstanceRequestBody](docs/CreateProxyResourceInstanceRequestBody.md)
 - [CreateResourceInstanceRequestBody](docs/CreateResourceInstanceRequestBody.md)
 - [CreateResourceInstanceResponseBody](docs/CreateResourceInstanceResponseBody.md)
 - [CreateServicesOrchestrationRequestBody](docs/CreateServicesOrchestrationRequestBody.md)
 - [CreateUpgradePathRequestBody](docs/CreateUpgradePathRequestBody.md)
 - [CustomDNSEndpoint](docs/CustomDNSEndpoint.md)
 - [CustomNetworkFleetDetail](docs/CustomNetworkFleetDetail.md)
 - [CustomNetworkResourceDetail](docs/CustomNetworkResourceDetail.md)
 - [CustomerOnboarding](docs/CustomerOnboarding.md)
 - [DeploymentCellHealthDetail](docs/DeploymentCellHealthDetail.md)
 - [DeploymentCellHealthSummary](docs/DeploymentCellHealthSummary.md)
 - [DeploymentCellSearchRecord](docs/DeploymentCellSearchRecord.md)
 - [DescribeInventorySummaryResult](docs/DescribeInventorySummaryResult.md)
 - [DescribeResourceInstanceResult](docs/DescribeResourceInstanceResult.md)
 - [DescribeResourceInstanceSnapshotFromTimeRequestBody](docs/DescribeResourceInstanceSnapshotFromTimeRequestBody.md)
 - [DescribeServiceOfferingResourceResult](docs/DescribeServiceOfferingResourceResult.md)
 - [DescribeServiceOfferingResult](docs/DescribeServiceOfferingResult.md)
 - [DescribeServiceWorkflowResult](docs/DescribeServiceWorkflowResult.md)
 - [DescribeServiceWorkflowSummaryResult](docs/DescribeServiceWorkflowSummaryResult.md)
 - [DescribeServicesOrchestrationResult](docs/DescribeServicesOrchestrationResult.md)
 - [DescribeSubscriptionRequestResult](docs/DescribeSubscriptionRequestResult.md)
 - [DetailedNodeHealthResult](docs/DetailedNodeHealthResult.md)
 - [EmailConfiguration](docs/EmailConfiguration.md)
 - [EnableFleetFeatureRequestBody](docs/EnableFleetFeatureRequestBody.md)
 - [EndCustomerEvent](docs/EndCustomerEvent.md)
 - [Error](docs/Error.md)
 - [EventsPerResource](docs/EventsPerResource.md)
 - [EventsPerWorkflowStep](docs/EventsPerWorkflowStep.md)
 - [FailoverResourceInstanceRequestBody](docs/FailoverResourceInstanceRequestBody.md)
 - [FleetCreateInstanceSnapshotResult](docs/FleetCreateInstanceSnapshotResult.md)
 - [FleetCustomNetwork](docs/FleetCustomNetwork.md)
 - [FleetCustomNetworkInstance](docs/FleetCustomNetworkInstance.md)
 - [FleetDescribeEventResult](docs/FleetDescribeEventResult.md)
 - [FleetDescribeHostClusterResult](docs/FleetDescribeHostClusterResult.md)
 - [FleetDescribeInstanceSnapshotFromTimeResult](docs/FleetDescribeInstanceSnapshotFromTimeResult.md)
 - [FleetDescribeInstanceSnapshotResult](docs/FleetDescribeInstanceSnapshotResult.md)
 - [FleetDescribeServicesOrchestrationResult](docs/FleetDescribeServicesOrchestrationResult.md)
 - [FleetDescribeSubscriptionResult](docs/FleetDescribeSubscriptionResult.md)
 - [FleetDescribeUserResult](docs/FleetDescribeUserResult.md)
 - [FleetFeature](docs/FleetFeature.md)
 - [FleetGenerateTokenForHostClusterDashboardResult](docs/FleetGenerateTokenForHostClusterDashboardResult.md)
 - [FleetListAllUsersResult](docs/FleetListAllUsersResult.md)
 - [FleetListCustomNetworksResult](docs/FleetListCustomNetworksResult.md)
 - [FleetListEventsResult](docs/FleetListEventsResult.md)
 - [FleetListHostClustersResult](docs/FleetListHostClustersResult.md)
 - [FleetListInstanceSnapshotResult](docs/FleetListInstanceSnapshotResult.md)
 - [FleetListLinkedInstancesResult](docs/FleetListLinkedInstancesResult.md)
 - [FleetListSubscriptionsResult](docs/FleetListSubscriptionsResult.md)
 - [FleetListUsersResult](docs/FleetListUsersResult.md)
 - [FleetNetworkFeaturesConfiguration](docs/FleetNetworkFeaturesConfiguration.md)
 - [GetWorkflowEventsResult](docs/GetWorkflowEventsResult.md)
 - [HelmPackage](docs/HelmPackage.md)
 - [HelmPackageInstallations](docs/HelmPackageInstallations.md)
 - [HostCluster](docs/HostCluster.md)
 - [InputParameterEntity](docs/InputParameterEntity.md)
 - [InstanceHealthSummary](docs/InstanceHealthSummary.md)
 - [InstanceUpgrade](docs/InstanceUpgrade.md)
 - [IntegrationStatus](docs/IntegrationStatus.md)
 - [IntegrationsHealth](docs/IntegrationsHealth.md)
 - [InventoryDescribeServiceOfferingResourceResult](docs/InventoryDescribeServiceOfferingResourceResult.md)
 - [InventoryDescribeServiceOfferingResult](docs/InventoryDescribeServiceOfferingResult.md)
 - [InventoryListServiceOfferingsResult](docs/InventoryListServiceOfferingsResult.md)
 - [KubernetesDashboardEndpoint](docs/KubernetesDashboardEndpoint.md)
 - [ListCustomerOnboardingResult](docs/ListCustomerOnboardingResult.md)
 - [ListCustomerOnboardingStagesResult](docs/ListCustomerOnboardingStagesResult.md)
 - [ListEligibleInstancesPerUpgradeResult](docs/ListEligibleInstancesPerUpgradeResult.md)
 - [ListEndCustomerEventsResult](docs/ListEndCustomerEventsResult.md)
 - [ListFleetFeaturesResult](docs/ListFleetFeaturesResult.md)
 - [ListFleetResourceInstancesResultInternal](docs/ListFleetResourceInstancesResultInternal.md)
 - [ListHelmPackageInstallationsResult](docs/ListHelmPackageInstallationsResult.md)
 - [ListHostClustersResult](docs/ListHostClustersResult.md)
 - [ListNotificationChannelsResult](docs/ListNotificationChannelsResult.md)
 - [ListOrganizationsResult](docs/ListOrganizationsResult.md)
 - [ListResourcesRequestBody](docs/ListResourcesRequestBody.md)
 - [ListResourcesResult](docs/ListResourcesResult.md)
 - [ListServiceOfferingsResult](docs/ListServiceOfferingsResult.md)
 - [ListServiceProviderEventsResult](docs/ListServiceProviderEventsResult.md)
 - [ListServiceWorkflowsResult](docs/ListServiceWorkflowsResult.md)
 - [ListSubscriptionRequestsResult](docs/ListSubscriptionRequestsResult.md)
 - [ListUpgradePathsResult](docs/ListUpgradePathsResult.md)
 - [ModifyServicesOrchestrationRequestBody](docs/ModifyServicesOrchestrationRequestBody.md)
 - [NodeHealthSummary](docs/NodeHealthSummary.md)
 - [NodeNetworkTopologyResult](docs/NodeNetworkTopologyResult.md)
 - [NodeVMInfoResult](docs/NodeVMInfoResult.md)
 - [NotificationSearchRecord](docs/NotificationSearchRecord.md)
 - [OfferingBillingPlan](docs/OfferingBillingPlan.md)
 - [OnboardingStage](docs/OnboardingStage.md)
 - [OneOffPatchResourceInstanceRequestBody](docs/OneOffPatchResourceInstanceRequestBody.md)
 - [Organization](docs/Organization.md)
 - [OutputParameterEntity](docs/OutputParameterEntity.md)
 - [PagerDutyConfiguration](docs/PagerDutyConfiguration.md)
 - [ProxyEndpoint](docs/ProxyEndpoint.md)
 - [ProxyInstanceSearchRecord](docs/ProxyInstanceSearchRecord.md)
 - [RecentDeploymentFailureStatus](docs/RecentDeploymentFailureStatus.md)
 - [RegionalHealthSummary](docs/RegionalHealthSummary.md)
 - [RemoveCapacityFromResourceInstanceRequestBody](docs/RemoveCapacityFromResourceInstanceRequestBody.md)
 - [Resource](docs/Resource.md)
 - [ResourceCapability](docs/ResourceCapability.md)
 - [ResourceDeploymentStatus](docs/ResourceDeploymentStatus.md)
 - [ResourceEntity](docs/ResourceEntity.md)
 - [ResourceHealthSummary](docs/ResourceHealthSummary.md)
 - [ResourceInstance](docs/ResourceInstance.md)
 - [ResourceInstanceSearchRecord](docs/ResourceInstanceSearchRecord.md)
 - [ResourceNetworkTopologyResult](docs/ResourceNetworkTopologyResult.md)
 - [ResourceSearchRecord](docs/ResourceSearchRecord.md)
 - [ResourceVersionSummary](docs/ResourceVersionSummary.md)
 - [RestoreResourceInstanceFromSnapshotRequestBody](docs/RestoreResourceInstanceFromSnapshotRequestBody.md)
 - [RestoreResourceInstanceRequestBody](docs/RestoreResourceInstanceRequestBody.md)
 - [SearchInventoryResult](docs/SearchInventoryResult.md)
 - [SearchRecord](docs/SearchRecord.md)
 - [SearchServiceInventoryRequestBody](docs/SearchServiceInventoryRequestBody.md)
 - [SearchServiceInventoryResult](docs/SearchServiceInventoryResult.md)
 - [ServerlessProxySearchRecord](docs/ServerlessProxySearchRecord.md)
 - [ServiceAssets](docs/ServiceAssets.md)
 - [ServiceDeploymentDetails](docs/ServiceDeploymentDetails.md)
 - [ServiceHealthSummary](docs/ServiceHealthSummary.md)
 - [ServiceModelFeatureDetail](docs/ServiceModelFeatureDetail.md)
 - [ServiceOffering](docs/ServiceOffering.md)
 - [ServicePlanSearchRecord](docs/ServicePlanSearchRecord.md)
 - [ServiceProviderEvent](docs/ServiceProviderEvent.md)
 - [ServiceProviderEventSummary](docs/ServiceProviderEventSummary.md)
 - [ServiceSearchRecord](docs/ServiceSearchRecord.md)
 - [ServiceWorkflow](docs/ServiceWorkflow.md)
 - [SlackConfiguration](docs/SlackConfiguration.md)
 - [StartResourceInstanceRequestBody](docs/StartResourceInstanceRequestBody.md)
 - [SubscriptionSearchRecord](docs/SubscriptionSearchRecord.md)
 - [UpdateCustomerOnboardingRequestBody](docs/UpdateCustomerOnboardingRequestBody.md)
 - [UpdateNotificationChannelRequestBody](docs/UpdateNotificationChannelRequestBody.md)
 - [UpdateResourceInstanceRequestBody](docs/UpdateResourceInstanceRequestBody.md)
 - [UpdateServiceWorkflowRequestBody](docs/UpdateServiceWorkflowRequestBody.md)
 - [UpgradePath](docs/UpgradePath.md)
 - [UpgradePathSearchRecord](docs/UpgradePathSearchRecord.md)
 - [User](docs/User.md)
 - [UserSearchRecord](docs/UserSearchRecord.md)
 - [UserSubscription](docs/UserSubscription.md)
 - [WebhookConfiguration](docs/WebhookConfiguration.md)
 - [WorkflowEvent](docs/WorkflowEvent.md)
 - [WorkflowFailure](docs/WorkflowFailure.md)
 - [WorkflowSearchRecord](docs/WorkflowSearchRecord.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### api_key_header_Authorization

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), fleet.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



