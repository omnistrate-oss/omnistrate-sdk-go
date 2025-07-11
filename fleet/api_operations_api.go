/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
	"reflect"
)


type OperationsApiAPI interface {

	/*
	OperationsApiDeploymentCellHealth DeploymentCellHealth operations-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOperationsApiDeploymentCellHealthRequest
	*/
	OperationsApiDeploymentCellHealth(ctx context.Context) ApiOperationsApiDeploymentCellHealthRequest

	// OperationsApiDeploymentCellHealthExecute executes the request
	//  @return DeploymentCellHealthDetail
	OperationsApiDeploymentCellHealthExecute(r ApiOperationsApiDeploymentCellHealthRequest) (*DeploymentCellHealthDetail, *http.Response, error)

	/*
	OperationsApiListEvents ListEvents operations-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOperationsApiListEventsRequest
	*/
	OperationsApiListEvents(ctx context.Context) ApiOperationsApiListEventsRequest

	// OperationsApiListEventsExecute executes the request
	//  @return ListServiceProviderEventsResult
	OperationsApiListEventsExecute(r ApiOperationsApiListEventsRequest) (*ListServiceProviderEventsResult, *http.Response, error)

	/*
	OperationsApiServiceHealth ServiceHealth operations-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOperationsApiServiceHealthRequest
	*/
	OperationsApiServiceHealth(ctx context.Context) ApiOperationsApiServiceHealthRequest

	// OperationsApiServiceHealthExecute executes the request
	//  @return ServiceHealthSummary
	OperationsApiServiceHealthExecute(r ApiOperationsApiServiceHealthRequest) (*ServiceHealthSummary, *http.Response, error)
}

// OperationsApiAPIService OperationsApiAPI service
type OperationsApiAPIService service

type ApiOperationsApiDeploymentCellHealthRequest struct {
	ctx context.Context
	ApiService OperationsApiAPI
	hostClusterID *string
	serviceID *string
	serviceEnvironmentID *string
}

// The host cluster ID to get the health for
func (r ApiOperationsApiDeploymentCellHealthRequest) HostClusterID(hostClusterID string) ApiOperationsApiDeploymentCellHealthRequest {
	r.hostClusterID = &hostClusterID
	return r
}

// The service ID to get the health for
func (r ApiOperationsApiDeploymentCellHealthRequest) ServiceID(serviceID string) ApiOperationsApiDeploymentCellHealthRequest {
	r.serviceID = &serviceID
	return r
}

// The service environment ID to get the health for
func (r ApiOperationsApiDeploymentCellHealthRequest) ServiceEnvironmentID(serviceEnvironmentID string) ApiOperationsApiDeploymentCellHealthRequest {
	r.serviceEnvironmentID = &serviceEnvironmentID
	return r
}

func (r ApiOperationsApiDeploymentCellHealthRequest) Execute() (*DeploymentCellHealthDetail, *http.Response, error) {
	return r.ApiService.OperationsApiDeploymentCellHealthExecute(r)
}

/*
OperationsApiDeploymentCellHealth DeploymentCellHealth operations-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOperationsApiDeploymentCellHealthRequest
*/
func (a *OperationsApiAPIService) OperationsApiDeploymentCellHealth(ctx context.Context) ApiOperationsApiDeploymentCellHealthRequest {
	return ApiOperationsApiDeploymentCellHealthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DeploymentCellHealthDetail
func (a *OperationsApiAPIService) OperationsApiDeploymentCellHealthExecute(r ApiOperationsApiDeploymentCellHealthRequest) (*DeploymentCellHealthDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeploymentCellHealthDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsApiAPIService.OperationsApiDeploymentCellHealth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/fleet/operations/deployment-cell-health"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.hostClusterID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hostClusterID", r.hostClusterID, "form", "")
	}
	if r.serviceID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "serviceID", r.serviceID, "form", "")
	}
	if r.serviceEnvironmentID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "serviceEnvironmentID", r.serviceEnvironmentID, "form", "")
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

type ApiOperationsApiListEventsRequest struct {
	ctx context.Context
	ApiService OperationsApiAPI
	nextPageToken *string
	pageSize *int64
	environmentType *string
	eventTypes *[]string
	serviceID *string
	serviceEnvironmentID *string
	instanceID *string
	startDate *time.Time
	endDate *time.Time
	productTierID *string
}

func (r ApiOperationsApiListEventsRequest) NextPageToken(nextPageToken string) ApiOperationsApiListEventsRequest {
	r.nextPageToken = &nextPageToken
	return r
}

func (r ApiOperationsApiListEventsRequest) PageSize(pageSize int64) ApiOperationsApiListEventsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiOperationsApiListEventsRequest) EnvironmentType(environmentType string) ApiOperationsApiListEventsRequest {
	r.environmentType = &environmentType
	return r
}

