/*
Omnistrate Registration API

Testing DeploymentConfigApiAPIService

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

func Test_v1_DeploymentConfigApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeploymentConfigApiAPIService DeploymentConfigApiCreateDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiCreateDeploymentConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigApiAPIService DeploymentConfigApiDeleteDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiDeleteDeploymentConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigApiAPIService DeploymentConfigApiDescribeDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiDescribeDeploymentConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigApiAPIService DeploymentConfigApiListDeploymentConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiListDeploymentConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigApiAPIService DeploymentConfigApiUpdateDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DeploymentConfigApiAPI.DeploymentConfigApiUpdateDeploymentConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
