// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package v20220315privatepreview

import (
	"time"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/linkrp"
	"github.com/project-radius/radius/pkg/to"
)

func toProvisioningStateDataModel(state *ProvisioningState) v1.ProvisioningState {
	if state == nil {
		return v1.ProvisioningStateAccepted
	}

	switch *state {
	case ProvisioningStateUpdating:
		return v1.ProvisioningStateUpdating
	case ProvisioningStateDeleting:
		return v1.ProvisioningStateDeleting
	case ProvisioningStateAccepted:
		return v1.ProvisioningStateAccepted
	case ProvisioningStateSucceeded:
		return v1.ProvisioningStateSucceeded
	case ProvisioningStateFailed:
		return v1.ProvisioningStateFailed
	case ProvisioningStateCanceled:
		return v1.ProvisioningStateCanceled
	case ProvisioningStateProvisioning:
		return v1.ProvisioningStateProvisioning
	default:
		return v1.ProvisioningStateAccepted
	}
}

func fromProvisioningStateDataModel(state v1.ProvisioningState) *ProvisioningState {
	var converted ProvisioningState
	switch state {
	case v1.ProvisioningStateUpdating:
		converted = ProvisioningStateUpdating
	case v1.ProvisioningStateDeleting:
		converted = ProvisioningStateDeleting
	case v1.ProvisioningStateAccepted:
		converted = ProvisioningStateAccepted
	case v1.ProvisioningStateSucceeded:
		converted = ProvisioningStateSucceeded
	case v1.ProvisioningStateFailed:
		converted = ProvisioningStateFailed
	case v1.ProvisioningStateCanceled:
		converted = ProvisioningStateCanceled
	default:
		converted = ProvisioningStateAccepted
	}

	return &converted
}

func toResourceProvisiongDataModel(provisioning *ResourceProvisioning) linkrp.ResourceProvisioning {
	if provisioning == nil {
		return linkrp.ResourceProvisioningRecipe
	}
	switch *provisioning {
	case ResourceProvisioningManual:
		return linkrp.ResourceProvisioningManual
	default:
		return linkrp.ResourceProvisioning(*provisioning)
	}
}

func fromResourceProvisioningDataModel(provisioning linkrp.ResourceProvisioning) *ResourceProvisioning {
	var converted ResourceProvisioning
	switch provisioning {
	case linkrp.ResourceProvisioningManual:
		converted = ResourceProvisioningManual
	default:
		converted = ResourceProvisioningRecipe
	}

	return &converted
}

func unmarshalTimeString(ts string) *time.Time {
	var tt timeRFC3339
	_ = tt.UnmarshalText([]byte(ts))
	return (*time.Time)(&tt)
}

func fromSystemDataModel(s v1.SystemData) *SystemData {
	return &SystemData{
		CreatedBy:          to.Ptr(s.CreatedBy),
		CreatedByType:      (*CreatedByType)(to.Ptr(s.CreatedByType)),
		CreatedAt:          unmarshalTimeString(s.CreatedAt),
		LastModifiedBy:     to.Ptr(s.LastModifiedBy),
		LastModifiedByType: (*CreatedByType)(to.Ptr(s.LastModifiedByType)),
		LastModifiedAt:     unmarshalTimeString(s.LastModifiedAt),
	}
}

func toRecipeDataModel(r *Recipe) linkrp.LinkRecipe {
	if r == nil {
		return linkrp.LinkRecipe{}
	}

	recipe := linkrp.LinkRecipe{
		Name: to.String(r.Name),
	}

	if r.Parameters != nil {
		recipe.Parameters = r.Parameters
	}
	return recipe
}

func fromRecipeDataModel(r linkrp.LinkRecipe) *Recipe {
	return &Recipe{
		Name:       to.Ptr(r.Name),
		Parameters: r.Parameters,
	}
}

func toResourcesDataModel(r []*ResourceReference) []*linkrp.ResourceReference {
	if r == nil {
		return nil
	}
	resources := make([]*linkrp.ResourceReference, len(r))
	for i, resource := range r {
		resources[i] = &linkrp.ResourceReference{
			ID: to.String(resource.ID),
		}
	}
	return resources
}

func fromResourcesDataModel(r []*linkrp.ResourceReference) []*ResourceReference {
	if r == nil {
		return nil
	}
	resources := make([]*ResourceReference, len(r))
	for i, resource := range r {
		resources[i] = &ResourceReference{
			ID: to.Ptr(resource.ID),
		}
	}
	return resources
}
