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

type LogAnalyticsSavedSearchObservation struct {

	// The Log Analytics Saved Search ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LogAnalyticsSavedSearchParameters struct {

	// The category that the Saved Search will be listed under. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Category *string `json:"category" tf:"category,omitempty"`

	// The name that Saved Search will be displayed as. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The function alias if the query serves as a function. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	FunctionAlias *string `json:"functionAlias,omitempty" tf:"function_alias,omitempty"`

	// The function parameters if the query serves as a function. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	FunctionParameters []*string `json:"functionParameters,omitempty" tf:"function_parameters,omitempty"`

	// Specifies the ID of the Log Analytics Workspace that the Saved Search will be associated with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// The query expression for the saved search. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// A mapping of tags which should be assigned to the Logs Analytics Saved Search. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LogAnalyticsSavedSearchSpec defines the desired state of LogAnalyticsSavedSearch
type LogAnalyticsSavedSearchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsSavedSearchParameters `json:"forProvider"`
}

// LogAnalyticsSavedSearchStatus defines the observed state of LogAnalyticsSavedSearch.
type LogAnalyticsSavedSearchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsSavedSearchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSavedSearch is the Schema for the LogAnalyticsSavedSearchs API. Manages a Log Analytics (formally Operational Insights) Saved Search.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsSavedSearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsSavedSearchSpec   `json:"spec"`
	Status            LogAnalyticsSavedSearchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSavedSearchList contains a list of LogAnalyticsSavedSearchs
type LogAnalyticsSavedSearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsSavedSearch `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsSavedSearch_Kind             = "LogAnalyticsSavedSearch"
	LogAnalyticsSavedSearch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsSavedSearch_Kind}.String()
	LogAnalyticsSavedSearch_KindAPIVersion   = LogAnalyticsSavedSearch_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsSavedSearch_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsSavedSearch_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsSavedSearch{}, &LogAnalyticsSavedSearchList{})
}
