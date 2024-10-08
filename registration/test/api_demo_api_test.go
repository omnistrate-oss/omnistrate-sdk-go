/*
Omnistrate Registration API

Testing DemoApiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package registration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func Test_registration_DemoApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DemoApiAPIService DemoApiDemo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.DemoApiAPI.DemoApiDemo(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
