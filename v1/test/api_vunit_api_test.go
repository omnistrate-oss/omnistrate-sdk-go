/*
Omnistrate Registration API

Testing VunitApiAPIService

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

func Test_v1_VunitApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VunitApiAPIService VunitApiDescribeNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		resp, httpRes, err := apiClient.VunitApiAPI.VunitApiDescribeNetwork(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VunitApiAPIService VunitApiDescribeVUnit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var id string

		resp, httpRes, err := apiClient.VunitApiAPI.VunitApiDescribeVUnit(context.Background(), serviceId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VunitApiAPIService VunitApiListVUnits", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var serviceModelId string

		resp, httpRes, err := apiClient.VunitApiAPI.VunitApiListVUnits(context.Background(), serviceId, serviceModelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
