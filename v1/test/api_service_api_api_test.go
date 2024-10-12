/*
Omnistrate Registration API

Testing ServiceApiApiAPIService

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

func Test_v1_ServiceApiApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceApiApiAPIService ServiceApiApiCreateServiceAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiCreateServiceAPI(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceApiApiAPIService ServiceApiApiDeleteServiceAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiDeleteServiceAPI(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceApiApiAPIService ServiceApiApiDeprecateServiceAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiDeprecateServiceAPI(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceApiApiAPIService ServiceApiApiDescribePendingChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		resp, httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiDescribePendingChanges(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceApiApiAPIService ServiceApiApiDescribeServiceAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		resp, httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiDescribeServiceAPI(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceApiApiAPIService ServiceApiApiDiscardPendingChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiDiscardPendingChanges(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceApiApiAPIService ServiceApiApiListServiceAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var serviceEnvironmentId string

		resp, httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiListServiceAPI(context.Background(), serviceId, serviceEnvironmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceApiApiAPIService ServiceApiApiReleaseServiceAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiReleaseServiceAPI(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceApiApiAPIService ServiceApiApiUpdateServiceAPI", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ServiceApiApiAPI.ServiceApiApiUpdateServiceAPI(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
