// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20240901preview

// ContainerGroupProfileClientCreateOrUpdateResponse contains the response from method ContainerGroupProfileClient.CreateOrUpdate.
type ContainerGroupProfileClientCreateOrUpdateResponse struct {
	ContainerGroupProfile

	// AzureAsyncOperation contains the information returned from the Azure-AsyncOperation header response.
	AzureAsyncOperation *string
}

// ContainerGroupProfileClientDeleteResponse contains the response from method ContainerGroupProfileClient.Delete.
type ContainerGroupProfileClientDeleteResponse struct {
	// AzureAsyncOperation contains the information returned from the Azure-AsyncOperation header response.
	AzureAsyncOperation *string

	// Location contains the information returned from the Location header response.
	Location *string
}

// ContainerGroupProfileClientGetResponse contains the response from method ContainerGroupProfileClient.Get.
type ContainerGroupProfileClientGetResponse struct {
	ContainerGroupProfile
}

// ContainerGroupProfileClientListByResourceGroupResponse contains the response from method ContainerGroupProfileClient.NewListByResourceGroupPager.
type ContainerGroupProfileClientListByResourceGroupResponse struct {
	// The container group profile list response
	ContainerGroupProfileListResult
}

// ContainerGroupProfileClientListBySubscriptionResponse contains the response from method ContainerGroupProfileClient.NewListBySubscriptionPager.
type ContainerGroupProfileClientListBySubscriptionResponse struct {
	// The container group profile list response
	ContainerGroupProfileListResult
}

// ContainerGroupProfileClientUpdateResponse contains the response from method ContainerGroupProfileClient.Update.
type ContainerGroupProfileClientUpdateResponse struct {
	ContainerGroupProfile

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string
}

// ContainerGroupsClientCreateOrUpdateResponse contains the response from method ContainerGroupsClient.BeginCreateOrUpdate.
type ContainerGroupsClientCreateOrUpdateResponse struct {
	// A container group.
	ContainerGroup
}

// ContainerGroupsClientDeleteResponse contains the response from method ContainerGroupsClient.BeginDelete.
type ContainerGroupsClientDeleteResponse struct {
	// A container group.
	ContainerGroup
}

// ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse contains the response from method ContainerGroupsClient.GetOutboundNetworkDependenciesEndpoints.
type ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse struct {
	// Response for network dependencies, always empty list.
	StringArray []*string
}

// ContainerGroupsClientGetResponse contains the response from method ContainerGroupsClient.Get.
type ContainerGroupsClientGetResponse struct {
	// A container group.
	ContainerGroup
}

// ContainerGroupsClientListByResourceGroupResponse contains the response from method ContainerGroupsClient.NewListByResourceGroupPager.
type ContainerGroupsClientListByResourceGroupResponse struct {
	// The container group list response that contains the container group properties.
	ContainerGroupListResult
}

// ContainerGroupsClientListResponse contains the response from method ContainerGroupsClient.NewListPager.
type ContainerGroupsClientListResponse struct {
	// The container group list response that contains the container group properties.
	ContainerGroupListResult
}

// ContainerGroupsClientRestartResponse contains the response from method ContainerGroupsClient.BeginRestart.
type ContainerGroupsClientRestartResponse struct {
	// placeholder for future response values
}

// ContainerGroupsClientStartResponse contains the response from method ContainerGroupsClient.BeginStart.
type ContainerGroupsClientStartResponse struct {
	// placeholder for future response values
}

// ContainerGroupsClientStopResponse contains the response from method ContainerGroupsClient.Stop.
type ContainerGroupsClientStopResponse struct {
	// placeholder for future response values
}

// ContainerGroupsClientUpdateResponse contains the response from method ContainerGroupsClient.Update.
type ContainerGroupsClientUpdateResponse struct {
	// A container group.
	ContainerGroup
}

// ContainersClientAttachResponse contains the response from method ContainersClient.Attach.
type ContainersClientAttachResponse struct {
	// The information for the output stream from container attach.
	ContainerAttachResponse
}

// ContainersClientExecuteCommandResponse contains the response from method ContainersClient.ExecuteCommand.
type ContainersClientExecuteCommandResponse struct {
	// The information for the container exec command.
	ContainerExecResponse
}

// ContainersClientListLogsResponse contains the response from method ContainersClient.ListLogs.
type ContainersClientListLogsResponse struct {
	// The logs.
	Logs
}

// LocationClientListCachedImagesResponse contains the response from method LocationClient.NewListCachedImagesPager.
type LocationClientListCachedImagesResponse struct {
	// The response containing cached images.
	CachedImagesListResult
}

// LocationClientListCapabilitiesResponse contains the response from method LocationClient.NewListCapabilitiesPager.
type LocationClientListCapabilitiesResponse struct {
	// The response containing list of capabilities.
	CapabilitiesListResult
}

// LocationClientListUsageResponse contains the response from method LocationClient.NewListUsagePager.
type LocationClientListUsageResponse struct {
	// The response containing the usage data
	UsageListResult
}

// NGroupsClientCreateOrUpdateResponse contains the response from method NGroupsClient.BeginCreateOrUpdate.
type NGroupsClientCreateOrUpdateResponse struct {
	// Describes a nGroup.
	NGroup
}

// NGroupsClientDeleteResponse contains the response from method NGroupsClient.BeginDelete.
type NGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// NGroupsClientGetResponse contains the response from method NGroupsClient.Get.
type NGroupsClientGetResponse struct {
	// Describes a nGroup.
	NGroup
}

// NGroupsClientListByResourceGroupResponse contains the response from method NGroupsClient.NewListByResourceGroupPager.
type NGroupsClientListByResourceGroupResponse struct {
	// The nGroups list response that contains the nGroups properties.
	NGroupsListResult
}

// NGroupsClientListResponse contains the response from method NGroupsClient.NewListPager.
type NGroupsClientListResponse struct {
	// The nGroups list response that contains the nGroups properties.
	NGroupsListResult
}

// NGroupsClientRestartResponse contains the response from method NGroupsClient.BeginRestart.
type NGroupsClientRestartResponse struct {
	// placeholder for future response values
}

// NGroupsClientStartResponse contains the response from method NGroupsClient.BeginStart.
type NGroupsClientStartResponse struct {
	// placeholder for future response values
}

// NGroupsClientStopResponse contains the response from method NGroupsClient.Stop.
type NGroupsClientStopResponse struct {
	// placeholder for future response values
}

// NGroupsClientUpdateResponse contains the response from method NGroupsClient.BeginUpdate.
type NGroupsClientUpdateResponse struct {
	// Describes a nGroup.
	NGroup
}

// NGroupsSKUsClientGetResponse contains the response from method NGroupsSKUsClient.NewGetPager.
type NGroupsSKUsClientGetResponse struct {
	// List of SKU definitions. NGroups offer a single sku
	NGroupsSKUsList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The operation list response that contains all operations for Azure Container Instance service.
	OperationListResult
}

// SubnetServiceAssociationLinkClientDeleteResponse contains the response from method SubnetServiceAssociationLinkClient.BeginDelete.
type SubnetServiceAssociationLinkClientDeleteResponse struct {
	// placeholder for future response values
}

