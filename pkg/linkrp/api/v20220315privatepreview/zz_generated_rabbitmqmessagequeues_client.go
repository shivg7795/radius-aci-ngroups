//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// RabbitMqMessageQueuesClient contains the methods for the RabbitMqMessageQueues group.
// Don't use this type directly, use NewRabbitMqMessageQueuesClient() instead.
type RabbitMqMessageQueuesClient struct {
	host string
	rootScope string
	pl runtime.Pipeline
}

// NewRabbitMqMessageQueuesClient creates a new instance of RabbitMqMessageQueuesClient with the specified values.
// rootScope - The scope in which the resource is present. For Azure resource this would be /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRabbitMqMessageQueuesClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RabbitMqMessageQueuesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RabbitMqMessageQueuesClient{
		rootScope: rootScope,
		host: ep,
pl: pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a RabbitMQMessageQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// rabbitMQMessageQueueName - The name of the RabbitMQMessageQueue link resource
// resource - Resource create parameters.
// options - RabbitMqMessageQueuesClientCreateOrUpdateOptions contains the optional parameters for the RabbitMqMessageQueuesClient.CreateOrUpdate
// method.
func (client *RabbitMqMessageQueuesClient) CreateOrUpdate(ctx context.Context, rabbitMQMessageQueueName string, resource RabbitMQMessageQueueResource, options *RabbitMqMessageQueuesClientCreateOrUpdateOptions) (RabbitMqMessageQueuesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, rabbitMQMessageQueueName, resource, options)
	if err != nil {
		return RabbitMqMessageQueuesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RabbitMqMessageQueuesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RabbitMqMessageQueuesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RabbitMqMessageQueuesClient) createOrUpdateCreateRequest(ctx context.Context, rabbitMQMessageQueueName string, resource RabbitMQMessageQueueResource, options *RabbitMqMessageQueuesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/rabbitMQMessageQueues/{rabbitMQMessageQueueName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQMessageQueueName == "" {
		return nil, errors.New("parameter rabbitMQMessageQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQMessageQueueName}", url.PathEscape(rabbitMQMessageQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RabbitMqMessageQueuesClient) createOrUpdateHandleResponse(resp *http.Response) (RabbitMqMessageQueuesClientCreateOrUpdateResponse, error) {
	result := RabbitMqMessageQueuesClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return RabbitMqMessageQueuesClientCreateOrUpdateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.RabbitMQMessageQueueResource); err != nil {
		return RabbitMqMessageQueuesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing RabbitMQMessageQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// rabbitMQMessageQueueName - The name of the RabbitMQMessageQueue link resource
// options - RabbitMqMessageQueuesClientDeleteOptions contains the optional parameters for the RabbitMqMessageQueuesClient.Delete
// method.
func (client *RabbitMqMessageQueuesClient) Delete(ctx context.Context, rabbitMQMessageQueueName string, options *RabbitMqMessageQueuesClientDeleteOptions) (RabbitMqMessageQueuesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, rabbitMQMessageQueueName, options)
	if err != nil {
		return RabbitMqMessageQueuesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RabbitMqMessageQueuesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return RabbitMqMessageQueuesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *RabbitMqMessageQueuesClient) deleteCreateRequest(ctx context.Context, rabbitMQMessageQueueName string, options *RabbitMqMessageQueuesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/rabbitMQMessageQueues/{rabbitMQMessageQueueName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQMessageQueueName == "" {
		return nil, errors.New("parameter rabbitMQMessageQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQMessageQueueName}", url.PathEscape(rabbitMQMessageQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *RabbitMqMessageQueuesClient) deleteHandleResponse(resp *http.Response) (RabbitMqMessageQueuesClientDeleteResponse, error) {
	result := RabbitMqMessageQueuesClientDeleteResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return RabbitMqMessageQueuesClientDeleteResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	return result, nil
}

// Get - Retrieves information about a RabbitMQMessageQueueResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// rabbitMQMessageQueueName - The name of the RabbitMQMessageQueue link resource
// options - RabbitMqMessageQueuesClientGetOptions contains the optional parameters for the RabbitMqMessageQueuesClient.Get
// method.
func (client *RabbitMqMessageQueuesClient) Get(ctx context.Context, rabbitMQMessageQueueName string, options *RabbitMqMessageQueuesClientGetOptions) (RabbitMqMessageQueuesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, rabbitMQMessageQueueName, options)
	if err != nil {
		return RabbitMqMessageQueuesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RabbitMqMessageQueuesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RabbitMqMessageQueuesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RabbitMqMessageQueuesClient) getCreateRequest(ctx context.Context, rabbitMQMessageQueueName string, options *RabbitMqMessageQueuesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/rabbitMQMessageQueues/{rabbitMQMessageQueueName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQMessageQueueName == "" {
		return nil, errors.New("parameter rabbitMQMessageQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQMessageQueueName}", url.PathEscape(rabbitMQMessageQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RabbitMqMessageQueuesClient) getHandleResponse(resp *http.Response) (RabbitMqMessageQueuesClientGetResponse, error) {
	result := RabbitMqMessageQueuesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RabbitMQMessageQueueResource); err != nil {
		return RabbitMqMessageQueuesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRootScopePager - Lists information about all RabbitMQMessageQueueResources in the given root scope
// Generated from API version 2022-03-15-privatepreview
// options - RabbitMqMessageQueuesClientListByRootScopeOptions contains the optional parameters for the RabbitMqMessageQueuesClient.ListByRootScope
// method.
func (client *RabbitMqMessageQueuesClient) NewListByRootScopePager(options *RabbitMqMessageQueuesClientListByRootScopeOptions) (*runtime.Pager[RabbitMqMessageQueuesClientListByRootScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[RabbitMqMessageQueuesClientListByRootScopeResponse]{
		More: func(page RabbitMqMessageQueuesClientListByRootScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RabbitMqMessageQueuesClientListByRootScopeResponse) (RabbitMqMessageQueuesClientListByRootScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRootScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RabbitMqMessageQueuesClientListByRootScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RabbitMqMessageQueuesClientListByRootScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RabbitMqMessageQueuesClientListByRootScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRootScopeHandleResponse(resp)
		},
	})
}

// listByRootScopeCreateRequest creates the ListByRootScope request.
func (client *RabbitMqMessageQueuesClient) listByRootScopeCreateRequest(ctx context.Context, options *RabbitMqMessageQueuesClientListByRootScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/rabbitMQMessageQueues"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRootScopeHandleResponse handles the ListByRootScope response.
func (client *RabbitMqMessageQueuesClient) listByRootScopeHandleResponse(resp *http.Response) (RabbitMqMessageQueuesClientListByRootScopeResponse, error) {
	result := RabbitMqMessageQueuesClientListByRootScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RabbitMQMessageQueueResourceListResult); err != nil {
		return RabbitMqMessageQueuesClientListByRootScopeResponse{}, err
	}
	return result, nil
}

// ListSecrets - Lists secrets values for the specified RabbitMQMessageQueue resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// rabbitMQMessageQueueName - The name of the RabbitMQMessageQueue link resource
// options - RabbitMqMessageQueuesClientListSecretsOptions contains the optional parameters for the RabbitMqMessageQueuesClient.ListSecrets
// method.
func (client *RabbitMqMessageQueuesClient) ListSecrets(ctx context.Context, rabbitMQMessageQueueName string, options *RabbitMqMessageQueuesClientListSecretsOptions) (RabbitMqMessageQueuesClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, rabbitMQMessageQueueName, options)
	if err != nil {
		return RabbitMqMessageQueuesClientListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RabbitMqMessageQueuesClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RabbitMqMessageQueuesClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *RabbitMqMessageQueuesClient) listSecretsCreateRequest(ctx context.Context, rabbitMQMessageQueueName string, options *RabbitMqMessageQueuesClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/rabbitMQMessageQueues/{rabbitMQMessageQueueName}/listSecrets"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if rabbitMQMessageQueueName == "" {
		return nil, errors.New("parameter rabbitMQMessageQueueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rabbitMQMessageQueueName}", url.PathEscape(rabbitMQMessageQueueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *RabbitMqMessageQueuesClient) listSecretsHandleResponse(resp *http.Response) (RabbitMqMessageQueuesClientListSecretsResponse, error) {
	result := RabbitMqMessageQueuesClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RabbitMQListSecretsResult); err != nil {
		return RabbitMqMessageQueuesClientListSecretsResponse{}, err
	}
	return result, nil
}

