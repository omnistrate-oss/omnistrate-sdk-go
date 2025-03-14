/*
Omnistrate Registration API

Testing ServicePlanApiAPIService

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

func Test_v1_ServicePlanApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServicePlanApiAPIService ServicePlanApiGetServicePlan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var productTierId string

		resp, httpRes, err := apiClient.ServicePlanApiAPI.ServicePlanApiGetServicePlan(context.Background(), serviceId, productTierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePlanApiAPIService ServicePlanApiListServicePlans", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var serviceEnvironmentId string

		resp, httpRes, err := apiClient.ServicePlanApiAPI.ServicePlanApiListServicePlans(context.Background(), serviceId, serviceEnvironmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
