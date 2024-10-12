/*
Omnistrate Registration API

Testing ServiceOfferingApiAPIService

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

func Test_v1_ServiceOfferingApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceOfferingApiAPIService ServiceOfferingApiDescribeServiceOffering", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.ServiceOfferingApiAPI.ServiceOfferingApiDescribeServiceOffering(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceOfferingApiAPIService ServiceOfferingApiDescribeServiceOfferingResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var resourceId string
		var instanceId string

		resp, httpRes, err := apiClient.ServiceOfferingApiAPI.ServiceOfferingApiDescribeServiceOfferingResource(context.Background(), serviceId, resourceId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceOfferingApiAPIService ServiceOfferingApiListServiceOffering", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceOfferingApiAPI.ServiceOfferingApiListServiceOffering(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
