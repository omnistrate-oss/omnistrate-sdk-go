/*
Omnistrate Registration API

Testing SaasPortalApiAPIService

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

func Test_v1_SaasPortalApiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SaasPortalApiAPIService SaasPortalApiCreateSaaSPortalCustomDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SaasPortalApiAPI.SaasPortalApiCreateSaaSPortalCustomDomain(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SaasPortalApiAPIService SaasPortalApiDeleteSaaSPortalCustomDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentType string

		httpRes, err := apiClient.SaasPortalApiAPI.SaasPortalApiDeleteSaaSPortalCustomDomain(context.Background(), environmentType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SaasPortalApiAPIService SaasPortalApiListSaaSPortalCustomDomains", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SaasPortalApiAPI.SaasPortalApiListSaaSPortalCustomDomains(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SaasPortalApiAPIService SaasPortalApiListSaaSPortals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SaasPortalApiAPI.SaasPortalApiListSaaSPortals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SaasPortalApiAPIService SaasPortalApiUpdateSaaSPortal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentType string

		httpRes, err := apiClient.SaasPortalApiAPI.SaasPortalApiUpdateSaaSPortal(context.Background(), environmentType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SaasPortalApiAPIService SaasPortalApiUpdateSaaSPortalCustomDomain", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentType string

		httpRes, err := apiClient.SaasPortalApiAPI.SaasPortalApiUpdateSaaSPortalCustomDomain(context.Background(), environmentType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}