/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	ctrl "github.com/radius-project/radius/pkg/armrpc/asyncoperation/controller"
	rpv1 "github.com/radius-project/radius/pkg/rp/v1"
	"github.com/radius-project/radius/pkg/ucp/resources"
)

var _ ctrl.Controller = (*DeleteResource)(nil)

// DeleteResource is the async operation controller to delete Applications.Core/Containers resource.
type DeleteResource struct {
	ctrl.BaseController
}

// NewDeleteResource creates a new DeleteResource controller with the given options.
func NewDeleteResource(opts ctrl.Options) (ctrl.Controller, error) {
	return &DeleteResource{ctrl.NewBaseAsyncController(opts)}, nil
}

// Run retrieves a resource from the database, parses its ID, gets its data model, converts it to a deployment
// data model, deletes the resource from the deployment processor, and deletes the resource from the database.
// It returns an error if any of these steps fail.
func (c *DeleteResource) Run(ctx context.Context, request *ctrl.Request) (ctrl.Result, error) {
	obj, err := c.DatabaseClient().Get(ctx, request.ResourceID)
	if err != nil {
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: err.Error()}), err
	}

	// This code is general and we might be processing an async job for a resource or a scope, so using the general Parse function.
	id, err := resources.Parse(request.ResourceID)
	if err != nil {
		return ctrl.Result{}, err
	}

	dataModel, err := getDataModel(id)
	if err != nil {
		return ctrl.Result{}, err
	}

	if err = obj.As(dataModel); err != nil {
		return ctrl.Result{}, err
	}

	deploymentDataModel, ok := dataModel.(rpv1.DeploymentDataModel)
	if !ok {
		return ctrl.NewFailedResult(v1.ErrorDetails{Message: "deployment data model conversion error"}), nil
	}

	err = c.DeploymentProcessor().Delete(ctx, id, deploymentDataModel.OutputResources())
	if err != nil {
		return ctrl.Result{}, err
	}

	err = c.DatabaseClient().Delete(ctx, request.ResourceID)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, err
}
