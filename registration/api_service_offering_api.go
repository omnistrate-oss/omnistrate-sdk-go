/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type ServiceOfferingApiAPI interface {

	/*
	ServiceOfferingApiDescribeServiceOffering DescribeServiceOffering service-offering-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId The service ID
	@return ApiServiceOfferingApiDescribeServiceOfferingRequest
	*/
	ServiceOfferingApiDescribeServiceOffering(ctx context.Context, serviceId string) ApiServiceOfferingApiDescribeServiceOfferingRequest

	// ServiceOfferingApiDescribeServiceOfferingExecute executes the request
	//  @return DescribeServiceOfferingResult
	ServiceOfferingApiDescribeServiceOfferingExecute(r ApiServiceOfferingApiDescribeServiceOfferingRequest) (*DescribeServiceOfferingResult, *http.Response, error)

	/*
	ServiceOfferingApiDescribeServiceOfferingResource DescribeServiceOfferingResource service-offering-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId The service ID
	@param resourceId The resource ID
	@param instanceId The instance ID
	@return ApiServiceOfferingApiDescribeServiceOfferingResourceRequest
	*/
	ServiceOfferingApiDescribeServiceOfferingResource(ctx context.Context, serviceId string, resourceId string, instanceId string) ApiServiceOfferingApiDescribeServiceOfferingResourceRequest

	// ServiceOfferingApiDescribeServiceOfferingResourceExecute executes the request
	//  @return DescribeServiceOfferingResourceResult
	ServiceOfferingApiDescribeServiceOfferingResourceExecute(r ApiServiceOfferingApiDescribeServiceOfferingResourceRequest) (*DescribeServiceOfferingResourceResult, *http.Response, error)

	/*
	ServiceOfferingApiListServiceOffering ListServiceOffering service-offering-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiServiceOfferingApiListServiceOfferingRequest
	*/
	ServiceOfferingApiListServiceOffering(ctx context.Context) ApiServiceOfferingApiListServiceOfferingRequest

	// ServiceOfferingApiListServiceOfferingExecute executes the request
	//  @return ListServiceOfferingsResult
	ServiceOfferingApiListServiceOfferingExecute(r ApiServiceOfferingApiListServiceOfferingRequest) (*ListServiceOfferingsResult, *http.Response, error)
}

// ServiceOfferingApiAPIService ServiceOfferingApiAPI service
type ServiceOfferingApiAPIService service

type ApiServiceOfferingApiDescribeServiceOfferingRequest struct {
	ctx context.Context
	ApiService ServiceOfferingApiAPI
	serviceId string
	visibility *string
	environmentType *string
}

// The visibility of service offering
func (r ApiServiceOfferingApiDescribeServiceOfferingRequest) Visibility(visibility string) ApiServiceOfferingApiDescribeServiceOfferingRequest {
	r.visibility = &visibility
	return r
}

// The environment type to filter by
func (r ApiServiceOfferingApiDescribeServiceOfferingRequest) EnvironmentType(environmentType string) ApiServiceOfferingApiDescribeServiceOfferingRequest {
	r.environmentType = &environmentType
	return r
}

func (r ApiServiceOfferingApiDescribeServiceOfferingRequest) Execute() (*DescribeServiceOfferingResult, *http.Response, error) {
	return r.ApiService.ServiceOfferingApiDescribeServiceOfferingExecute(r)
}

/*
ServiceOfferingApiDescribeServiceOffering DescribeServiceOffering service-offering-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId The service ID
 @return ApiServiceOfferingApiDescribeServiceOfferingRequest
*/
func (a *ServiceOfferingApiAPIService) ServiceOfferingApiDescribeServiceOffering(ctx context.Context, serviceId string) ApiServiceOfferingApiDescribeServiceOfferingRequest {
	return ApiServiceOfferingApiDescribeServiceOfferingRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return DescribeServiceOfferingResult
func (a *ServiceOfferingApiAPIService) ServiceOfferingApiDescribeServiceOfferingExecute(r ApiServiceOfferingApiDescribeServiceOfferingRequest) (*DescribeServiceOfferingResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DescribeServiceOfferingResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceOfferingApiAPIService.ServiceOfferingApiDescribeServiceOffering")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/service-offering/{serviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.visibility != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "visibility", r.visibility, "form", "")
	}
	if r.environmentType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "environmentType", r.environmentType, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.goa.error"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServiceOfferingApiDescribeServiceOfferingResourceRequest struct {
	ctx context.Context
	ApiService ServiceOfferingApiAPI
	serviceId string
	resourceId string
	instanceId string
}

