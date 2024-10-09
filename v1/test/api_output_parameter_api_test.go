/*
Omnistrate Registration API

Testing OutputParameterApiAPIService

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

func Test_v1_OutputParameterApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OutputParameterApiAPIService OutputParameterApiCreateOutputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.OutputParameterApiAPI.OutputParameterApiCreateOutputParameter(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutputParameterApiAPIService OutputParameterApiDeleteOutputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.OutputParameterApiAPI.OutputParameterApiDeleteOutputParameter(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutputParameterApiAPIService OutputParameterApiDescribeOutputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		resp, httpRes, err := apiClient.OutputParameterApiAPI.OutputParameterApiDescribeOutputParameter(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutputParameterApiAPIService OutputParameterApiListOutputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var resourceId string

		resp, httpRes, err := apiClient.OutputParameterApiAPI.OutputParameterApiListOutputParameter(context.Background(), serviceId, resourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutputParameterApiAPIService OutputParameterApiUpdateOutputParameter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		httpRes, err := apiClient.OutputParameterApiAPI.OutputParameterApiUpdateOutputParameter(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}