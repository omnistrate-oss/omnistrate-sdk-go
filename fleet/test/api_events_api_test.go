/*
Omnistrate Fleet API

Testing EventsApiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fleet

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func Test_fleet_EventsApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EventsApiAPIService EventsApiAcknowledgeEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.EventsApiAPI.EventsApiAcknowledgeEvent(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiAPIService EventsApiListEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsApiAPI.EventsApiListEvents(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
