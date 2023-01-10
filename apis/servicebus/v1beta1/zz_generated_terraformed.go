/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ServiceBusNamespace
func (mg *ServiceBusNamespace) GetTerraformResourceType() string {
	return "azurerm_servicebus_namespace"
}

// GetConnectionDetailsMapping for this ServiceBusNamespace
func (tr *ServiceBusNamespace) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"default_primary_connection_string": "status.atProvider.defaultPrimaryConnectionString", "default_primary_key": "status.atProvider.defaultPrimaryKey", "default_secondary_connection_string": "status.atProvider.defaultSecondaryConnectionString", "default_secondary_key": "status.atProvider.defaultSecondaryKey"}
}

// GetObservation of this ServiceBusNamespace
func (tr *ServiceBusNamespace) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceBusNamespace
func (tr *ServiceBusNamespace) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceBusNamespace
func (tr *ServiceBusNamespace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceBusNamespace
func (tr *ServiceBusNamespace) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceBusNamespace
func (tr *ServiceBusNamespace) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ServiceBusNamespace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceBusNamespace) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceBusNamespaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceBusNamespace) GetTerraformSchemaVersion() int {
	return 1
}