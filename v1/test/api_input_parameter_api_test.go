/*
Omnistrate Registration API

Testing InputParameterApiAPIService

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

func Test_v1_InputParameterApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InputParameterApiAPIService InputParameterApiCreateInputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.InputParameterApiAPI.InputParameterApiCreateInputParameter(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InputParameterApiAPIService InputParameterApiDeleteInputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.InputParameterApiAPI.InputParameterApiDeleteInputParameter(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InputParameterApiAPIService InputParameterApiDescribeInputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		resp, httpRes, err := apiClient.InputParameterApiAPI.InputParameterApiDescribeInputParameter(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InputParameterApiAPIService InputParameterApiListInputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var resourceId string

		resp, httpRes, err := apiClient.InputParameterApiAPI.InputParameterApiListInputParameter(context.Background(), serviceId, resourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InputParameterApiAPIService InputParameterApiUpdateInputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.InputParameterApiAPI.InputParameterApiUpdateInputParameter(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
