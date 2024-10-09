/*
Omnistrate Fleet API

Testing CustomerOnboardingsApiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fleet

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/fleet"
)

func Test_fleet_CustomerOnboardingsApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomerOnboardingsApiAPIService CustomerOnboardingsApiCreateCustomerOnboarding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiCreateCustomerOnboarding(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerOnboardingsApiAPIService CustomerOnboardingsApiDeleteCustomerOnboarding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiDeleteCustomerOnboarding(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerOnboardingsApiAPIService CustomerOnboardingsApiDescribeCustomerOnboarding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiDescribeCustomerOnboarding(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerOnboardingsApiAPIService CustomerOnboardingsApiListCustomerOnboardingStages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiListCustomerOnboardingStages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerOnboardingsApiAPIService CustomerOnboardingsApiListCustomerOnboardings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiListCustomerOnboardings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerOnboardingsApiAPIService CustomerOnboardingsApiUpdateCustomerOnboarding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CustomerOnboardingsApiAPI.CustomerOnboardingsApiUpdateCustomerOnboarding(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}