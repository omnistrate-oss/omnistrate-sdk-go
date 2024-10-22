/*
Omnistrate Fleet API

Testing FleetCustomNetworkApiAPIService

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

func Test_fleet_FleetCustomNetworkApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FleetCustomNetworkApiAPIService FleetCustomNetworkApiCreateCustomNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FleetCustomNetworkApiAPI.FleetCustomNetworkApiCreateCustomNetwork(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetCustomNetworkApiAPIService FleetCustomNetworkApiDeleteCustomNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.FleetCustomNetworkApiAPI.FleetCustomNetworkApiDeleteCustomNetwork(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetCustomNetworkApiAPIService FleetCustomNetworkApiDescribeCustomNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.FleetCustomNetworkApiAPI.FleetCustomNetworkApiDescribeCustomNetwork(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetCustomNetworkApiAPIService FleetCustomNetworkApiListCustomNetworks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FleetCustomNetworkApiAPI.FleetCustomNetworkApiListCustomNetworks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
