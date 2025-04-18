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

package datamodel

import (
	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	"github.com/radius-project/radius/pkg/portableresources"
	rpv1 "github.com/radius-project/radius/pkg/rp/v1"
)

// ExtenderResourceType is the resource type for Extender portable resources.
const ExtenderResourceType = "Applications.Core/extenders"

// Extender represents Extender portable resource.
type Extender struct {
	v1.BaseResource

	// Properties is the properties of the resource.
	Properties ExtenderProperties `json:"properties"`

	// PortableResourceMetadata represents internal DataModel properties common to all portable resource types.
	PortableResourceMetadata
}

// ApplyDeploymentOutput updates the Status of Properties of the Extender resource with the DeployedOutputResources and returns no error.
func (r *Extender) ApplyDeploymentOutput(do rpv1.DeploymentOutput) error {
	r.Properties.Status.OutputResources = do.DeployedOutputResources
	return nil
}

// OutputResources returns the OutputResources of the Extender resource.
func (r *Extender) OutputResources() []rpv1.OutputResource {
	return r.Properties.Status.OutputResources
}

// ResourceMetadata returns an adapter that provides standardized access to BasicResourceProperties of the Extender resource.
func (r *Extender) ResourceMetadata() rpv1.BasicResourcePropertiesAdapter {
	return &r.Properties.BasicResourceProperties
}

// ResourceTypeName returns the resource type of the extender resource.
func (r *Extender) ResourceTypeName() string {
	return ExtenderResourceType
}

// Recipe returns the ResourceRecipe associated with the Extender if the ResourceProvisioning is not set to Manual,
// otherwise it returns nil.
func (r *Extender) GetRecipe() *portableresources.ResourceRecipe {
	if r.Properties.ResourceProvisioning == portableresources.ResourceProvisioningManual {
		return nil
	}
	return &r.Properties.ResourceRecipe
}

// SetRecipe sets the recipe information.
func (r *Extender) SetRecipe(recipe *portableresources.ResourceRecipe) {
	r.Properties.ResourceRecipe = *recipe
}

// ExtenderProperties represents the properties of Extender resource.
type ExtenderProperties struct {
	rpv1.BasicResourceProperties
	// Additional properties for the resource
	AdditionalProperties map[string]any `json:"additionalProperties,omitempty"`
	// Secrets values provided for the resource
	Secrets map[string]any `json:"secrets,omitempty"`
	// The recipe used to automatically deploy underlying infrastructure for the Extender
	ResourceRecipe portableresources.ResourceRecipe `json:"recipe,omitempty"`
	// Specifies how the underlying service/resource is provisioned and managed
	ResourceProvisioning portableresources.ResourceProvisioning `json:"resourceProvisioning,omitempty"`
}
