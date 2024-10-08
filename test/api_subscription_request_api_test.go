/*
Omnistrate Registration API

Testing SubscriptionRequestApiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package omnistrategosdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func Test_omnistrategosdk_SubscriptionRequestApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubscriptionRequestApiAPIService SubscriptionRequestApiCancelSubscriptionRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.SubscriptionRequestApiAPI.SubscriptionRequestApiCancelSubscriptionRequest(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionRequestApiAPIService SubscriptionRequestApiCreateSubscriptionRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SubscriptionRequestApiAPI.SubscriptionRequestApiCreateSubscriptionRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionRequestApiAPIService SubscriptionRequestApiDescribeSubscriptionRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SubscriptionRequestApiAPI.SubscriptionRequestApiDescribeSubscriptionRequest(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionRequestApiAPIService SubscriptionRequestApiListSubscriptionRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SubscriptionRequestApiAPI.SubscriptionRequestApiListSubscriptionRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
