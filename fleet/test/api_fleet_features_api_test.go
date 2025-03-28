/*
Omnistrate Fleet API

Testing FleetFeaturesApiAPIService

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

func Test_fleet_FleetFeaturesApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FleetFeaturesApiAPIService FleetFeaturesApiDescribeFleetFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var feature string

		resp, httpRes, err := apiClient.FleetFeaturesApiAPI.FleetFeaturesApiDescribeFleetFeature(context.Background(), feature).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetFeaturesApiAPIService FleetFeaturesApiDisableFleetFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.FleetFeaturesApiAPI.FleetFeaturesApiDisableFleetFeature(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetFeaturesApiAPIService FleetFeaturesApiEnableFleetFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.FleetFeaturesApiAPI.FleetFeaturesApiEnableFleetFeature(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetFeaturesApiAPIService FleetFeaturesApiListFleetFeatures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FleetFeaturesApiAPI.FleetFeaturesApiListFleetFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
