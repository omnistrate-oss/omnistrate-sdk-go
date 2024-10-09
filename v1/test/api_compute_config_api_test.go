/*
Omnistrate Registration API

Testing ComputeConfigApiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v1

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func Test_v1_ComputeConfigApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ComputeConfigApiAPIService ComputeConfigApiAddComputeInstanceType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiAddComputeInstanceType(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputeConfigApiAPIService ComputeConfigApiCreateComputeConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiCreateComputeConfig(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputeConfigApiAPIService ComputeConfigApiDeleteComputeConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiDeleteComputeConfig(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputeConfigApiAPIService ComputeConfigApiDescribeComputeConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		resp, httpRes, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiDescribeComputeConfig(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputeConfigApiAPIService ComputeConfigApiListComputeConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiListComputeConfig(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputeConfigApiAPIService ComputeConfigApiListComputeInstanceTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var cloudProviderName string

		resp, httpRes, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiListComputeInstanceTypes(context.Background(), serviceId, cloudProviderName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputeConfigApiAPIService ComputeConfigApiRemoveComputeInstanceType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiRemoveComputeInstanceType(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputeConfigApiAPIService ComputeConfigApiUpdateComputeConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.ComputeConfigApiAPI.ComputeConfigApiUpdateComputeConfig(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}