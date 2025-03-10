/*
Omnistrate Fleet API

Testing InventoryApiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fleet

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func Test_fleet_InventoryApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InventoryApiAPIService InventoryApiAddCapacityToResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiAddCapacityToResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiAddCustomDNSToResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var resourceKey string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiAddCustomDNSToResourceInstance(context.Background(), serviceId, environmentId, resourceKey, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiApproveSubscriptionRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiApproveSubscriptionRequest(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiCancelUpgradePath", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var productTierId string
		var upgradePathId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiCancelUpgradePath(context.Background(), serviceId, productTierId, upgradePathId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiCreateConsumptionUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiCreateConsumptionUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiCreateProxyResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceProviderId string
		var serviceKey string
		var serviceAPIVersion string
		var serviceEnvironmentKey string
		var serviceModelKey string
		var productTierKey string
		var sourceResourceKey string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiCreateProxyResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, sourceResourceKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiCreateResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceProviderId string
		var serviceKey string
		var serviceAPIVersion string
		var serviceEnvironmentKey string
		var serviceModelKey string
		var productTierKey string
		var resourceKey string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiCreateResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, resourceKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiCreateResourceInstanceSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiCreateResourceInstanceSnapshot(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiCreateServicesOrchestration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiCreateServicesOrchestration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiCreateUpgradePath", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var productTierId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiCreateUpgradePath(context.Background(), serviceId, productTierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDebugResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDebugResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDeleteProxyResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceProviderId string
		var serviceKey string
		var serviceAPIVersion string
		var serviceEnvironmentKey string
		var serviceModelKey string
		var productTierKey string
		var id string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiDeleteProxyResourceInstance(context.Background(), serviceProviderId, serviceKey, serviceAPIVersion, serviceEnvironmentKey, serviceModelKey, productTierKey, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDeleteResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiDeleteResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDeleteResourceInstanceSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string
		var snapshotId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiDeleteResourceInstanceSnapshot(context.Background(), serviceId, environmentId, instanceId, snapshotId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDeleteServicesOrchestration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiDeleteServicesOrchestration(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiDeleteUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDenySubscriptionRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiDenySubscriptionRequest(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeHostCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeHostCluster(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeInstanceEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string
		var id string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeInstanceEvent(context.Background(), serviceId, environmentId, instanceId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeInventorySummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeInventorySummary(context.Background(), serviceId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeOrgUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeOrgUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var organizationId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeOrganization(context.Background(), serviceId, environmentId, organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var serviceModelId string
		var productTierId string
		var resourceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeResource(context.Background(), serviceId, environmentId, serviceModelId, productTierId, resourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeResourceInstanceSnapshotFromTime", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeResourceInstanceSnapshotFromTime(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeServiceOffering", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeServiceOffering(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeServiceOfferingResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var resourceId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeServiceOfferingResource(context.Background(), serviceId, resourceId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeServicesOrchestration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeServicesOrchestration(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeSubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeSubscription(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeSubscriptionRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeSubscriptionRequest(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeUpgradePath", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var productTierId string
		var upgradePathId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeUpgradePath(context.Background(), serviceId, productTierId, upgradePathId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiDescribeUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var userId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiDescribeUser(context.Background(), serviceId, environmentId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiEnableResourceInstanceManualOverride", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiEnableResourceInstanceManualOverride(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiFailoverResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiFailoverResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiGenerateTokenForHostClusterDashboard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiGenerateTokenForHostClusterDashboard(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListActiveOrganizations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListActiveOrganizations(context.Background(), serviceId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListAllUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListAllUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListDependentComponents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListDependentComponents(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListEligibleInstancesPerUpgrade", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var productTierId string
		var upgradePathId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListEligibleInstancesPerUpgrade(context.Background(), serviceId, productTierId, upgradePathId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListHostClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListHostClusters(context.Background(), serviceId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListInstanceEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListInstanceEvents(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListLinkedInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListLinkedInstances(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListResourceInstanceSnapshots", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListResourceInstanceSnapshots(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListResourceInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListResourceInstances(context.Background(), serviceId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var serviceModelId string
		var productTierId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListResources(context.Background(), serviceId, environmentId, serviceModelId, productTierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListServiceOfferings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListServiceOfferings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListServicesOrchestrations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListServicesOrchestrations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListSubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListSubscription(context.Background(), serviceId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListSubscriptionRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListSubscriptionRequests(context.Background(), serviceId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListUpgradePaths", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var productTierId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListUpgradePaths(context.Background(), serviceId, productTierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiListUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiListUsers(context.Background(), serviceId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiManageUpgradePath", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var productTierId string
		var upgradePathId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiManageUpgradePath(context.Background(), serviceId, productTierId, upgradePathId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiModifyServicesOrchestration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiModifyServicesOrchestration(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiRemoveCapacityFromResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiRemoveCapacityFromResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiRemoveCustomDNSFromResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var resourceKey string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiRemoveCustomDNSFromResourceInstance(context.Background(), serviceId, environmentId, resourceKey, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiRestartResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiRestartResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiRestoreResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiRestoreResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiRestoreResourceInstanceFromSnapshot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var snapshotId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiRestoreResourceInstanceFromSnapshot(context.Background(), serviceId, environmentId, snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiResumeSubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiResumeSubscription(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiSearchInventory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiSearchInventory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiSearchServiceInventory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string

		resp, httpRes, err := apiClient.InventoryApiAPI.InventoryApiSearchServiceInventory(context.Background(), serviceId, environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiStartResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiStartResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiStopResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiStopResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiSuspendSubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiSuspendSubscription(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiSuspendUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiSuspendUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiTerminateSubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var id string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiTerminateSubscription(context.Background(), serviceId, environmentId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiUnsuspendUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiUnsuspendUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiAPIService InventoryApiUpdateResourceInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var environmentId string
		var instanceId string

		httpRes, err := apiClient.InventoryApiAPI.InventoryApiUpdateResourceInstance(context.Background(), serviceId, environmentId, instanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
