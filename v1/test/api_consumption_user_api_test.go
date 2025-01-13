/*
Omnistrate Registration API

Testing ConsumptionUserApiAPIService

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

func Test_v1_ConsumptionUserApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConsumptionUserApiAPIService ConsumptionUserApiDescribeUserBillingDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiDescribeUserBillingDetails(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumptionUserApiAPIService ConsumptionUserApiDescribeUsersBySubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiDescribeUsersBySubscription(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumptionUserApiAPIService ConsumptionUserApiInviteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiInviteUser(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumptionUserApiAPIService ConsumptionUserApiListAllSubscriptionUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiListAllSubscriptionUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumptionUserApiAPIService ConsumptionUserApiRevokeUserRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiRevokeUserRole(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConsumptionUserApiAPIService ConsumptionUserApiSignin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConsumptionUserApiAPI.ConsumptionUserApiSignin(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
