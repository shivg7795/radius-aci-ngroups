// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package utils

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2015-06-15/storage"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/Azure/radius/test/config"
)

// GetGroup gets info on the resource group in use
func GetGroup(ctx context.Context, groupName string) (resources.Group, error) {
	groupsClient, err := config.AzureConfig.GetGroupsClient()
	if err != nil {
		return resources.Group{}, err
	}
	return groupsClient.Get(ctx, groupName)
}

// DeleteGroup deletes the resource group
func DeleteGroup(ctx context.Context, groupName string) (result resources.GroupsDeleteFuture, err error) {
	groupsClient, err := config.AzureConfig.GetGroupsClient()
	if err != nil {
		return resources.GroupsDeleteFuture{}, err
	}
	return groupsClient.Delete(ctx, groupName)
}

// ListResourcesInResourceGroup gets all resources in resource group
func ListResourcesInResourceGroup(ctx context.Context, groupName string) (resources.ListResultPage, error) {
	resourcesClient, err := config.AzureConfig.GetResourcesClient()
	if err != nil {
		return resources.ListResultPage{}, err
	}
	var top10 int32 = 10
	resourcesInRg, err := resourcesClient.ListByResourceGroup(ctx, groupName, "", "", &top10)
	return resourcesInRg, err
}

// GetContainer gets info about an existing container.
func GetContainer(ctx context.Context, accountName, accountGroupName, containerName string) (azblob.ContainerURL, error) {
	c, err := getContainerURL(ctx, accountName, accountGroupName, containerName)
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	_, err = c.GetProperties(ctx, azblob.LeaseAccessConditions{})
	return c, err
}

var (
	blobFormatString = `https://%s.blob.core.windows.net`
)

func getContainerURL(ctx context.Context, accountName, accountGroupName, containerName string) (azblob.ContainerURL, error) {
	key, err := getAccountPrimaryKey(ctx, accountName, accountGroupName)
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	c, err := azblob.NewSharedKeyCredential(accountName, key)
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	p := azblob.NewPipeline(c, azblob.PipelineOptions{})
	u, err := url.Parse(fmt.Sprintf(blobFormatString, accountName))
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	service := azblob.NewServiceURL(*u, p)
	container := service.NewContainerURL(containerName)
	return container, nil
}

func getAccountPrimaryKey(ctx context.Context, accountName, accountGroupName string) (string, error) {
	response, err := GetAccountKeys(ctx, accountName, accountGroupName)
	if err != nil {
		return "", err
	}
	return *response.Key1, nil
}

// GetAccountKeys gets the storage account keys
func GetAccountKeys(ctx context.Context, accountName, accountGroupName string) (storage.AccountKeys, error) {
	accountsClient, err := config.AzureConfig.GetStorageAccountsClient()
	if err != nil {
		return storage.AccountKeys{}, err
	}
	return accountsClient.ListKeys(ctx, accountGroupName, accountName)
}

// AcquireStorageContainerLease acquires an infinite lease on the storage container
func AcquireStorageContainerLease(ctx context.Context, accountName, accountGroupName, containerName string) error {
	container, err := GetContainer(ctx, accountName, accountGroupName, containerName)
	if err != nil {
		return err
	}

	_, err = container.AcquireLease(ctx, "", -1, azblob.ModifiedAccessConditions{})
	return err
}

// BreakStorageContainerLease breaks a lease on the storage container
func BreakStorageContainerLease(ctx context.Context, accountName, accountGroupName, containerName string) error {
	// Break lease on the test cluster to make it available for other tests
	container, err := GetContainer(ctx, accountName, accountGroupName, containerName)
	if err != nil {
		return nil
	}
	_, err = container.BreakLease(ctx, -1, azblob.ModifiedAccessConditions{})
	return err
}