// The event types to filter by
func (r ApiOperationsApiListEventsRequest) EventTypes(eventTypes []string) ApiOperationsApiListEventsRequest {
	r.eventTypes = &eventTypes
	return r
}

// The service ID to list events for
func (r ApiOperationsApiListEventsRequest) ServiceID(serviceID string) ApiOperationsApiListEventsRequest {
	r.serviceID = &serviceID
	return r
}

// The service environment ID to list events for
func (r ApiOperationsApiListEventsRequest) ServiceEnvironmentID(serviceEnvironmentID string) ApiOperationsApiListEventsRequest {
	r.serviceEnvironmentID = &serviceEnvironmentID
	return r
}

// The instance ID to list events for
func (r ApiOperationsApiListEventsRequest) InstanceID(instanceID string) ApiOperationsApiListEventsRequest {
	r.instanceID = &instanceID
	return r
}

// Start date of the events
func (r ApiOperationsApiListEventsRequest) StartDate(startDate time.Time) ApiOperationsApiListEventsRequest {
	r.startDate = &startDate
	return r
}

// End date of the events
func (r ApiOperationsApiListEventsRequest) EndDate(endDate time.Time) ApiOperationsApiListEventsRequest {
	r.endDate = &endDate
	return r
}

func (r ApiOperationsApiListEventsRequest) ProductTierID(productTierID string) ApiOperationsApiListEventsRequest {
	r.productTierID = &productTierID
	return r
}

func (r ApiOperationsApiListEventsRequest) Execute() (*ListServiceProviderEventsResult, *http.Response, error) {
	return r.ApiService.OperationsApiListEventsExecute(r)
}

/*
OperationsApiListEvents ListEvents operations-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOperationsApiListEventsRequest
*/
func (a *OperationsApiAPIService) OperationsApiListEvents(ctx context.Context) ApiOperationsApiListEventsRequest {
	return ApiOperationsApiListEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListServiceProviderEventsResult
func (a *OperationsApiAPIService) OperationsApiListEventsExecute(r ApiOperationsApiListEventsRequest) (*ListServiceProviderEventsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListServiceProviderEventsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsApiAPIService.OperationsApiListEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/fleet/operations/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextPageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextPageToken", r.nextPageToken, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.environmentType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "environmentType", r.environmentType, "form", "")
	}
	if r.eventTypes != nil {
		t := *r.eventTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "eventTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "eventTypes", t, "form", "multi")
		}
	}
	if r.serviceID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "serviceID", r.serviceID, "form", "")
	}
	if r.serviceEnvironmentID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "serviceEnvironmentID", r.serviceEnvironmentID, "form", "")
	}
	if r.instanceID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "instanceID", r.instanceID, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.productTierID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productTierID", r.productTierID, "form", "")
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

type ApiOperationsApiServiceHealthRequest struct {
	ctx context.Context
	ApiService OperationsApiAPI
	serviceID *string
	serviceEnvironmentID *string
}

// The service ID to get the health for
func (r ApiOperationsApiServiceHealthRequest) ServiceID(serviceID string) ApiOperationsApiServiceHealthRequest {
	r.serviceID = &serviceID
	return r
}

// The service environment ID to get the health for
func (r ApiOperationsApiServiceHealthRequest) ServiceEnvironmentID(serviceEnvironmentID string) ApiOperationsApiServiceHealthRequest {
	r.serviceEnvironmentID = &serviceEnvironmentID
	return r
}

func (r ApiOperationsApiServiceHealthRequest) Execute() (*ServiceHealthSummary, *http.Response, error) {
	return r.ApiService.OperationsApiServiceHealthExecute(r)
}

/*
OperationsApiServiceHealth ServiceHealth operations-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOperationsApiServiceHealthRequest
*/
func (a *OperationsApiAPIService) OperationsApiServiceHealth(ctx context.Context) ApiOperationsApiServiceHealthRequest {
	return ApiOperationsApiServiceHealthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServiceHealthSummary
func (a *OperationsApiAPIService) OperationsApiServiceHealthExecute(r ApiOperationsApiServiceHealthRequest) (*ServiceHealthSummary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceHealthSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsApiAPIService.OperationsApiServiceHealth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/fleet/operations/service-health"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceID == nil {
		return localVarReturnValue, nil, reportError("serviceID is required and must be specified")
	}
	if r.serviceEnvironmentID == nil {
		return localVarReturnValue, nil, reportError("serviceEnvironmentID is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "serviceID", r.serviceID, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "serviceEnvironmentID", r.serviceEnvironmentID, "form", "")
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
