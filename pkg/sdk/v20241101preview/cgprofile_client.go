//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20241101preview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type CGProfileClient struct {
	subscriptionID string
	pl runtime.Pipeline
}

// NewCGProfileClient creates a new instance of CGProfileClient with the specified values.
// subscriptionID - The ID of the target subscription. The value must be an UUID.
// pl - the pipeline used for sending requests and handling responses.
func NewCGProfileClient(subscriptionID string, pl runtime.Pipeline) *CGProfileClient {
	client := &CGProfileClient{
		subscriptionID: subscriptionID,
		pl: pl,
	}
	return client
}

// CreateOrUpdate - Create a CGProfile if it doesn't exist or update an existing CGProfile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerGroupProfileName - ContainerGroupProfile name.
// containerGroupProfile - The ContainerGroupProfile object.
// options - CGProfileClientCreateOrUpdateOptions contains the optional parameters for the CGProfileClient.CreateOrUpdate
// method.
func (client *CGProfileClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerGroupProfileName string, containerGroupProfile ContainerGroupProfile, options *CGProfileClientCreateOrUpdateOptions) (CGProfileClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, containerGroupProfileName, containerGroupProfile, options)
	if err != nil {
		return CGProfileClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CGProfileClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CGProfileClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CGProfileClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerGroupProfileName string, containerGroupProfile ContainerGroupProfile, options *CGProfileClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroupProfiles/{containerGroupProfileName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupProfileName == "" {
		return nil, errors.New("parameter containerGroupProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupProfileName}", url.PathEscape(containerGroupProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, containerGroupProfile)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CGProfileClient) createOrUpdateHandleResponse(resp *http.Response) (CGProfileClientCreateOrUpdateResponse, error) {
	result := CGProfileClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroupProfile); err != nil {
		return CGProfileClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a container group profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerGroupProfileName - ContainerGroupProfile name.
// options - CGProfileClientBeginDeleteOptions contains the optional parameters for the CGProfileClient.BeginDelete method.
func (client *CGProfileClient) BeginDelete(ctx context.Context, resourceGroupName string, containerGroupProfileName string, options *CGProfileClientBeginDeleteOptions) (*runtime.Poller[CGProfileClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, containerGroupProfileName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CGProfileClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[CGProfileClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a container group profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-11-01-preview
func (client *CGProfileClient) deleteOperation(ctx context.Context, resourceGroupName string, containerGroupProfileName string, options *CGProfileClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerGroupProfileName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	 return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CGProfileClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerGroupProfileName string, options *CGProfileClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroupProfiles/{containerGroupProfileName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupProfileName == "" {
		return nil, errors.New("parameter containerGroupProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupProfileName}", url.PathEscape(containerGroupProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of the specified container group profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerGroupProfileName - ContainerGroupProfile name.
// options - CGProfileClientGetOptions contains the optional parameters for the CGProfileClient.Get method.
func (client *CGProfileClient) Get(ctx context.Context, resourceGroupName string, containerGroupProfileName string, options *CGProfileClientGetOptions) (CGProfileClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerGroupProfileName, options)
	if err != nil {
		return CGProfileClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CGProfileClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CGProfileClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CGProfileClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerGroupProfileName string, options *CGProfileClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroupProfiles/{containerGroupProfileName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupProfileName == "" {
		return nil, errors.New("parameter containerGroupProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupProfileName}", url.PathEscape(containerGroupProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CGProfileClient) getHandleResponse(resp *http.Response) (CGProfileClientGetResponse, error) {
	result := CGProfileClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroupProfile); err != nil {
		return CGProfileClientGetResponse{}, err
	}
	return result, nil
}

// GetByRevisionNumber - Gets the properties of the specified revision of the container group profile in the given subscription
// and resource group. The operation returns the properties of container group profile including
// containers, image registry credentials, restart policy, IP address type, OS type, volumes, current revision number, etc.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerGroupProfileName - ContainerGroupProfile name.
// revisionNumber - The revision number of the container group profile.
// options - CGProfileClientGetByRevisionNumberOptions contains the optional parameters for the CGProfileClient.GetByRevisionNumber
// method.
func (client *CGProfileClient) GetByRevisionNumber(ctx context.Context, resourceGroupName string, containerGroupProfileName string, revisionNumber string, options *CGProfileClientGetByRevisionNumberOptions) (CGProfileClientGetByRevisionNumberResponse, error) {
	req, err := client.getByRevisionNumberCreateRequest(ctx, resourceGroupName, containerGroupProfileName, revisionNumber, options)
	if err != nil {
		return CGProfileClientGetByRevisionNumberResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CGProfileClientGetByRevisionNumberResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CGProfileClientGetByRevisionNumberResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByRevisionNumberHandleResponse(resp)
}

// getByRevisionNumberCreateRequest creates the GetByRevisionNumber request.
func (client *CGProfileClient) getByRevisionNumberCreateRequest(ctx context.Context, resourceGroupName string, containerGroupProfileName string, revisionNumber string, options *CGProfileClientGetByRevisionNumberOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroupProfiles/{containerGroupProfileName}/revisions/{revisionNumber}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupProfileName == "" {
		return nil, errors.New("parameter containerGroupProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupProfileName}", url.PathEscape(containerGroupProfileName))
	if revisionNumber == "" {
		return nil, errors.New("parameter revisionNumber cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{revisionNumber}", url.PathEscape(revisionNumber))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByRevisionNumberHandleResponse handles the GetByRevisionNumber response.
func (client *CGProfileClient) getByRevisionNumberHandleResponse(resp *http.Response) (CGProfileClientGetByRevisionNumberResponse, error) {
	result := CGProfileClientGetByRevisionNumberResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroupProfile); err != nil {
		return CGProfileClientGetByRevisionNumberResponse{}, err
	}
	return result, nil
}

// NewListAllRevisionsPager - Get a list of all the revisions of the specified container group profile in the given subscription
// and resource group. This operation returns properties of each revision of the specified container
// group profile including containers, image registry credentials, restart policy, IP address type, OS type volumes, revision
// number, etc.
// Generated from API version 2024-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerGroupProfileName - ContainerGroupProfile name.
// options - CGProfileClientListAllRevisionsOptions contains the optional parameters for the CGProfileClient.ListAllRevisions
// method.
func (client *CGProfileClient) NewListAllRevisionsPager(resourceGroupName string, containerGroupProfileName string, options *CGProfileClientListAllRevisionsOptions) (*runtime.Pager[CGProfileClientListAllRevisionsResponse]) {
	return runtime.NewPager(runtime.PagingHandler[CGProfileClientListAllRevisionsResponse]{
		More: func(page CGProfileClientListAllRevisionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CGProfileClientListAllRevisionsResponse) (CGProfileClientListAllRevisionsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllRevisionsCreateRequest(ctx, resourceGroupName, containerGroupProfileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CGProfileClientListAllRevisionsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CGProfileClientListAllRevisionsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CGProfileClientListAllRevisionsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllRevisionsHandleResponse(resp)
		},
	})
}

// listAllRevisionsCreateRequest creates the ListAllRevisions request.
func (client *CGProfileClient) listAllRevisionsCreateRequest(ctx context.Context, resourceGroupName string, containerGroupProfileName string, options *CGProfileClientListAllRevisionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroupProfiles/{containerGroupProfileName}/revisions"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupProfileName == "" {
		return nil, errors.New("parameter containerGroupProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupProfileName}", url.PathEscape(containerGroupProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllRevisionsHandleResponse handles the ListAllRevisions response.
func (client *CGProfileClient) listAllRevisionsHandleResponse(resp *http.Response) (CGProfileClientListAllRevisionsResponse, error) {
	result := CGProfileClientListAllRevisionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroupProfileListResult); err != nil {
		return CGProfileClientListAllRevisionsResponse{}, err
	}
	return result, nil
}

// Update - Update a specified container group profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2024-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// containerGroupProfileName - ContainerGroupProfile name.
// properties - The container group profile properties that need to be updated.
// options - CGProfileClientUpdateOptions contains the optional parameters for the CGProfileClient.Update method.
func (client *CGProfileClient) Update(ctx context.Context, resourceGroupName string, containerGroupProfileName string, properties ContainerGroupProfilePatch, options *CGProfileClientUpdateOptions) (CGProfileClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, containerGroupProfileName, properties, options)
	if err != nil {
		return CGProfileClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CGProfileClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CGProfileClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CGProfileClient) updateCreateRequest(ctx context.Context, resourceGroupName string, containerGroupProfileName string, properties ContainerGroupProfilePatch, options *CGProfileClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroupProfiles/{containerGroupProfileName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupProfileName == "" {
		return nil, errors.New("parameter containerGroupProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupProfileName}", url.PathEscape(containerGroupProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// updateHandleResponse handles the Update response.
func (client *CGProfileClient) updateHandleResponse(resp *http.Response) (CGProfileClientUpdateResponse, error) {
	result := CGProfileClientUpdateResponse{}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroupProfile); err != nil {
		return CGProfileClientUpdateResponse{}, err
	}
	return result, nil
}

