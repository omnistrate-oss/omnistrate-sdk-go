/*
Omnistrate Registration API

Testing AvailabilityZoneApiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v1

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func Test_v1_AvailabilityZoneApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AvailabilityZoneApiAPIService AvailabilityZoneApiDescribeAvailabilityZone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AvailabilityZoneApiAPI.AvailabilityZoneApiDescribeAvailabilityZone(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvailabilityZoneApiAPIService AvailabilityZoneApiGetAvailabilityZoneByCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var code string
		var cloudProviderName string

		resp, httpRes, err := apiClient.AvailabilityZoneApiAPI.AvailabilityZoneApiGetAvailabilityZoneByCode(context.Background(), code, cloudProviderName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvailabilityZoneApiAPIService AvailabilityZoneApiListAvailabilityZone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudProviderName string

		resp, httpRes, err := apiClient.AvailabilityZoneApiAPI.AvailabilityZoneApiListAvailabilityZone(context.Background(), cloudProviderName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvailabilityZoneApiAPIService AvailabilityZoneApiListAvailabilityZonesByRegionCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var regionCode string
		var cloudProviderName string

		resp, httpRes, err := apiClient.AvailabilityZoneApiAPI.AvailabilityZoneApiListAvailabilityZonesByRegionCode(context.Background(), regionCode, cloudProviderName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
