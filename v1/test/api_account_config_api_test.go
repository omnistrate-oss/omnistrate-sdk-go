/*
Omnistrate Registration API

Testing AccountConfigApiAPIService

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

func Test_v1_AccountConfigApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountConfigApiAPIService AccountConfigApiAccountConfigIdentityID", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiAccountConfigIdentityID(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountConfigApiAPIService AccountConfigApiCreateAccountConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiCreateAccountConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountConfigApiAPIService AccountConfigApiDeleteAccountConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiDeleteAccountConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountConfigApiAPIService AccountConfigApiDescribeAccountConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountConfigApiAPIService AccountConfigApiDescribeAccountConfigByAWSAccountID", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var awsAccountID string

		resp, httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByAWSAccountID(context.Background(), awsAccountID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountConfigApiAPIService AccountConfigApiDescribeAccountConfigByGCPProjectID", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gcpProjectID string

		resp, httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByGCPProjectID(context.Background(), gcpProjectID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountConfigApiAPIService AccountConfigApiListAccountConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudProviderName string

		resp, httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiListAccountConfig(context.Background(), cloudProviderName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountConfigApiAPIService AccountConfigApiListBYOAConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiListBYOAConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountConfigApiAPIService AccountConfigApiVerifyAccountConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AccountConfigApiAPI.AccountConfigApiVerifyAccountConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
