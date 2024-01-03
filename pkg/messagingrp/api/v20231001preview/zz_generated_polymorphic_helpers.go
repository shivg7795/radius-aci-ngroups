// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20231001preview

import "encoding/json"

func unmarshalEnvironmentComputeClassification(rawMsg json.RawMessage) (EnvironmentComputeClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EnvironmentComputeClassification
	switch m["kind"] {
	case "aci":
		b = &AzureContainerInstanceCompute{}
	case "kubernetes":
		b = &KubernetesCompute{}
	default:
		b = &EnvironmentCompute{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