func (r ApiServiceOfferingApiDescribeServiceOfferingResourceRequest) Execute() (*DescribeServiceOfferingResourceResult, *http.Response, error) {
	return r.ApiService.ServiceOfferingApiDescribeServiceOfferingResourceExecute(r)
}

/*
ServiceOfferingApiDescribeServiceOfferingResource DescribeServiceOfferingResource service-offering-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId The service ID
 @param resourceId The resource ID
 @param instanceId The instance ID
 @return ApiServiceOfferingApiDescribeServiceOfferingResourceRequest
*/
func (a *ServiceOfferingApiAPIService) ServiceOfferingApiDescribeServiceOfferingResource(ctx context.Context, serviceId string, resourceId string, instanceId string) ApiServiceOfferingApiDescribeServiceOfferingResourceRequest {
	return ApiServiceOfferingApiDescribeServiceOfferingResourceRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
		resourceId: resourceId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return DescribeServiceOfferingResourceResult
func (a *ServiceOfferingApiAPIService) ServiceOfferingApiDescribeServiceOfferingResourceExecute(r ApiServiceOfferingApiDescribeServiceOfferingResourceRequest) (*DescribeServiceOfferingResourceResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DescribeServiceOfferingResourceResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceOfferingApiAPIService.ServiceOfferingApiDescribeServiceOfferingResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/service-offering/{serviceId}/resource/{resourceId}/instance/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterValueToString(r.resourceId, "resourceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", url.PathEscape(parameterValueToString(r.instanceId, "instanceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.goa.error"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServiceOfferingApiListServiceOfferingRequest struct {
	ctx context.Context
	ApiService ServiceOfferingApiAPI
	orgId *string
	visibility *string
	environmentType *string
}

// Org Id
func (r ApiServiceOfferingApiListServiceOfferingRequest) OrgId(orgId string) ApiServiceOfferingApiListServiceOfferingRequest {
	r.orgId = &orgId
	return r
}

// The visibility of service offering
func (r ApiServiceOfferingApiListServiceOfferingRequest) Visibility(visibility string) ApiServiceOfferingApiListServiceOfferingRequest {
	r.visibility = &visibility
	return r
}

// The environment type to filter by
func (r ApiServiceOfferingApiListServiceOfferingRequest) EnvironmentType(environmentType string) ApiServiceOfferingApiListServiceOfferingRequest {
	r.environmentType = &environmentType
	return r
}

func (r ApiServiceOfferingApiListServiceOfferingRequest) Execute() (*ListServiceOfferingsResult, *http.Response, error) {
	return r.ApiService.ServiceOfferingApiListServiceOfferingExecute(r)
}

/*
ServiceOfferingApiListServiceOffering ListServiceOffering service-offering-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServiceOfferingApiListServiceOfferingRequest
*/
func (a *ServiceOfferingApiAPIService) ServiceOfferingApiListServiceOffering(ctx context.Context) ApiServiceOfferingApiListServiceOfferingRequest {
	return ApiServiceOfferingApiListServiceOfferingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListServiceOfferingsResult
func (a *ServiceOfferingApiAPIService) ServiceOfferingApiListServiceOfferingExecute(r ApiServiceOfferingApiListServiceOfferingRequest) (*ListServiceOfferingsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListServiceOfferingsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceOfferingApiAPIService.ServiceOfferingApiListServiceOffering")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/service-offering"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.orgId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orgId", r.orgId, "form", "")
	}
	if r.visibility != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "visibility", r.visibility, "form", "")
	}
	if r.environmentType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "environmentType", r.environmentType, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.goa.error"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
