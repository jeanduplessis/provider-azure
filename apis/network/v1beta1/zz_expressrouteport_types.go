/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ExpressRoutePortIdentityObservation struct {
}

type ExpressRoutePortIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Express Route Port.
	// +kubebuilder:validation:Required
	IdentityIds []*string `json:"identityIds" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Express Route Port. Only possible value is UserAssigned.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ExpressRoutePortObservation struct {

	// The EtherType of the Express Route Port.
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype,omitempty"`

	// The resource GUID of the Express Route Port.
	GUID *string `json:"guid,omitempty" tf:"guid,omitempty"`

	// The ID of the Express Route Port.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of link blocks as defined below.
	// +kubebuilder:validation:Optional
	Link1 []Link1Observation `json:"link1,omitempty" tf:"link1,omitempty"`

	// A list of link blocks as defined below.
	// +kubebuilder:validation:Optional
	Link2 []Link2Observation `json:"link2,omitempty" tf:"link2,omitempty"`

	// The maximum transmission unit of the Express Route Port.
	Mtu *string `json:"mtu,omitempty" tf:"mtu,omitempty"`
}

type ExpressRoutePortParameters struct {

	// Bandwidth of the Express Route Port in Gbps. Changing this forces a new Express Route Port to be created.
	// +kubebuilder:validation:Required
	BandwidthInGbps *float64 `json:"bandwidthInGbps" tf:"bandwidth_in_gbps,omitempty"`

	// The billing type of the Express Route Port. Possible values are MeteredData and UnlimitedData.
	// +kubebuilder:validation:Optional
	BillingType *string `json:"billingType,omitempty" tf:"billing_type,omitempty"`

	// The encapsulation method used for the Express Route Port. Changing this forces a new Express Route Port to be created. Possible values are: Dot1Q, QinQ.
	// +kubebuilder:validation:Required
	Encapsulation *string `json:"encapsulation" tf:"encapsulation,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []ExpressRoutePortIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// A list of link blocks as defined below.
	// +kubebuilder:validation:Optional
	Link1 []Link1Parameters `json:"link1,omitempty" tf:"link1,omitempty"`

	// A list of link blocks as defined below.
	// +kubebuilder:validation:Optional
	Link2 []Link2Parameters `json:"link2,omitempty" tf:"link2,omitempty"`

	// The Azure Region where the Express Route Port should exist. Changing this forces a new Express Route Port to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the peering location that this Express Route Port is physically mapped to. Changing this forces a new Express Route Port to be created.
	// +kubebuilder:validation:Required
	PeeringLocation *string `json:"peeringLocation" tf:"peering_location,omitempty"`

	// The name of the Resource Group where the Express Route Port should exist. Changing this forces a new Express Route Port to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Express Route Port.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type Link1Observation struct {

	// The connector type of the Express Route Port Link.
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type,omitempty"`

	// The ID of this Express Route Port Link.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The interface name of the Azure router associated with the Express Route Port Link.
	InterfaceName *string `json:"interfaceName,omitempty" tf:"interface_name,omitempty"`

	// The ID that maps from the Express Route Port Link to the patch panel port.
	PatchPanelID *string `json:"patchPanelId,omitempty" tf:"patch_panel_id,omitempty"`

	// The ID that maps from the patch panel port to the rack.
	RackID *string `json:"rackId,omitempty" tf:"rack_id,omitempty"`

	// The name of the Azure router associated with the Express Route Port Link.
	RouterName *string `json:"routerName,omitempty" tf:"router_name,omitempty"`
}

type Link1Parameters struct {

	// Whether enable administration state on the Express Route Port Link? Defaults to false.
	// +kubebuilder:validation:Optional
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`

	// The ID of the Key Vault Secret that contains the Mac security CAK key for this Express Route Port Link.
	// +kubebuilder:validation:Optional
	MacsecCakKeyvaultSecretID *string `json:"macsecCakKeyvaultSecretId,omitempty" tf:"macsec_cak_keyvault_secret_id,omitempty"`

	// The MACSec cipher used for this Express Route Port Link. Possible values are GcmAes128 and GcmAes256. Defaults to GcmAes128.
	// +kubebuilder:validation:Optional
	MacsecCipher *string `json:"macsecCipher,omitempty" tf:"macsec_cipher,omitempty"`

	// The ID of the Key Vault Secret that contains the MACSec CKN key for this Express Route Port Link.
	// +kubebuilder:validation:Optional
	MacsecCknKeyvaultSecretID *string `json:"macsecCknKeyvaultSecretId,omitempty" tf:"macsec_ckn_keyvault_secret_id,omitempty"`
}

type Link2Observation struct {

	// The connector type of the Express Route Port Link.
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type,omitempty"`

	// The ID of this Express Route Port Link.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The interface name of the Azure router associated with the Express Route Port Link.
	InterfaceName *string `json:"interfaceName,omitempty" tf:"interface_name,omitempty"`

	// The ID that maps from the Express Route Port Link to the patch panel port.
	PatchPanelID *string `json:"patchPanelId,omitempty" tf:"patch_panel_id,omitempty"`

	// The ID that maps from the patch panel port to the rack.
	RackID *string `json:"rackId,omitempty" tf:"rack_id,omitempty"`

	// The name of the Azure router associated with the Express Route Port Link.
	RouterName *string `json:"routerName,omitempty" tf:"router_name,omitempty"`
}

type Link2Parameters struct {

	// Whether enable administration state on the Express Route Port Link? Defaults to false.
	// +kubebuilder:validation:Optional
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`

	// The ID of the Key Vault Secret that contains the Mac security CAK key for this Express Route Port Link.
	// +kubebuilder:validation:Optional
	MacsecCakKeyvaultSecretID *string `json:"macsecCakKeyvaultSecretId,omitempty" tf:"macsec_cak_keyvault_secret_id,omitempty"`

	// The MACSec cipher used for this Express Route Port Link. Possible values are GcmAes128 and GcmAes256. Defaults to GcmAes128.
	// +kubebuilder:validation:Optional
	MacsecCipher *string `json:"macsecCipher,omitempty" tf:"macsec_cipher,omitempty"`

	// The ID of the Key Vault Secret that contains the MACSec CKN key for this Express Route Port Link.
	// +kubebuilder:validation:Optional
	MacsecCknKeyvaultSecretID *string `json:"macsecCknKeyvaultSecretId,omitempty" tf:"macsec_ckn_keyvault_secret_id,omitempty"`
}

// ExpressRoutePortSpec defines the desired state of ExpressRoutePort
type ExpressRoutePortSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRoutePortParameters `json:"forProvider"`
}

// ExpressRoutePortStatus defines the observed state of ExpressRoutePort.
type ExpressRoutePortStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRoutePortObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRoutePort is the Schema for the ExpressRoutePorts API. Manages a Express Route Port.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRoutePort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRoutePortSpec   `json:"spec"`
	Status            ExpressRoutePortStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRoutePortList contains a list of ExpressRoutePorts
type ExpressRoutePortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRoutePort `json:"items"`
}

// Repository type metadata.
var (
	ExpressRoutePort_Kind             = "ExpressRoutePort"
	ExpressRoutePort_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRoutePort_Kind}.String()
	ExpressRoutePort_KindAPIVersion   = ExpressRoutePort_Kind + "." + CRDGroupVersion.String()
	ExpressRoutePort_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRoutePort_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRoutePort{}, &ExpressRoutePortList{})
}
