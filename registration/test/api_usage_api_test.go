/*
Omnistrate Registration API

Testing UsageApiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package registration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go"
)

func Test_registration_UsageApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsageApiAPIService UsageApiGetCurrentPlanUsage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UsageApiAPI.UsageApiGetCurrentPlanUsage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
